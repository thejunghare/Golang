package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const TASK_SERVER = "http://localhost:5000"

func Read() []Task {
	content, err := ioutil.ReadFile("task.json")
	if err != nil {
		log.Fatal(err)
	}
	tasks := []Task{}
	_ = json.Unmarshal(content, &tasks)
	return tasks
}

func Write(task Task) {
	tasks := Read()
	tasks = append(tasks, task)
	content, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("task.json", content, 0644)
	if err != nil {
		log.Println(err)
	}
}

func ReadAPI() []Task {
	response, err := http.Get(TASK_SERVER + "/list")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	tasks := []Task{}
	_ = json.Unmarshal(responseData, &tasks)

	return tasks
}

func WriteAPI(task Task) TaskCreateResp {
	url := TASK_SERVER + "/save"
	taskJson, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(taskJson)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)

	taskCreateResp := TaskCreateResp{}
	_ = json.Unmarshal(responseData, &taskCreateResp)

	return taskCreateResp
}

func InitConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
}

func SetCommandValues(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them,
		// so bind them to their equivalent keys with underscores,
		// e.g. --server_url to server_url
		envVarSuffix := ""
		if strings.Contains(f.Name, "-") {
			envVarSuffix = strings.ReplaceAll(f.Name, "-", "_")
			viper.BindEnv(f.Name, envVarSuffix)
		}

		// Apply the viper config value to the flag
		// when the flag is not set and viper has a value
		if !f.Changed && viper.GetString(envVarSuffix) != "" {
			val := viper.GetString(envVarSuffix)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

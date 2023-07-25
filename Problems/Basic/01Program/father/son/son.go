package son

func init()  {
	fmt.Println("Son Package initalized")
}

type Son struct {
	Name string
}

func (s Son) Data (name string) string {
	return  "Son" + name

}
## Variables

_Variables are defined as the container to store data values._

Types of variables:

1. int
2. float64
3. string
4. bool

Syntax:

1. _Declaring variables_

   - var variablename type = value

2. _Using := operator_

   - variablename := value

3. _Variable declaration without initial value_

   - var variablename type

4. _Difference between var and :=_
   - **var** : can be used both inside and outside of an fuction
   - **:=** : can only be used inside functions
   - **:=** : declaration and value assignment cannot be done separately.
5. _Declaring multiple variables_

   - var a,b,c ,d int = 1,2,3,4
   - d, e := 7, "Variable"

6. _Variable declaration in a block_
   - var ( <br>
     a int <br>
     b int = 1 <br>
     c string = "Hello" <br>
     )

7. *Multi-word variable names*
	- Camel Case : variableName := value
	- Pascal Case : VariableName := value
	- Snake Case : variable_name := value
	
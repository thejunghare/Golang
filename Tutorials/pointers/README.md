## Pointers

Pointers are variable used to store memory address of another variable, memory address are always stored in ***hexadecimal format***.

*Variables are the names given to a memory location where the actual data is stored.To access the stored data we need the address of that particular memory location, which is stored in pointers.*

## Two important operators used in pointers

1. `*` opreator :- Know as Dereferncing operator <br>
   used to decalre an pointer variable and access the value stored in address.
2. & opreator :- Know as Address operator<br>
   used to return address of variable or access address of a variable.

## Declaring a pointer
```
var pointer_name *Datatype
```
## Initalizing a pointer
```
pointer_name = &variable_name
```

### Some important points
1. uninitailized pointer will always have a nil value.
```
var pointer_name *int
fmt.Print(pointer_name) // output --> nil
```

2. Declaring and initalizing of pointer can be done in one go.
```
var pointer_name *int = &variable_name
```

3. Using := syntax for declaring and initalizing an pointer
```
variable_name := 123
pointer_variable := &variable_name
```
***You can do same with the var keyword -- this is type interence concept of var keyword as there is no need to specify the type of the pointer variable*** 

4. Using dereferncing operator to access value stored in vaiable
```
variable_name := 123
fmt.Print(*variable_name)
```

## important notes

For some composite type, such as struct, you can get address without declaring other variable to get the address from.
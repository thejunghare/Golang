## Structures
Structures are user-defined type that allows to group/combine items of possibly different types into a single type.

Any real-world entity which has some set of properties/fields can be  represented as struct.
For example address can have fields name(string), state(string), pincode(int) etc, this all type can be stored under one type Address

### syntax:
```
type Type_name struct {
    field1 type
    ...
}
```

### Declaring variable of type struct
```
var variable_name struct_name
fmt.Print(variable_name) // Output --> {0 0}
```
**NOTE**: *All fields will be initalized with zero-value by default*

### Initalizing struct
```
variable_name := struct_name {
    fieldValue,
    ....,
}
```

### Declaring and initalizing with name fields
```
variable_name := struct_name {
    field_name: field_value,
    field_name: field_value,
    ...
}
```

### Access the fields of struct
```
variable_name.field_name
```

# Interface

Interfaces define a set of method signatures. An interface specifies what methods a type must have in order to be considered as implementing that interface.

In simple words interface lets us group together different types based on their behaviour.

Interfaces are like rule book which define define a method (properties) to be considered as a specific type.

for example: type Human has properties like walk(), talk(), etc, so to be a type of Human you should implement the methods(properties) of type Human.

### syntax

```
type Human interface {}
```

### Implementing multiple interface

A type can implement more than one interface.

### Embedding intrefaces

Although go does not offer inheritance, it is possible to ceate a new interfaces by embedding other interfaces.

```
type interface_one interface{
    method() type
}

type interface_two interface {
    method() type
}

type interface_three interface {
    interface_one
    interface_two
}
```

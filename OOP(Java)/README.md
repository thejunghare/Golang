## Programming languages

There are five major categories of programming languages.

1. Procedural programming languages :- Programming languages that follow a procedural approach, where programs are structured around procedures or steps to be executed in a specific order.
   1. C and C++
   2. Java
2. Functional programming languages :- Rather than focusing on the execution of statements, functional languages focus on the output of mathematical functions and evaluations.
   1. Scala
   2. Erlang
3. Object-oriented programming languages :- oriented programming is about creating objects that contain both data and methods.
   1. Java
   2. C++
4. Scripting languages :- Scripting languages are used to automate repetitive tasks, manage dynamic web content, or support processes in larger applications
   1. PHP
   2. Node.js
5. Logic programming languages :- Instead of telling a computer what to do, a logic programming language expresses a series of facts and rules to instruct the computer on how to make decisions.
   1. Prolog
   2. Absys

We attempt to provide a straightforward examination of the object-oriented programming language by utilizing the java programming language.

## OOP (Java)

OOP stands for object-Oriented Programming. Object-Oriented Programming is a methodology or paradigm to design a program using classes and objects.

It simplifies software development and maintenance by providing some concepts:

1. Object
2. Class
3. Inheritance
4. Polymorphism
5. Abstraction
6. Encapsulation

## Class

Classes can be defined as group of objects with common properties.

## Object

An object is an instance of a class.

Object can have three properties:-

1. state :- represents data
2. Behavior :- represents behavior
3. identity :- an object identity is typically implemented via a unique ID.

## Method

A method sometimes aslo called as function is a block of code which only runs when it is called.

## Constructors

A constructor in java is a special method, used to set initial values of object attributes.

_Note_ :-

1. constructor name must match the class name, and it cannot have a return type.
2. Constructor is called when object is create.
3. All classes have constructor by default.

## Encapsulation

Hidding user data from users is encapsulation, to achieve encapsulation you must :-

1. Declare class variables/attributes as private.
2. provide public get and set methods to access and update the value of a private variable.

## Inheritance

Inheriting methods and attributes from one class to another is called as inheritance. Inheritance can be grouped in two categories :-

1. subclass (child) :- The class that inherits from another class.
2. superclass (parent) :- The class being inherited from.

In java we use _extends_ keyword to inherit.

## Plymorphism

Polymorphoism means _many forms_. polymorphoism like inheritance lets us inherit attributes and methods from another class. Basically allows us to perform a single action in different ways.

## Abstraction

Abstraction is the process of hiding certain details and showing only essential information.

Absraction can be achieved with either abstract class or interface

1. Abstract class :- is a restricted class that cannot be used to create objects.
2. Abstract method :- can only used in an abstract class, it does not have a body. Body is provided by subclass.

## Interface

In simple words interface lets us group together different types based on their behaviour.

Interfaces are like rule book which define define a method (properties) to be considered as a specific type.

for example: type Human has properties like walk(), talk(), etc, so to be a type of Human you should implement the methods(properties) of type Human.

FOR JAVA
Another way to achieve abstraction in Java, is with interfaces.
An interface is a completely "abstract class"

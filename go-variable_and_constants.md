## GO VARIABLES
Variables are containers for storing data values.

# Go Variable Types

    int- stores integers (whole numbers), such as 123 or -123
    float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
    string - stores text, such as "Hello World". String values are surrounded by double quotes
    bool- stores values with two states: true or false

# Declaring Variables

In Go, there are two ways to declare a variable:
1. With the var keyword:

Use the var keyword, followed by variable name and type:
Syntax
**var variablename type = value**

Note: You always have to specify either type or value (or both).
2. With the := sign:

Use the := sign, followed by the variable value:
Syntax
**variablename := value**

In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

It is not possible to declare a variable using :=, without assigning a value to it.

# Variable Declaration Without Initial Value

In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:
```go
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```
It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.

# Difference Between var and :=

1. Can be used inside and outside of functions while Can only be used inside functions

2. Variable declaration and value assignment can be done separately while Variable declaration and value assignment cannot be done separately (must be done in the same line)

# Go Multiple Variable Declaration

In Go, it is possible to declare multiple variables on the same line.
```go
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```
 If you use the type keyword, it is only possible to declare one type of variable per line.
 If the type keyword is not specified, you can declare different types of variables on the same line:
```go
package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```
# Go Variable Declaration in a Block

Multiple variable declarations can also be grouped together into a block for greater readability:
```go
package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```
# Go Variable Naming Rules

A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Go variable naming rules:

    A variable name must start with a letter or an underscore character (_)
    A variable name cannot start with a digit
    A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
    Variable names are case-sensitive (age, Age and AGE are three different variables)
    There is no limit on the length of the variable name
    A variable name cannot contain spaces
    The variable name cannot be any Go keywords

# Multi-Word Variable Names

Variable names with more than one word can be difficult to read.
Here ar techniques to make them more readable :
1. Camel case
2. Pascal case
3. Snake case

# Go Constants

If a variable should have a fixed value that cannot be changed, you can use the const keyword.

The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
```go 
 const CONSTNAME type = value 
 ```
 # Constant Rules

    Constant names follow the same naming rules as variables
    Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
    Constants can be declared both inside and outside of a function

# Constant Types

There are two types of constants:

    Typed constants
    Untyped constants



# GO ARRAYS

Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
Declare an Array

In Go, there are two ways to declare an array:
1. With the var keyword:
Syntax
**var array_name = [length]datatype{values}**

or

**var array_name = [...]datatype{values}**
2. With the := sign:
Syntax
**array_name := [length]datatype{values}**

or

**array_name := [...]datatype{values}**

The length specifies the number of elements to store in the array. In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).

# GO SLICES
Slices are similar to arrays but more flexible in the sense that unlike arrays, the length of slices can grow and shrink as you see fit.

In Go, there are several ways to create a slice:

1. Using the []datatype{values} format
Syntax :
```go
 slice_name := []datatype{values} 
```
2. Create a slice from an array
Syntax:
```go
var myarray = [length]datatype{values} // An array
myslice := myarray[start:end] // A slice made from the array
```
3. Using the make() function
```go
 slice_name := make([]type, length, capacity)
```
Go can accept, change, append or modify slices.
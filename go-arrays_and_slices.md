# Go Arrays

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

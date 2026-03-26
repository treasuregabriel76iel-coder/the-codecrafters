# GO OUTPUT FUNCTIONS
Go has three functions to output text:
```go
    Print() // prints argumentes with their default format
    Println() // prints arguments on new line
    Printf() // prints function first formats its argument based on the given formatting verb and then prints them.
```
# Go Formatting Verbs
**Formatting Verbs for Printf()**

Go offers several formatting verbs that can be used with the Printf() function.
General Formatting Verbs
The following verbs can be used with all data types:
%v 	Prints the value in the default format
%#v 	Prints the value in Go-syntax format
%T 	Prints the type of the value
%% 	Prints the % sign

# Integer Formatting Verbs
The following verbs can be used with the integer data type:
%b 	Base 2
%d 	Base 10
%+d 	Base 10 and always show sign
%o 	Base 8
%O 	Base 8, with leading 0o
%x 	Base 16, lowercase
%X 	Base 16, uppercase
%#x 	Base 16, with leading 0x
%4d 	Pad with spaces (width 4, right justified)
%-4d 	Pad with spaces (width 4, left justified)
%04d 	Pad with zeroes (width 4)

# String Formatting Verbs
The following verbs can be used with the string data type:
%s 	Prints the value as plain string
%q 	Prints the value as a double-quoted string
%8s 	Prints the value as plain string (width 8, right justified)
%-8s 	Prints the value as plain string (width 8, left justified)
%x 	Prints the value as hex dump of byte values
% x 	Prints the value as hex dump with spaces

# Boolean Formatting Verbs
The following verb can be used with the boolean data type:
%t 	Value of the boolean operator in true or false format (same as using %v)

# Float Formatting Verbs
The following verbs can be used with the float data type:
%e 	Scientific notation with 'e' as exponent
%f 	Decimal point, no exponent
%.2f 	Default width, precision 2
%6.2f 	Width 6, precision 2
%g 	Exponent as needed, only necessary digits
# Go For Loops

The for loop loops through a block of code a specified number of times.

The for loop is the only loop available in Go.

Loops are handy if you want to run the same code over and over again, each time with a different value.

Each execution of a loop is called an iteration.

The for loop can take up to three statements:
Syntax
```go
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
```
statement1 Initializes the loop counter value.

statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

statement3 Increases the loop counter value.

 These statements don't need to be present as loops arguments. However, they need to be present in the code in some form.
# The continue Statement

The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

# The break Statement

The break statement is used to break/terminate the loop execution.

# Nested Loops

It is a loop inside another loop 
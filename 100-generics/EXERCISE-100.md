# Exercises for chapter Generics

## a) Implementing a generic function

Write a generic function that takes a slice of any type and returns a slice of the same type, but with all elements reversed.

## b) Implementing a generic type

Create a generic type that represents a stack. The stack should be able to hold any type of element. Implement the following methods:

- `Push` to add an element to the stack.
- `Pop` to remove and return the last element from the stack.
- `Peek` to return the last element from the stack without removing it.
- `Len` to return the number of elements in the stack.
- `IsEmpty` to return a boolean indicating whether the stack is empty.
- `String` to return a string representation of the stack.
- `Iter` to return a `Seq` instance that can be used to iteratore over the stack in LIFO order.

Feel free to reuse the stack code from the Golang Basics training.
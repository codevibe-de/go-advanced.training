# Exercises for chapter Generics

## a) Implementing a generic function

Write a generic function that takes a slice of any type and returns a slice of the same type,
but with all elements reversed.

## b) Implementing a generic type

Create a generic type that represents a **set of items**. A set should be able to hold any type of element.

Implement the following methods:

- `Add` to add an element to the set
- `Remove` to remove an element from the set
- `Contains` to check if an element is in the set
- `Len` to return the number of elements in the set
- `IsEmpty` to return a boolean indicating whether the set is empty
- `String` to return a string representation of the set
- `Iter` to return a `Seq` instance that can be used to iterate over the set

Which existing Golang data structure can we use to implement a set efficiently?
Which constraint do we need to enforce on the generic type?
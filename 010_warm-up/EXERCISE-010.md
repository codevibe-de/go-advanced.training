# Warm-up Exercise

In this exercise, we want to recap several topics from the Golang Basics training,
especially:

- structs
- methods
- pointers
- loops
- functions as first-class citizens
- optionally closures

Also, the data structure created here will serve as a building block when working with generics later.

## a) A chained list

Develop a data structure for a chained list, which meets these criteria:

- the basic type is a list-item that holds a current value of type `string` and a pointer to the next list item
- a method (not function) named `add` takes a value argument and creates and returns a new list-item
- a method named `forEach` takes a function argument that will be called for the list item's value and each following one

Create a list holding at least three values and let those values be printed out and/or collected into a slice.
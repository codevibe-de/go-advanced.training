# Exercises for chapter Reflection

## a) Uppercasing struct fields

Write a utility function that takes a `reflect.Value` parameter, which is a _reflection object_ of a struct.

The function should uppercase all letters of each field in the struct, which is of type `string`.

## b) Calling a command struct

Create a struct that represents a command. Each exported method of the struct represents a subcommand, which
can be called from the command line. 

Your application greets the user on StdOut and then scans StdIn for the name
of a subcommand (method). 
If a method with that name is found in the struct, it is called and its result is printed to StdOut.

Your main function **could** start like this:

```go
	cmd := MyCommand{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for scanner.Scan() {
		cmd.callSubcommand(scanner.Text())
		fmt.Print("$ ")
	}
```

However, feel free to modify this exercise to your liking. 
You could also use a package for parsing command line arguments, such as `flag` or `flaggy` or `kong` or `cobra`.

Before you start coding, think about the following questions:

- What is the general purpose of your command?
- And which typical subcommands could you also implement, i.e., which are commonly found in other commands?

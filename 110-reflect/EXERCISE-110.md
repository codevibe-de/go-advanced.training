# Exercises for chapter Reflection

## a) Uppercasing struct fields

Write a utility function that takes a `reflect.Value` parameter being a reflection object of a struct.

The function should uppercase all letters of each field in the struct, which is of type string.

## b) Calling a command struct

Create a struct that represents a command. Each exported method of the struct represents a subcommand, which
can be called from the command line. Your application greets the user on StdOut and then scans StdIn for the name
of a subcommand (method). If the corresponding method is found, it is called and its result is printed to StdOut.

Your main function can start like this:

```go
	cmd := MyCommand{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for scanner.Scan() {
		cmd.callSubcommand(scanner.Text())
		fmt.Print("$ ")
	}
```

Which typical subcommands could you implement, i.e., are commonly found in other commands?
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func stuff() {
	someStruct := GenerateCommand{}
	m := reflect.ValueOf(someStruct).MethodByName("Help")
	//var argValues []reflect.Value = []reflect.Value{
	//	reflect.ValueOf(1),
	//	reflect.ValueOf(2),
	//}
	var returnValues []reflect.Value = m.Call(nil)
	fmt.Println(returnValues)
}

func main() {
	//stuff()

	cmd := GenerateCommand{}
	fmt.Println(reflect.TypeOf(cmd))
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for scanner.Scan() {
		cmd.callSubcommand(scanner.Text())
		fmt.Print("$ ")
	}
}

type GenerateCommand struct {
}

func (gc GenerateCommand) usage() {
	fmt.Println(gc.Help())
}

func (gc GenerateCommand) Version() string {
	return "v1.0.0"
}

func (gc GenerateCommand) Help() string {
	s := "Call any of these methods:"
	t := reflect.TypeOf(gc)
	names := make([]string, 0, t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		names = append(names, t.Method(i).Name)
	}
	s += "\n- " + strings.Join(names, "\n- ")
	return s
}

func (gc GenerateCommand) Generate() string {
	return "Generated content"
}

func (gc GenerateCommand) Exit() {
	os.Exit(0)
}

func (gc GenerateCommand) callSubcommand(c string) {
	c = strings.Title(c)
	method := reflect.ValueOf(gc).MethodByName(c)
	if !method.IsValid() {
		fmt.Printf("ERROR: No subcommand for name '%s'\n", c)
		return
	}
	resultArr := method.Call(nil)
	if len(resultArr) > 0 {
		fmt.Println(resultArr[0])
	}
}

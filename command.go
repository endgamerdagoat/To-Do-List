package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new item, with a name")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an item by its index, with a new name. id:new_name")
	flag.IntVar(&cf.Del, "delete", -1, "Delete an item at a specific index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Mark an item as done by its index")
	flag.BoolVar(&cf.List, "list", false, "List all items")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute (todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format. Please use id:new_name")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: Invalid Index")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid Command")
	}
}

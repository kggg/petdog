package commands

import (
	"flag"
	"fmt"
)

type Command struct {
	FirstArgs string
	Flatset   flag.FlagSet
}

func (c *Command) Usage() {
	fmt.Println("Usage of project")
}

func (c *Command) Help(arg string) {
	fmt.Println("help ", arg)
}

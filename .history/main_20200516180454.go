package main

import (
       "flag"
       "fmt"
       "petdog/api"
       "petdog/commands"
)

func main() {
	cmd := &commands.Command{}
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		cmd.Usage()
		return
	}

	switch args[0]{
	case "help":
		cmd.Help(args[1])
	case "new":
		if len(args) < 2 {
			fmt.Println("no project name append params new")
			return
	}
	err := api.NewProject(args[1])
	if err != nil {
			fmt.Println(err)
			return
	}
	case "make":
	case "run":
	default:
		fmt.Println("nothing to do")
	}

	if args[0] == "help" {
		cmd.Help(args[1])
		return
	}

	if args[0] == "new" {

}
}
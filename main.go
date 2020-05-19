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

	switch args[0] {
	case "help":
		if len(args) < 2 {
			fmt.Println("no project name append params new")
			return
		}
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
		if len(args) < 2 {
			fmt.Println("no specified template and directory")
			cmd.Usage()
			return
		}
		if len(args) < 3 {
			fmt.Println("no specified filename")
			cmd.Usage()
			return
		}
		maker, err := api.NewMaker(args[1], args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := maker.MakeFile(); err != nil {
			fmt.Println(err)
			return
		}
	case "run":
	default:
		fmt.Println("nothing to do")
	}
}

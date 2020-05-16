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

       if args[0] == "help" {
               cmd.Help(args[1])
               return
       }

       if args[0] == "new" {
               if len(args) < 2 {
                       fmt.Println("no project name append params new")
                       return
               }
               err := api.NewProject(args[1])
               if err != nil {
                       fmt.Println(err)
                       return
               }
       }
}
	
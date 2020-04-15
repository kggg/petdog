package main

import(
	""{{ .appname }}/routes""
)

func main()  {
	route := routes.Initrouter()
	route.Run(":8080")
}
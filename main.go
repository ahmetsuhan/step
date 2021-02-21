package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
		return
	}

	helpCommand := *flag.String("h",""," -h")
	saveCommand := *flag.String("s", "", " -s alias_name")

	flag.Parse()

	if helpCommand != "" {
		help()
		return
	}

	if saveCommand != "" {
		//
	}

}

func help(){
	fmt.Println(`
	-h			Outputs this help
	-s			Saves the alias for the spesified ssh
	-s -i			Save alias with a key path

	============
	Create Alias
	============
	step -s demo_1 -i /full/path/key IP@server
`)
}
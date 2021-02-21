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
		fmt.Printf("Please provide a password to use as a key")
		pass := readline()
		fmt.Printf("Retype the password")
		if pass != readline() {
			fmt.Printf("%s","Please make sure that password matches.")
			return
		}

		ciphertext := encrypt("line2", pass)
		fmt.Printf("%s",ciphertext)
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
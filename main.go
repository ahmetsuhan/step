package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
)

var dbPath string

func init(){
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	fmt.Println( usr.HomeDir )
	path := usr.HomeDir +"/step/"
	dbPath = path + "data.db"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	if _, err := os.Stat(dbPath); err != nil || os.IsNotExist(err) {
		// your code here if file exists
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
		return
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("Cannot stat database")
		return;
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
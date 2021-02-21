package main

import (
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/exec"
	"os/user"
)

var dbPath string

func init(){
	if len(os.Args) < 2 {
		help()
		return
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
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


	helpCommand := flag.String("h",""," -h")
	saveCommand := flag.String("s", "", " -s alias_name")
	path := flag.String("i",""," -i full path")
	flag.Parse()

	if *helpCommand != "" {
		help()
		return
	}

	if *saveCommand != "" {
		save(*path)
		return
	}
	return
	cmd:= exec.Command("ssh","")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run() // add error checking

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
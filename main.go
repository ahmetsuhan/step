package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
	}
}

func help(){
	fmt.Println(`step -h
						line 2
						line 3`)
}
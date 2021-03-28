package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(2)
	}
	//cmd := os.Args[1]
	//swtich cmd {
	//
	//}
}

package main

// go:generate sqliboiler --wipe sqlite3

import (
	"fmt"
	"github.com/amin1024/xtelbot/cmd"
)

func main() {
	fmt.Println("=========================================")
	fmt.Println("\t\txTelBot")
	fmt.Println("=========================================")

	cmd.Execute()
}

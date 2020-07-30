package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/etc/shadow")
	defer file.Close()
	if err == nil {
		fmt.Println("Yay! You're running as root!")
	} else {
		fmt.Println(":( You're a regular user and got the following error:", err)
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Hi , I am Hector"
	fmt.Println(strings.ToLower(message))
	fmt.Println(strings.ToUpper(message))
	fmt.Println(strings.ReplaceAll(message, "Hi , I am Hector", "Hi , I'm Hector"))

}

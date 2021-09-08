package main

import "fmt"

func main() {
	var (
		n1  int
		out string
	)
	fmt.Println("Enter one number of 1 to 5:")
	fmt.Scanln(&n1)
	switch n1 {
	case 1:
		out = "One"
	case 2:
		out = "Two"
	case 3:
		out = "Three"
	case 4:
		out = "Four"
	case 5:
		out = "Five"
	default:
		out = "isn't in 1 to 5"
	}
	fmt.Println("The number ingresed is ", n1, "\nThe name of number is ", out)
}

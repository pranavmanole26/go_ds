package main

import "fmt"

func main() {
	// str := "अआabc"
	var str string
	fmt.Println("Enter a string value")
	fmt.Scanf("%s", &str)
	revStr := ""

	for _, s := range str {
		revStr = string(s) + " " + revStr
	}

	fmt.Print(revStr)
}

package main

import "fmt"

func main() {
	name := "{{ .Name }}"
	fmt.Println("Hello " + name)
}

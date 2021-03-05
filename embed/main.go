package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed html1/* html2/*
var html embed.FS

func main() {

	index1, err := html.ReadFile("html1/index1.html")
	if err != nil {
		log.Fatal(err)
	}

	index2, err := html.ReadFile("html2/index2.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(index1))

	fmt.Println()
	fmt.Printf("----------------------------------------------------")
	fmt.Println()

	fmt.Print(string(index2))

	fmt.Println()
	fmt.Printf("----------------------------------------------------")
	fmt.Println()

}

//func main() {
//	hello.HelloEmbed()
//	fmt.Println(hello2)
//}

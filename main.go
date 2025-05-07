package main

import (
	"fmt"
	"pogodoc/go/sdk/client"
)

func main() {
	client := client.NewClient()
	fmt.Println(client)
}


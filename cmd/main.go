package main

import (
	"fmt"
	"os"

	"github.com/sandronister/fast-bolierplate-go/internals/cli"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please input folder name and package name")
		return
	}

	client := cli.NewService(os.Args[1])

	err := client.CreateInitialFolder()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = client.CreatePackage()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = client.CreatePrincipals()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = client.CreateInternal()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = client.CreateInfra()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}

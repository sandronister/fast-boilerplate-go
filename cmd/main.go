package main

import (
	"fmt"
	"os"

	"github.com/sandronister/fast-bolierplate-go/internals/cli"
	"github.com/sandronister/fast-bolierplate-go/pkg/utils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("To execute this program example")
		fmt.Println("fast-go <folder-name> <module-name>")
		return
	}

	client := cli.NewService(os.Args[1], os.Args[2])

	err := client.CreateInitialFolder()
	utils.Catch(&err)

	err = client.CreatePackage()
	utils.Catch(&err)

	err = client.GetPackage()
	utils.Catch(&err)

	err = client.CreatePrincipals()
	utils.Catch(&err)

	err = client.CreateConfig()
	utils.Catch(&err)

	err = client.CreateInternal()
	utils.Catch(&err)

	err = client.CreateInfra()
	utils.Catch(&err)

}

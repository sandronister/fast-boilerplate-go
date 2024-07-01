package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input package name")
		return
	}

	packageName := os.Args[1]
	cmd := exec.Command("go", "mod", "init", packageName)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	principals := []string{"internal", "pkg", "cmd", "config"}

	for _, principal := range principals {
		err := os.Mkdir(principal, 0755)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	internals := []string{"internal/entity", "internal/repository", "internal/infra", "internal/usecase"}

	for _, internal := range internals {
		err := os.MkdirAll(internal, 0755)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	infra := []string{"internal/infra/http", "internal/infra/database", "internal/infra/web", "internal/infra/handler"}

	for _, inf := range infra {
		err := os.MkdirAll(inf, 0755)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}

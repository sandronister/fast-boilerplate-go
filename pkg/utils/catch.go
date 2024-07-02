package utils

import (
	"fmt"
	"os"
)

func Catch(err *error) {
	if *err != nil {
		fmt.Println("Error: ", *err)
		os.Exit(1)
	}
}

package main

import (
	"boilerplate-cleanarch/bootstrap"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {

		fmt.Println("you must specified the project name")
		return

	}

	projectName := os.Args[1]
	err := bootstrap.Generate(projectName)
	if err != nil {

		fmt.Println(err.Error())

	}

}

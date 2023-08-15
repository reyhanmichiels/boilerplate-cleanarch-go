package bootstrap

import (
	"github.com/reyhanmichiels/boilerplate-cleanarch-go/util"
	"fmt"
	"os"
)

func Generate(projectName string) error {

	//initialize go mod
	if err := os.MkdirAll(projectName, os.ModePerm); err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName)

	}

	if err := os.Chdir(projectName); err != nil {

		return err

	}

	command := "go mod init " + projectName
	_, err := util.ExecuteCommand(command)
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Initialize go mod")

	}

	//bootstraping

	err = util.CreateFile("app/main.go")
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"app")
		fmt.Println("successed: Make file", projectName+"app/main.go")

	}

	err = os.MkdirAll("src/business/entity", os.ModePerm)
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"/src")
		fmt.Println("successed: Make directory", projectName+"/src/business")
		fmt.Println("successed: Make directory", projectName+"/src/business/entity")

	}

	err = util.CreateFile("src/business/handler/rest/rest.go")
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"/src/business/handler")
		fmt.Println("successed: Make directory", projectName+"/src/business/handler/rest")
		fmt.Println("successed: Make file", projectName+"/src/business/handler/rest/rest.go")

	}

	err = util.CreateFile("src/business/repository/repository.go")
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"/src/business/repository")
		fmt.Println("successed: Make file", projectName+"/src/business/repository/repository.go")

	}

	err = util.CreateFile("src/business/usecase/usecase.go")
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"/src/business/usecase")
		fmt.Println("successed: Make file", projectName+"/src/business/usecase/usecase.go")

	}

	err = util.CreateFile("src/sdk/conf/conf.go")
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"/src/sdk")
		fmt.Println("successed: Make directory", projectName+"/src/sdk/conf")
		fmt.Println("successed: Make file", projectName+"/src/sdk/conf/conf.go")

	}

	err = os.MkdirAll("src/sdk/database", os.ModePerm)
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"/src/sdk/database")

	}

	err = os.MkdirAll("src/sdk/library", os.ModePerm)
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"/src/sdk/library")

	}

	err = os.MkdirAll("src/sdk/middleware", os.ModePerm)
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make directory", projectName+"/src/sdk/middleware")

	}

	err = util.CreateFile(".env")
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make file", projectName+"/.env")

	}

	err = util.CreateFile(".env.example")
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make file", projectName+"/.env.example")

	}

	err = util.CreateFile(".gitignore")
	if err != nil {

		return err

	} else {

		fmt.Println("successed: Make file", projectName+"/.gitignore")

	}

	if err := os.Chdir("../"); err != nil {

		return err

	}

	return nil

}

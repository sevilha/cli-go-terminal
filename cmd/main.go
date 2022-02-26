package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
)

var commands = []string{"build", "deploy", "test", "exit"}

func build() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error to load .env file")
	}
}

func deploy() {

}

func test() {
	option := ""
	prompt := &survey.Select{
		Message: "Select a command:",
		Options: []string{"test-all", "test-file"},
	}

	survey.AskOne(prompt, &option)

	switch option {
	case "test-all":
		testAll()
	case "test-file":
		testFile()
	}
}

func testAll() {
	out, err := exec.Command("go", "test", "./test").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func testFile() {

}

func exit() {
	fmt.Print("Exiting")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
	os.Exit(1)
}

func main() {
	command := ""
	prompt := &survey.Select{
		Message: "Select a command:",
		Options: *&commands,
	}

	survey.AskOne(prompt, &command)

	switch command {
	case "build":
		build()
	case "deploy":
		deploy()
	case "test":
		test()
	case "exit":
		exit()
	default:
		fmt.Println("Option not suported")
	}
}

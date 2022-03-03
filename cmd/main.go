package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

var commands = []string{"build", "deploy", "tests", "exit"}

func build() {
	out, err := exec.Command("go", "build", "../pkg/main.go").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func deploy() {

}

func test() {
	option := ""
	prompt := &survey.Select{
		Message: "Select a test command:",
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
	files, err := IOReadDir("./test")

	if err != nil {
		log.Fatal(err)
	}

	file := ""
	prompt := &survey.Select{
		Message: "Select file to test:",
		Options: files,
	}

	survey.AskOne(prompt, &file)

	if &file == nil {
		log.Fatal("File not selected")
	}

	out, err := exec.Command("go", "test", "-v", "./test/"+file).Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func IOReadDir(dir string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir("./test/")

	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	return files, nil
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
	case "tests":
		test()
	case "exit":
		exit()
	default:
		fmt.Println("Option not suported")
	}
}

package main

// random editor

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func create(directory string) {
	files, _ := os.ReadDir(directory) // reading directory (.)
	newfilename := "1"
	if len(files) != 0 {
		lastfile := files[len(files)-1].Name()             // get last element of type []fs.DirEntry
		lastfilename, _, _ := strings.Cut(lastfile, ".")   // cutting text by separator (.)
		lastfilenameint, err := strconv.Atoi(lastfilename) // covert string to int
		if err != nil {
			fmt.Println("Non-int files are exist:", err)
			return
		}
		newfilename = fmt.Sprintf("%d", lastfilenameint+1) // convert int to string
	}
	newfile, _ := os.Create(directory + newfilename + ".go") // creating a file
	defer newfile.Close()                                    // close file
	abspath, _ := filepath.Abs(newfile.Name())
	cmd := exec.Command("mc", abspath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func search(directory string) {
	fmt.Printf("🔎︎ ")
	var search_query string
	fmt.Scanln(&search_query)
	clear := exec.Command("clear")
	clear.Stdout=os.Stdout
	clear.Run()
	grep := exec.Command("grep", "-rni", search_query, directory)
	grep.Stdout = os.Stdout
	grep.Run()
}

func main() {
	user, _ := user.Current()
	directory := fmt.Sprintf("/home/%s/.note/", user.Username)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, 0750)
	}
	fmt.Printf("Mode: Search(S) or Create(C) ")
	var mode_type_input string
	fmt.Scanln(&mode_type_input)
	if mode_type_input == "C" || mode_type_input == "c" {
		create(directory)
	} else if mode_type_input == "S" || mode_type_input == "s" {
		search(directory)
	} else {
		fmt.Println("Aborted. Cause: Wrong Key")
	}

	// fmt.Printf("%T\n",files) // get type
}

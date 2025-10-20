package main

// random editor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	files, _ := os.ReadDir(".") // reading directory (.)
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
	newfile, _ := os.Create(newfilename + ".go") // creating a file
	defer newfile.Close()                        // close file
	abspath, _ := filepath.Abs(newfile.Name())
	cmd := exec.Command("mc", abspath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	// fmt.Printf("%T\n",files) // get type
}

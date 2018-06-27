package main

import (
	"fmt"
	"io/ioutil"
	"os"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	// Load json data from file or stdin as appropriate
	var data []byte
	var err error
	if len(os.Args) <= 1 {
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = ioutil.ReadFile(os.Args[1])
	}

	// Quit on error
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Initialise termbox; pass off to browse()
	termbox.Init()
	termbox.SetInputMode(termbox.InputAlt)
	defer termbox.Close()
	browse(data)
}

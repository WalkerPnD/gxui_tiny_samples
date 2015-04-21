package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"

	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func fileText() string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file, err := os.Open(path + "/sample.txt")
	if err != nil {
		fmt.Println("No such file\nerror:", err)
		os.Exit(1)
	}
	defer file.Close()

	buffer := bufio.NewScanner(file)
	text := ""
	for buffer.Scan() {
		text += buffer.Text() + "\n"
	}
	return text
}

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(400, 225, "CreateTextBox")

	textbox := theme.CreateTextBox()
	textbox.SetMultiline(true)
	textbox.SetText(fileText())

	window.AddChild(textbox)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

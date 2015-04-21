package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	// Create label and set a word
	label1 := theme.CreateLabel()
	label1.SetText("Hello, world")

	window := theme.CreateWindow(380, 100, "Window for My Program")
	window.AddChild(label1)

	// when close
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

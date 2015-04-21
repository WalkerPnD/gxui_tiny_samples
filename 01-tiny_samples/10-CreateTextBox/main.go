package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(400, 225, "CreateTextBox")

	textbox := theme.CreateTextBox()
	textbox.SetText("fileText\nHello")

	window.AddChild(textbox)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

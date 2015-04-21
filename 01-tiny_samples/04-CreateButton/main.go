package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	button := theme.CreateButton()
	button.SetText("Yes!")

	window := theme.CreateWindow(380, 100, "Empty Window")
	window.AddChild(button)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

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
	// Criating a function for on click
	onClickFunc := func(gxui.MouseEvent) {
		if button.Text() == "Yes!" {
			button.SetText("No!")
		} else {
			button.SetText("Yes!")
		}
	}
	button.OnClick(onClickFunc)

	window := theme.CreateWindow(380, 100, "Empty Window")
	window.AddChild(button)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

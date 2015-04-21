package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	// flag to check the existance of the sub window
	subWindowIsOpen := false

	button := theme.CreateButton()
	button.SetText("Open subWindow")
	// fuction for the button
	onClickFunc := func(gxui.MouseEvent) {
		// check if sub window exists
		if subWindowIsOpen == false {
			subWindow := theme.CreateWindow(200, 60, "subWindow")
			// when close sub window put the flag back to false
			subWindow.OnClose(func() {
				subWindowIsOpen = false
			})
			subWindowIsOpen = true
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

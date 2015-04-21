package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	subWindowIsOpen := false

	button := theme.CreateButton()
	button.SetText("Open subWindow")
	onClickFunc := func(gxui.MouseEvent) {
		if subWindowIsOpen == false {
			subWindow := theme.CreateWindow(200, 60, "subWindow")
			subWindow.OnClose(func() {
				subWindowIsOpen = false
			})
			subWindowIsOpen = true
			return
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

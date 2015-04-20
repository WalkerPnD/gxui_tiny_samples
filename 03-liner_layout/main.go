package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	label1 := theme.CreateLabel()
	label1.SetText("Hello, world")
	label2 := theme.CreateLabel()
	label2.SetText("Hello, world again")

	layout := theme.CreateLinearLayout()
	layout.SetSizeMode(gxui.Fill)
	//layout.SetDirection(gxui.LeftToRight)
	layout.AddChild(label1)
	layout.AddChild(label2)

	window := theme.CreateWindow(380, 100, "Empty Window")
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

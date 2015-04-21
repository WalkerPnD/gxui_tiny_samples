package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(400, 225, "CreateSplitterLayout")

	// Labels for each Layout
	labelL := theme.CreateLabel()
	labelL.SetText("Label L")
	labelR := theme.CreateLabel()
	labelR.SetText("Label R")

	splitterLR := theme.CreateSplitterLayout()
	splitterLR.SetOrientation(gxui.Horizontal) // or "gxui.Vertical"
	splitterLR.AddChild(labelL)
	splitterLR.AddChild(labelR)

	window.AddChild(splitterLR)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(400, 225, "CreatePinelHolder")

	// Labels for Holder
	labelA := theme.CreateLabel()
	labelA.SetText("Label A")
	labelB := theme.CreateLabel()
	labelB.SetText("Label B")

	// CreatePinelHolder
	holder := theme.CreatePanelHolder()
	holder.AddPanel(labelA, "tab A")
	holder.AddPanel(labelB, "tab B")

	window.AddChild(holder)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

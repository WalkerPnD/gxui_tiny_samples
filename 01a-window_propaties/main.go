package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)

	window := theme.CreateWindow(380, 100, "Empty Window")

	// Back ground color
	color := gxui.Yellow // same with Color{1.0, 1.0, 0.0, 1.0}

	// Type and List are here: https://github.com/google/gxui/blob/master/color.go#L60
	brush := gxui.CreateBrush(color)
	window.SetBackgroundBrush(brush)

	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

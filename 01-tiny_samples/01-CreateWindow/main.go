package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	// Creating window using "themes.dark"
	theme := dark.CreateTheme(driver)

	// CreateWindow(WIDTH, HEIGHT, TITLE)
	window := theme.CreateWindow(380, 100, "Empty Window")

	// when close
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

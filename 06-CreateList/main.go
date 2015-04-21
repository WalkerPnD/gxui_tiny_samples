package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

// More Info:
// https://github.com/google/gxui/blob/master/mixins/list.go
// https://github.com/google/gxui/blob/master/themes/dark/list.go
// https://godoc.org/github.com/google/gxui#List

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(400, 225, "CreateList")

	items := []string{
		"Orange", "Lemon", "Apple",
	}

	// List needs adapter
	adapter := gxui.CreateDefaultAdapter()
	adapter.SetItems(items)
	list := theme.CreateList()
	list.SetAdapter(adapter)

	window.AddChild(list)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

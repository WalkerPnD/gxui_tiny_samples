package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

// More Info:
// https://godoc.org/github.com/google/gxui#DropDownList
// https://github.com/google/gxui/blob/master/mixins/drop_down_list.go
// https://github.com/google/gxui/blob/master/themes/dark/drop_down_list.go

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(400, 225, "CreateList")

	items := []string{
		"Orange", "Lemon", "Apple",
	}

	adapter := gxui.CreateDefaultAdapter()
	adapter.SetItems(items)

	// DropDownList needs overlay
	overlay := theme.CreateBubbleOverlay()
	ddlist := theme.CreateDropDownList()
	ddlist.SetAdapter(adapter)
	ddlist.SetBubbleOverlay(overlay)

	window.AddChild(ddlist)
	window.AddChild(overlay) // Overlay has to be added
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

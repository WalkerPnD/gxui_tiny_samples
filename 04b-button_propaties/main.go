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

	// Info: https://github.com/google/gxui/blob/master/themes/dark/button.go
	// func CreateButton(theme *Theme) gxui.Button {
	// 	b := &Button{}
	// 	b.Init(b, theme)
	// 	b.theme = theme
	// 	b.SetPadding(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	// 	b.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	// 	b.SetBackgroundBrush(b.theme.ButtonDefaultStyle.Brush)
	// 	b.SetBorderPen(b.theme.ButtonDefaultStyle.Pen)
	// 	b.OnMouseEnter(func(gxui.MouseEvent) { b.Redraw() })
	// 	b.OnMouseExit(func(gxui.MouseEvent) { b.Redraw() })
	// 	b.OnMouseDown(func(gxui.MouseEvent) { b.Redraw() })
	// 	b.OnMouseUp(func(gxui.MouseEvent) { b.Redraw() })
	// 	b.OnGainedFocus(b.Redraw)
	// 	b.OnLostFocus(b.Redraw)
	// 	return b
	// }

	window := theme.CreateWindow(380, 100, "Empty Window")
	window.AddChild(button)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

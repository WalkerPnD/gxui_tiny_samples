package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"

	"image"      // to decode image
	"image/draw" // to draw image
	_ "image/jpeg"
	"os" // to read file
	"path/filepath"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(400, 225, "CriateImage")

	path, _ := filepath.Abs(filepath.Dir(os.Args[0])) // Get current path
	file, _ := os.Open(path + "/cappuccino.jpg")      // "_" to ignore the error state
	defer file.Close()                                // register to close on exit

	img, _, _ := image.Decode(file)
	rgbaBuffer := image.NewRGBA(img.Bounds()) // img.Bound = image size information

	draw.Draw(rgbaBuffer, img.Bounds(), img, image.Point{0, 0}, draw.Src) // draw img to rgbaBuffer
	imgArea := theme.CreateImage()
	texture := driver.CreateTexture(rgbaBuffer, 1)
	imgArea.SetTexture(texture)

	window.AddChild(imgArea)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

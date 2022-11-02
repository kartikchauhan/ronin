package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Ronin",
		Bounds:    pixel.R(0, 0, 1024, 768),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	img, err := loadPicture("./assets/ronin.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(img, img.Bounds())

	win.Clear(colornames.Green)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

// func loadPicture(path string) (pixel.Picture, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pixel.PictureDataFromImage(img), nil
// }

func main() {
	pixelgl.Run(run)
}

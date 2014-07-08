package main

import (
    "code.google.com/p/go-tour/pic"
    "image/color"
    "image"
)

type Image struct{}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0,0,200,200)
}

func (img *Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x + y), uint8(x ^ y), uint8(x * y), 255}
}


func main() {
    m := &Image{}
    
    pic.ShowImage(m)
}

package main

import (
	"github.com/bonggeek/element"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"github.com/llgcode/draw2d"
	"fmt"
)

func createElementThumbNail(element element.Element) {
	fmt.Println("Generating image for", element)
	w := 75.0
	h := 75.0
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(image.White)
	gc.SetStrokeColor(image.Black)
	gc.SetLineWidth(4)

	// Draw a closed shape
	if !element.Real {
		gc.SetFillColor(color.RGBA{0xef, 0xef,0xef, 0xff})
	}

	gc.MoveTo(0, 0) // should always be called first for a new path
	gc.LineTo(w, 0)
	gc.LineTo(w, h)
	gc.LineTo(0, h)
	gc.LineTo(0, 0)
	gc.Close()
	gc.FillStroke()

	draw2d.SetFontFolder("/usr/share/fonts/truetype/freefont")
	//gc.SetFontData(draw2d.FontData{Name: "FreeSans"})
	gc.SetFillColor(color.RGBA{0, 0,0, 0xff})
	//gc.SetLineWidth(4)


	// Draw the main element Symbol
	str := element.Symbol
	gc.SetFontSize(20)
	l, _, r, _ := gc.GetStringBounds(str)
	strW := r - l
	x := (w - strW) / 2
	gc.FillStringAt(str, x, h / 2 + 5)

	// Draw the element atomic number
	gc.SetFontSize(10)
	atStr := fmt.Sprintf("%v", element.AtNumber)
	gc.FillStringAt(atStr, 5, 14)

	// Draw the element name
	str = element.Name
	gc.SetFontSize(8)

	l, _, r, _ = gc.GetStringBounds(str)
	strW = r - l
	x = (w - strW) / 2

	gc.FillStringAt(str, x, h - 10)

	fileName := fmt.Sprintf("%v.png", element.AtNumber)
	// Save to file
	draw2dimg.SaveToPngFile(fileName, dest)
	fmt.Println("Generated", fileName)
}

func main() {

	elements := element.GetElements().GetAllElements()
	for _, element := range elements{
		createElementThumbNail(element)
	}
}

package visualization

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"

	core "github.com/jirgl/quadient-events-no1/core"
	m "github.com/jirgl/quadient-events-no1/model"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var cellSize = 15
var fontColor = color.RGBA{60, 60, 60, 255}
var regularCellColor = color.RGBA{200, 200, 200, 255}
var visitedCellColor = color.RGBA{0, 200, 0, 128}
var pathCellColor = color.RGBA{200, 40, 40, 255}

func drawLabel(label string, x, y int, img *image.RGBA) {
	point := fixed.Point26_6{
		fixed.Int26_6(x * cellSize * 64),
		fixed.Int26_6((y + 1) * cellSize * 64),
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(fontColor),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func drawCell(x, y int, img *image.RGBA, at *core.ArrayTraveler) {
	node := at.GetNode(x, y)
	cost, _ := m.ParseNode(node.OriginData)

	isPath := core.IsPath(node.Position)
	wasVisited := core.WasVisited(node.Position)
	for relX := 0; relX < cellSize; relX++ {
		for relY := 0; relY < cellSize; relY++ {
			if isPath {
				img.Set(x*cellSize+relX, y*cellSize+relY, pathCellColor)
			} else if wasVisited {
				img.Set(x*cellSize+relX, y*cellSize+relY, visitedCellColor)
			} else {
				img.Set(x*cellSize+relX, y*cellSize+relY, regularCellColor)
			}
		}
	}

	drawLabel(strconv.Itoa(cost), x, y, img)
}

//Visualize func Flush result into image
func Visualize(filename string, at *core.ArrayTraveler) {
	size := at.DimSize * cellSize
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for x := 0; x < at.DimSize; x++ {
		for y := 0; y < at.DimSize; y++ {
			drawCell(x, y, img, at)
		}
	}

	file, _ := os.OpenFile(filename+".png", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	png.Encode(file, img)
}

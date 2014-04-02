package colorshow

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func createSwatch(col color.Color, size int) image.Image {
	log.Printf("Creating swatch for %v\n", col)
	bounds := image.Rect(0, 0, size, size)
	img := image.NewNRGBA(bounds)
	draw.Draw(img, img.Bounds(), image.NewUniform(col), img.Bounds().Min, draw.Src)
	return img
}

func drawSwatches(cols []color.Color, size, numCols int) image.Image {
	log.Println("Creating swatches...")
	destBounds := image.Rect(0, 0, size*numCols, size*(len(cols)/numCols))
	destImg := image.NewNRGBA(destBounds)
	for i, c := range cols {
		x := (i % numCols) * size
		y := (i / numCols) * size
		img := createSwatch(c, size)
		drawBounds := image.Rect(x, y, x+size, y+size)
		draw.Draw(destImg, drawBounds, img, img.Bounds().Min, draw.Src)
	}
	return destImg
}

func DisplaySwatches(cols []color.Color) {
	width := 8
	if len(cols) < width {
		width = len(cols)
	}
	DisplayImage(drawSwatches(cols, 100, width))
}

func WriteSwatches(cols []color.Color, filename string) {
	fd, err := os.Create(filename)
	if err != nil {
		log.Println("error:", err)
		return
	}
	defer fd.Close()
	png.Encode(fd, drawSwatches(cols, 100, 8))
}

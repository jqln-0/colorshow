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

	// SDL will freak out if we leave some pixels unset, so make everything black initially.
	// TODO: Instead of black do like a pretty squares pattern.
	black := color.NRGBA{0, 0, 0, 255}
	draw.Draw(destImg, destBounds, image.NewUniform(black), destBounds.Min, draw.Src)

	for i, c := range cols {
		x := (i % numCols) * size
		y := (i / numCols) * size
		img := createSwatch(c, size)
		drawBounds := image.Rect(x, y, x+size, y+size)
		draw.Draw(destImg, drawBounds, img, img.Bounds().Min, draw.Src)
	}
	return destImg
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// badFactorize attempts to find a factor of i that makes for a square-ish grid. It is bad
// because the algorithm used is really bad and if I wasn't sick I'd make it better.
func badFactorize(i int) int {
	bestCandidate := 1
	// 'Score' here is the difference between the width and height. Low is good.
	bestScore := i
	for candidate := 1; candidate <= i/2; candidate++ {
		if i%candidate != 0 {
			continue
		}
		width, height := candidate, i/candidate
		score := abs(width - height)
		if score < bestScore {
			bestCandidate = candidate
			bestScore = score
		}
	}
	return i / bestCandidate
}

func DisplaySwatches(cols []color.Color) {
	width := badFactorize(len(cols))
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

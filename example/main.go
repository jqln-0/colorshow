package main

import (
	"image/color"

	"github.com/jqln-0/colorshow"
)

func main() {
	example := []color.NRGBA{
		{255, 255, 255, 255},
		{255, 255, 0, 255},
		{255, 0, 255, 255},
		{255, 0, 0, 255},
		{0, 255, 255, 255},
		{0, 255, 0, 255},
		{0, 0, 255, 255},
		{0, 0, 0, 255},
		{128, 128, 128, 255},
		{128, 128, 0, 255},
		{128, 0, 128, 255},
		{128, 0, 0, 255},
		{0, 128, 128, 255},
		{0, 128, 0, 255},
		{0, 0, 128, 255},
		{0, 0, 0, 255},
	}

	exampleInterface := make([]color.Color, 0, len(example))
	for _, c := range example {
		exampleInterface = append(exampleInterface, color.Color(c))
	}

	colorshow.DisplaySwatches(exampleInterface)
}

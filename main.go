package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

func main() {
	const S = 50

	rand.Seed(time.Now().UnixNano())

	for _, c := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		dc := gg.NewContext(S, S)
		dc.SetRGB(float64(randomInt(0, 255)), float64(randomInt(0, 255)), float64(randomInt(0, 255))) //to set the background color
		dc.Clear()
		if err := dc.LoadFontFace("PlayfairDisplay-VariableFont_wght.ttf", 25); err != nil {
			panic(err)
		}
		dc.SetRGB(float64(randomInt(0, 255)), float64(randomInt(0, 255)), float64(randomInt(0, 255))) //to set the letter color
		dc.DrawStringAnchored(string(c), S/2, S/2, 0.5, 0.5)
		dc.SavePNG(string(c) + ".png")
		dc.Clear()
	}
}

// Returns an int >= min, <= max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

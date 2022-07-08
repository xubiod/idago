package idago

import (
	"image"
	"image/color"
)

func Encode(data []byte, w, h int) *image.NRGBA {
	r := image.Rect(0, 0, w, h)
	m := image.NewNRGBA(r)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			go func(x int, y int, w int) {
				idx := (x + (y * w))
				if idx < len(data) {
					m.SetNRGBA(x, y, getColor(data[idx]))
				}
			}(x, y, w)
		}
	}
	return m
}

func getColors(data byte) (color.NRGBA, color.NRGBA) {
	cs := [16]color.NRGBA{
		{43, 45, 66, 255},    // 0
		{255, 173, 173, 255}, // 1
		{255, 214, 165, 255}, // 2
		{253, 255, 182, 255}, // 3
		{202, 255, 191, 255}, // 4
		{155, 246, 255, 255}, // 5
		{160, 196, 255, 255}, // 6
		{189, 178, 255, 255}, // 7
		{255, 198, 255, 255}, // 8
		{255, 89, 94, 255},   // 9
		{255, 202, 58, 255},  // A
		{138, 201, 38, 255},  // B
		{25, 130, 196, 255},  // C
		{106, 76, 147, 255},  // D
		{141, 153, 174, 255}, // E
		{237, 242, 244, 255}, // F
	}

	th := (data >> 4) & 0x0F
	bh := data & 0x0F

	lhr := cs[th]
	rhr := cs[bh]

	return lhr, rhr
}

func getColor(data byte) color.NRGBA {
	lhr, rhr := getColors(data)

	var c color.NRGBA

	c.R = lhr.R&0x0F | rhr.R&0xF0
	c.G = lhr.G&0x0F | rhr.G&0xF0
	c.B = lhr.B&0x0F | rhr.B&0xF0
	c.A = 255

	return c
}

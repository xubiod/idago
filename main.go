package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"
	idago "xubiod/idago/elements"
)

func main() {
	// generateByteencodedImage("out", Example0{})
}

type Example0 struct{}

func (Example0) GenerateByteplane() ([]byte, image.Point) {
	var byteplane0 []byte
	var byteplane1 []byte
	count := 16384 * 32 // Now only *approximately* what you want. Rounds to nearest square with 8x8 cells n shit
	var x, y int = -1, -1
	x++
	y++

	w := int(math.Floor(math.Sqrt(float64(count))))
	h := int(math.Ceil(math.Sqrt(float64(count))))
	w = int(math.Ceil(float64(w)/8) * 8)
	h = int(math.Ceil(float64(h)/8) * 8)
	count = w * h

	for idx := 0; idx < count; idx++ {
		x = (idx % w)
		y = (int(idx/w) % h)
		x = x + (y / 4)
		y = y + (x / 8)
		byteplane0 = append(byteplane0, byte(int(math.Abs(float64(math.Sin(float64((x/8)+y*16))*0xFF)))&0xFF))
		byteplane1 = append(byteplane1, byteplane0[idx])
	}

	rm := 360.0
	for r := 0.0; r < rm; r++ {
		for j := 1.0; j < 3.0; j += 0.05 {
			x = int(math.Sin(r/rm*2.0*math.Pi)*float64(w)/j) + w/2
			y = int(math.Cos(r/rm*2.0*math.Pi)*float64(h)/j) + h/2
			if x > 0 && y > 0 && x < w && y < h {
				idx := (x + (y * w))
				byteplane0[idx] = 0x00
			}
		}
	}

	for r := 225.0; r < rm+135; r += 0.2 {
		for j := 3.7; j < 8.0; j += 0.01 {
			x = int(math.Sin(r/rm*2.0*math.Pi)*float64(w)/j) + w/2
			y = int(math.Cos(r/rm*2.0*math.Pi)*float64(h)/j) + h/2
			idx := (x + (y * w))
			byteplane0[idx] = 0x00
		}
	}

	for lx := w/2 - 20; lx < w/2+20; lx++ {
		for ly := 0; ly < h/2+4; ly++ {
			idx := (lx + (ly * w))
			byteplane0[idx] = 0x00
		}
	}

	for idx := range byteplane0 {
		byteplane0[idx] ^= byteplane1[idx]
	}

	return byteplane0, image.Point{w, h}
}

type byteplaneGenerator interface {
	GenerateByteplane() ([]byte, image.Point)
}

func generateByteencodedImage(fn string, gen byteplaneGenerator) {
	byteplane, point := gen.GenerateByteplane()

	w := point.X
	h := point.Y

	img := idago.Encode(byteplane, w, h)

	f, err := os.Create(fn + ".png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	f, err = os.Create(fn + ".bin")
	if err != nil {
		log.Fatal(err)
	}

	n, err := f.Write(byteplane)
	if err != nil {
		log.Fatal(n, err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

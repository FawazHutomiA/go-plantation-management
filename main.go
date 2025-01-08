package main

import (
	"fmt"
	"math"
)

// Struktur untuk menyimpan informasi plot
type Plot struct {
	X      int // Posisi X (Kolom)
	Y      int // Posisi Y (Baris)
	Height int // Tinggi pohon (meter)
}

func calculateDistance(length, width int, trees []Plot) int {
	// Inisialisasi
	grid := make(map[int]map[int]int)
	for _, t := range trees {
		if grid[t.Y] == nil {
			grid[t.Y] = make(map[int]int)
		}
		grid[t.Y][t.X] = t.Height
	}

	// Jarak Horizontal
	horizontalDistance := 0
	for y := 1; y <= width; y++ {
		horizontalDistance += (length - 1) * 10
	}

	// Jarak Vertikal
	verticalDistance := 0
	currentHeight := 0

	// Naik dari tanah ke plot pertama
	verticalDistance += 1

	for y := 1; y <= width; y++ {
		for x := 1; x <= length; x++ {
			height := grid[y][x]
			// Naik/turun ke ketinggian plot saat ini
			verticalDistance += int(math.Abs(float64(height - currentHeight)))
			currentHeight = height
		}
		// Turun ke ketinggian plot di awal baris berikutnya
		if y < width {
			nextHeight := grid[y+1][1]
			verticalDistance += int(math.Abs(float64(nextHeight - currentHeight)))
			currentHeight = nextHeight
		}
	}

	// Turun kembali ke tanah
	verticalDistance += currentHeight
	fmt.Println("Total vertical distance:", verticalDistance, "meter")
	fmt.Println("Total horizontal distance:", horizontalDistance, "meter")

	// Total Jarak
	return horizontalDistance + verticalDistance
}

func main() {
	// Contoh input dimensi perkebunan
	length := 6 // Panjang (60 meter / 10)
	width := 3  // Lebar (30 meter / 10)

	// Input dari pengguna untuk tinggi pohon
	trees := []Plot{
		{X: 3, Y: 1, Height: 5},
		{X: 3, Y: 2, Height: 10},
		{X: 4, Y: 2, Height: 15},
		{X: 6, Y: 2, Height: 10},
		{X: 5, Y: 3, Height: 4},
	}

	totalDistance := calculateDistance(length, width, trees)
	fmt.Printf("Total Jarak Drone: %d meter\n", totalDistance)
}

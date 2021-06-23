package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/haunt98/dpi-go"
)

func main() {
	if len(os.Args[1:]) != 3 {
		fmt.Println("Example: go run ./cmd/dpi/main.go [width] [height] [diagonal]")
		return
	}

	width, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	height, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	diagonal, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := dpi.NewResolution(width, height)

	fmt.Printf("DPI: %.2f\n", dpi.CalculateDPI(r, diagonal))

	wRatio, hRatio := r.CalculateAspectRatio()
	fmt.Printf("Aspect Ratio: %d:%d\n", wRatio, hRatio)
}

package resources

import (
	"fmt"
	"image"
	"math"
	"os"

	// Seems to be needed for the Pixel lib
	_ "image/png"

	errorUtil "github.com/pkg/errors"

	"github.com/faiface/pixel"
)

// LoadPNGPicture loads a PNG file
func LoadPNGPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errorUtil.Wrap(err, "LoadPNGPicture() os.Open(path)")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, errorUtil.Wrap(err, "LoadPNGPicture() image.Decode(file)")
	}
	return pixel.PictureDataFromImage(img), nil
}

// CalculateDistance between two points
func CalculateDistance(x1 float64, y1 float64, x2 float64, y2 float64) bool {
	distance := math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
	fmt.Printf("Distance: %f\n", distance)
	return distance > 550
}

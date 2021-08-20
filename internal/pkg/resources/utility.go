/*                                  
    Cheemswave - Side Scroller Game 
    Copyright (C) 2021  David Kirwan                                                                                                                                                                                                          
                                    
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
                                    
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
                                    
    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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

package resources

import (
	"strings"
	"testing"
)

// ErrorContains checks if the error message in out contains the text in
// want.
//
// This is safe when out is nil. Use an empty string for want if you want to
// test that err is nil.
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

func TestLoadPNGPicture(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test for file which exists",
			args: args{
				path: "../../../assets/images/trees.png",
			},
			want: "",
		},
		{
			name: "test with file that does not exist",
			args: args{
				path: "./assets/images/queen.png",
			},
			want: "LoadPNGPicture() os.Open(path): open ./assets/images/queen.png: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPNGPicture(tt.args.path)
			if !ErrorContains(err, tt.want) {
				t.Errorf("LoadPNGPicture(): %v, want %v", err, tt.want)
			}
		})
	}
}

func TestCalculateDistance(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test point is greater than 550",
			args: args{
				x1: 0.0,
				y1: 0.0,
				x2: 551,
				y2: 0.0,
			},
			want: true,
		},
		{
			name: "test point is less than 550",
			args: args{
				x1: 0.0,
				y1: 0.0,
				x2: 549,
				y2: 0.0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDistance(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("CalculateDistance(): %v, want %v", got, tt.want)
			}
		})
	}
}

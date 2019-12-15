package main

import (
	"fmt"
	"testing"
)

func Test_getFuelRequirements(t *testing.T) {
	tests := []struct {
		name string
		mass []int
		want int
	}{
		{
			name: "mass is 12+12, expect fuel: 4",
			mass: []int{12, 12},
			want: 4,
		},
		{
			name: "mass is 14+14, expect fuel: 4",
			mass: []int{14, 14},
			want: 4,
		},
		{
			name: "mass is 1969+1969+1969, expect fuel: 1962",
			mass: []int{1969, 1969, 1969},
			want: 1962,
		},
		{
			name: "mass is 100756+100756+100756, expect fuel: 100749",
			mass: []int{100756, 100756, 100756},
			want: 100749,
		},
	}
	for _, tt := range tests {
		actual := getFuelRequirements(tt.mass)

		if actual != tt.want {
			t.Errorf("getFuelRequirements() = %v, want %v", actual, tt.want)
		}
	}
}

func Test_getFuelRequirement(t *testing.T) {
	tests := []struct {
		name string
		mass int
		want int
	}{
		{
			name: "mass is 12, expect fuel: 2",
			mass: 12,
			want: 2,
		},
		{
			name: "mass is 14, expect fuel: 2",
			mass: 14,
			want: 2,
		},
		{
			name: "mass is 1969, expect fuel: 654",
			mass: 1969,
			want: 654,
		},
		{
			name: "mass is 100756, expect fuel: 33583",
			mass: 100756,
			want: 33583,
		},
	}
	for _, tt := range tests {
		actual := getFuelRequirement(tt.mass)

		if actual != tt.want {
			t.Errorf("getFuelRequirement() = %v, want %v", actual, tt.want)
		}
	}
}

func Test_solvePuzzle(t *testing.T) {
	input := []int{102480, 121446, 118935, 54155, 102510, 142419, 73274, 57571, 123916, 99176, 143124, 141318, 72224,
		145479, 97027, 126427, 94990, 100521, 105589, 123009, 77143, 142861, 92366, 66478, 102195, 128373, 128447,
		120178, 99122, 98671, 89541, 125720, 107984, 126544, 145231, 110241, 123926, 72793, 76705, 128338, 74262, 68845,
		65297, 112536, 59892, 57115, 73230, 80569, 146118, 108843, 59221, 140492, 122616, 140652, 64404, 99782, 104375,
		86926, 143145, 114969, 108948, 77236, 143655, 71406, 97588, 64892, 105345, 104393, 93442, 54525, 94116, 123606,
		106813, 59904, 149253, 81620, 80892, 66309, 142604, 97984, 79743, 79448, 123756, 64927, 139703, 71448, 135964,
		86083, 94767, 116856, 73786, 141083, 122581, 82239, 122282, 96092, 80029, 52957, 72062, 52124}

	fmt.Println(getFuelRequirements(input))
}

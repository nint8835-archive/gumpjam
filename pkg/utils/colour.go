package utils

import "image/color"

func GetContrastingTextColour(background color.Color) color.Color {
	r, g, b, _ := background.RGBA()

	// https://stackoverflow.com/a/3943023
	yiq := ((r * 299) + (g * 587) + (b * 114)) / 1000

	if yiq >= 128 {
		return color.Black
	}

	return color.White
}

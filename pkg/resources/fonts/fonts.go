package fonts

import (
	_ "embed"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed MINGLIU.ttf
var mingLiu []byte

var MingLiU = make(map[int]font.Face)

const SizeSmall = 12

var fontSizes = []int{
	SizeSmall,
}

func init() {
	var err error

	mingLiuOt, err := opentype.Parse(mingLiu)
	if err != nil {
		panic(err)
	}

	for _, size := range fontSizes {
		MingLiU[size], err = opentype.NewFace(mingLiuOt, &opentype.FaceOptions{
			Size:    float64(size),
			DPI:     72,
			Hinting: font.HintingNone,
		})
		if err != nil {
			panic(err)
		}
	}
}

// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package text

import (
	"image/color"
	"strings"

	"github.com/jeefies/vim-survival/log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	arcadeFontBaseSize = 4
	fontMaxScale = 16
)

var (
	arcadeFonts map[int]font.Face
)

func getArcadeFonts(scale int) font.Face {
	if arcadeFonts == nil {
		log.Logger.Printf("arcadeFonts == nil")
		tt, err := opentype.Parse(fonts.PressStart2P_ttf)
		if err != nil {
			log.Logger.Fatal(err)
		}

		arcadeFonts = map[int]font.Face{}
		for i := 1; i <= fontMaxScale; i++ {
			const dpi = 72
			arcadeFonts[i], err = opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    float64(arcadeFontBaseSize * i),
				DPI:     dpi,
				Hinting: font.HintingFull,
			})
			if err != nil {
				log.Logger.Fatal(err)
			}
		}
	}
	return arcadeFonts[scale]
}

func TextWidth(str string) int {
	maxW := 0
	for _, line := range strings.Split(str, "\n") {
		b, _ := font.BoundString(getArcadeFonts(1), line)
		w := (b.Max.X - b.Min.X).Ceil()
		if maxW < w {
			maxW = w
		}
	}
	return maxW
}

func TextSize(str string) (int, int) {
	maxW := 0
	lineCount := 0
	for _, line := range strings.Split(str, "\n") {
		lineCount++
		b, _ := font.BoundString(getArcadeFonts(1), line)
		w := (b.Max.X - b.Min.X).Ceil()
		if maxW < w {
			maxW = w
		}
	}
	return maxW, arcadeFontBaseSize * lineCount
}

var (
	shadowColor = color.NRGBA{0, 0, 0, 0x80}
)

func DrawText(scr * ebiten.Image, str string, x, y, scale int, clr color.Color) {
	offsetY := arcadeFontBaseSize * scale
	for _, line := range strings.Split(str, "\n") {
		y += offsetY
		log.Logger.Printf("arcade fonts: %p", getArcadeFonts(scale))
		text.Draw(scr, line, getArcadeFonts(scale), x, y, clr)
	}
}

func DrawTextWithCenter(scr *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := TextWidth(str) * scale
	x += (width - w) / 2
	DrawText(scr, str, x, y, scale, clr)
}

func DrawTextWithRight(scr *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := TextWidth(str) * scale
	x += width - w
	DrawText(scr, str, x, y, scale, clr)
}

func DrawTextCentered(scr *ebiten.Image, str string, x, y, scale int, clr color.Color) {
	w := TextWidth(str) * scale
	x -= w / 2
	DrawText(scr, str, x, y, scale, clr)
}

func DrawTextRight(scr *ebiten.Image, str string, x, y, scale int, clr color.Color) {
	w := TextWidth(str) * scale
	x -= w
	DrawText(scr, str, x, y, scale, clr)
}

func DrawTextWithShadow(rt *ebiten.Image, str string, x, y, scale int, clr color.Color) {
	offsetY := arcadeFontBaseSize * scale
	for _, line := range strings.Split(str, "\n") {
		y += offsetY
		text.Draw(rt, line, getArcadeFonts(scale), x+1, y+1, shadowColor)
		text.Draw(rt, line, getArcadeFonts(scale), x, y, clr)
	}
}

func DrawTextWithShadowCenter(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := TextWidth(str) * scale
	x += (width - w) / 2
	DrawTextWithShadow(rt, str, x, y, scale, clr)
}

func DrawTextWithShadowRight(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := TextWidth(str) * scale
	x += width - w
	DrawTextWithShadow(rt, str, x, y, scale, clr)
}

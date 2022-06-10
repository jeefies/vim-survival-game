package scenes

import (
	"math"
	"image/color"

	"github.com/jeefies/vim-survival/log"
	"github.com/jeefies/vim-survival/text"
	"github.com/jeefies/vim-survival/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type TitleScene struct {
	titleRotateAngle float64
	titleScale float64
	titleImage * ebiten.Image
}

var (
	Title = "Vim Survival"
	TitleFontScale = 10
	SubTitle = "Test Edition"
	SubTitleFontScale = 4
	TitleRotateMax = 4.0
	TitleRotateMin = -4.0
	TitleRotateSpeed = 0.07
	TitleScaleMax = 2.3
	TitleScaleMin = 2.0
	TitleScaleSpeed = 0.003

	TitleColor = color.RGBA{0, 0, 0, 0xff}
	SubTitleColor = color.NRGBA{0, 0, 0, 0xA0}

	WHITE = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

func CalcTheta(angle float64) float64 {
	return math.Pi * angle / 180
}

func (ts * TitleScene) StartTransition() {
}

func (ts * TitleScene) EndTransition() {
}

func (ts * TitleScene) Update(in * input.Input) error {
	if ts.titleScale == 0 {
		ts.titleScale = TitleScaleMin
	}
	ts.titleRotateAngle += TitleRotateSpeed
	if (ts.titleRotateAngle >= TitleRotateMax) || (ts.titleRotateAngle <= TitleRotateMin) {
		TitleRotateSpeed = -TitleRotateSpeed
	}

	ts.titleScale += TitleScaleSpeed
	if (ts.titleScale >= TitleScaleMax) || (ts.titleScale <= TitleScaleMin) {
		TitleScaleSpeed = -TitleScaleSpeed
	}

	return nil
}

func (ts * TitleScene) Draw(screen * ebiten.Image) {
	screen.Fill(WHITE)

	if ts.titleImage == nil {
		titleWidth, titleHeight := text.TextSize(Title)
		// font size 4 * 10, scale 10
		titleWidth, titleHeight = titleWidth * TitleFontScale, titleHeight * TitleFontScale

		subtitleWidth, subtitleHeight := text.TextSize(SubTitle)
		subtitleWidth, subtitleHeight = subtitleWidth * SubTitleFontScale, subtitleHeight * SubTitleFontScale

		totalHeight := titleHeight + subtitleHeight

		ts.titleImage = ebiten.NewImage(titleWidth, totalHeight)

		log.Logger.Printf("ts.titleImage = %p, Title is %s", ts.titleImage, Title)

		text.DrawText(ts.titleImage, Title, 0, 0, TitleFontScale, TitleColor)
		text.DrawTextRight(ts.titleImage, SubTitle, titleWidth * 9 / 10, titleHeight, SubTitleFontScale, SubTitleColor)
	}

	op := &ebiten.DrawImageOptions{}
	tisize := ts.titleImage.Bounds()
	width, height := tisize.Max.X - tisize.Min.X, tisize.Max.Y - tisize.Min.Y
	log.Logger.Printf("tsi w, h = %d %d", width, height)

	screenWidth, screenHeight := ebiten.ScreenSizeInFullscreen()

	// op.GeoM.Translate(float64(screenWidth - width) / 10, float64(screenHeight - height) / 10)
	op.GeoM.Translate(float64(screenHeight) / 10, float64(screenHeight) / 10)

	op.GeoM.Scale(ts.titleScale, ts.titleScale)
	op.GeoM.Rotate(CalcTheta(ts.titleRotateAngle))
	log.Logger.Printf("title Scale %f, rotate angle %f", ts.titleScale, ts.titleRotateAngle)

	screen.DrawImage(ts.titleImage, op)

	startInfo := "type :help for tutor\npress space to start"
	text.DrawTextCentered(screen, startInfo, screenWidth / 2, screenHeight * 4 / 5, SubTitleFontScale, SubTitleColor)
}

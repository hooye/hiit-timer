package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/sirupsen/logrus"
)

func main() {
	hiitTimerApp := app.NewWithID("com.hiitTimer")
	oneWindow := hiitTimerApp.NewWindow("Hiit Timer")
	circle := canvas.NewCircle(color.RGBA{0x00, 0xff, 0x00, 0xff})
	circle.StrokeColor = color.Gray{0x99}
	circle.StrokeWidth = 5

	text := canvas.NewText("100", color.Black)
	text.TextSize = 24
	text.Alignment = fyne.TextAlignCenter

	logrus.Info("start Hiit Timer")

	oneWindow.Resize(fyne.NewSize(500, 500))
	oneWindow.ShowAndRun()

	go func() {
		for i := 100; i >= 0; i-- {
			oneWindow.SetContent(canvas.NewText(strconv.Itoa(i), color.Black))
		}
	}()

	logrus.Info("end Hiit Timer")
}

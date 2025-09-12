package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "strconv"
)

var count = 0

func add(num int) string {
    count = count + num
    return strconv.Itoa(count)
}

func main() {
    myApp := app.New()
    window := myApp.NewWindow("Counter App")

    label := widget.NewLabel("0")

    plusonebutton := widget.NewButton("+ 1", func() {
        label.SetText(add(1))
    })
    plusfivebutton := widget.NewButton("+ 5", func() {
        label.SetText(add(5))
    })
    minonebutton := widget.NewButton("- 1", func() {
        label.SetText(add(-1))
    })
    minfivebutton := widget.NewButton("- 5", func() {
        label.SetText(add(-5))
    })

    centered := container.NewCenter(label)
    content := container.NewVBox(plusonebutton, plusfivebutton, minonebutton, minfivebutton)

    layout := container.NewVBox(centered, content)

    window.Resize(fyne.NewSize(300, 300))
    window.SetContent(layout)
    window.ShowAndRun()
}
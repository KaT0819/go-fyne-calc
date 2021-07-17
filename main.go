package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// cdata is data structure.
type cdata struct {
	mem int
	cal string
	flg bool
}

// 数字ボタンの生成・配置
func createNumButtons(f func(v int, data *cdata, l *widget.Label), data *cdata, l *widget.Label) *fyne.Container {
	c := container.New(
		layout.NewGridLayout(3),
		widget.NewButton(strconv.Itoa(7), func() { f(7, data, l) }),
		widget.NewButton(strconv.Itoa(8), func() { f(8, data, l) }),
		widget.NewButton(strconv.Itoa(9), func() { f(9, data, l) }),
		widget.NewButton(strconv.Itoa(4), func() { f(4, data, l) }),
		widget.NewButton(strconv.Itoa(5), func() { f(5, data, l) }),
		widget.NewButton(strconv.Itoa(6), func() { f(6, data, l) }),
		widget.NewButton(strconv.Itoa(1), func() { f(1, data, l) }),
		widget.NewButton(strconv.Itoa(2), func() { f(2, data, l) }),
		widget.NewButton(strconv.Itoa(3), func() { f(3, data, l) }),
		widget.NewButton(strconv.Itoa(0), func() { f(0, data, l) }),
	)
	return c
}

// 計算ボタンの生成・配置
func createCalcButtons(f func(c string, data *cdata, l *widget.Label), data *cdata, l *widget.Label) *fyne.Container {
	c := container.New(
		layout.NewGridLayout(1),
		widget.NewButton("CL", func() { f("CL", data, l) }),
		widget.NewButton("/", func() { f("/", data, l) }),
		widget.NewButton("*", func() { f("*", data, l) }),
		widget.NewButton("+", func() { f("+", data, l) }),
		widget.NewButton("-", func() { f("-", data, l) }),
	)
	return c
}

// main function.
func main() {
	a := app.New()
	w := a.NewWindow("Calc")
	w.SetFixedSize(true)
	l := widget.NewLabel("0")
	l.Alignment = fyne.TextAlignTrailing

	data := cdata{
		mem: 0,
		cal: "",
		flg: false,
	}

	k := createNumButtons(pushNum, &data, l)
	c := createCalcButtons(pushCalc, &data, l)
	e := widget.NewButton("Enter", func() { pushEnter(&data, l) })

	w.SetContent(
		container.New(
			layout.NewBorderLayout(
				l, e, nil, c,
			),
			l, e, k, c,
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

// 計算処理
func calc(n int, data *cdata, l *widget.Label) {
	switch data.cal {
	case "":
		data.mem = n
	case "+":
		data.mem += n
	case "-":
		data.mem -= n
	case "*":
		data.mem *= n
	case "/":
		data.mem /= n
	}
	l.SetText(strconv.Itoa(data.mem))
	data.flg = true
}

func pushNum(v int, data *cdata, l *widget.Label) {
	s := l.Text
	if data.flg {
		s = "0"
		data.flg = false
	}
	s += strconv.Itoa(v)
	n, err := strconv.Atoi(s)
	if err == nil {
		l.SetText(strconv.Itoa(n))
	}
}

func pushCalc(c string, data *cdata, l *widget.Label) {
	if c == "CL" {
		l.SetText("0")
		data.mem = 0
		data.flg = false
		data.cal = ""
		return
	}
	n, er := strconv.Atoi(l.Text)
	if er != nil {
		return
	}
	calc(n, data, l)

	data.cal = c
}

func pushEnter(data *cdata, l *widget.Label) {
	n, er := strconv.Atoi(l.Text)
	if er != nil {
		return
	}
	calc(n, data, l)
	data.cal = ""
}

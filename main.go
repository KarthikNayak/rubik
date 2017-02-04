package main

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	red     = 0
	green   = 1
	yellow  = 2
	blue    = 3
	magenta = 4
	cyan    = 5
)

type cube struct {
	f [3][3]int
	l [3][3]int
	r [3][3]int
	u [3][3]int
	d [3][3]int
	b [3][3]int
}

func initSide(s *[3][3]int, clr int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s[i][j] = clr
		}
	}
}

func (c *cube) init() {
	initSide(&c.f, red)
	initSide(&c.l, green)
	initSide(&c.r, yellow)
	initSide(&c.u, blue)
	initSide(&c.b, magenta)
	initSide(&c.d, cyan)
}

func colorSq() func(clr int) string {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	return func(clr int) string {
		switch clr {
		case 0:
			return red("■")
		case 1:
			return green("■")
		case 2:
			return yellow("■")
		case 3:
			return blue("■")
		case 4:
			return magenta("■")
		case 5:
			return cyan("■")
		}
		return "■"
	}
}

func (c *cube) printCube() {
	clr := colorSq()

	fmt.Println("        ---------")
	for i := 0; i < 3; i++ {
		fmt.Printf("        | %s %s %s |\n", clr(c.u[i][0]),
			clr(c.u[i][1]), clr(c.u[i][2]))
	}
	fmt.Println("---------------------------------")

	for i := 0; i < 3; i++ {
		fmt.Printf("| %s %s %s |", clr(c.l[i][0]), clr(c.l[i][1]), clr(c.l[i][2]))
		fmt.Printf(" %s %s %s |", clr(c.f[i][0]), clr(c.f[i][1]), clr(c.f[i][2]))
		fmt.Printf(" %s %s %s |", clr(c.r[i][0]), clr(c.r[i][1]), clr(c.r[i][2]))
		fmt.Printf(" %s %s %s |", clr(c.b[i][0]), clr(c.b[i][1]), clr(c.b[i][2]))
		fmt.Println()
	}
	fmt.Println("---------------------------------")

	for i := 0; i < 3; i++ {
		fmt.Printf("        | %s %s %s |\n", clr(c.d[i][0]),
			clr(c.d[i][1]), clr(c.d[i][2]))
	}
	fmt.Println("        ---------")
}

func swap(a, b *int) {
	tmp := *a
	*a, *b = *b, tmp
}

func swapR(a, b *[3]int) {
	for i := 0; i < 3; i++ {
		swap(&a[i], &b[i])
	}
}

func transpose(a *[3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < i; j++ {
			if i != j {
				swap(&a[i][j], &a[j][i])
			}
		}
	}
}

func rotateHelper(a, b *[3][3]int, inv bool) {
	for i := 0; i < 3; i++ {
		if inv != true {
			swap(&a[0][i], &b[2][i])
		} else {
			swap(&a[i][0], &b[i][2])
		}
	}
}

func (c *cube) rotateRow(row int, inv bool) {
	if inv != true {
		swapR(&c.b[row], &c.l[row])
		swapR(&c.r[row], &c.b[row])
		swapR(&c.f[row], &c.r[row])
	} else {
		swapR(&c.f[row], &c.r[row])
		swapR(&c.r[row], &c.b[row])
		swapR(&c.b[row], &c.l[row])
	}

	switch row {
	case 0:
		transpose(&c.u)
		rotateHelper(&c.u, &c.u, false || inv)
	case 2:
		transpose(&c.d)
		rotateHelper(&c.d, &c.d, true != inv)
	}
}

func swapC(a, b *[3][3]int, col1, col2 int) {
	for i := 0; i < 3; i++ {
		swap(&a[i][col1], &b[i][col2])
	}
}

func (c *cube) rotateCol(col int, inv bool) {
	if inv != true {
		swapC(&c.f, &c.d, col, col)
		swapC(&c.d, &c.b, col, 2-col)
		swapC(&c.b, &c.u, 2-col, col)
	} else {
		swapC(&c.b, &c.u, 2-col, col)
		swapC(&c.d, &c.b, col, 2-col)
		swapC(&c.f, &c.d, col, col)
	}

	switch col {
	case 0:
		transpose(&c.l)
		rotateHelper(&c.l, &c.l, false || inv)
	case 2:
		transpose(&c.r)
		rotateHelper(&c.r, &c.r, true != inv)
	}
}

func main() {
	var c cube
	c.init()
	c.l[0][2] = red
	c.printCube()
	c.rotateCol(0, true)
	c.printCube()
}

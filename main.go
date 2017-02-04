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

func (c *cube) rotateRow(row int) {
	swapR(&c.b[row], &c.l[row])
	swapR(&c.r[row], &c.b[row])
	swapR(&c.f[row], &c.r[row])

	switch row {
	case 0:
		transpose(&c.u)
		for i := 0; i < 3; i++ {
			swap(&c.u[0][i], &c.u[2][i])
		}
	case 2:
		transpose(&c.d)
		for i := 0; i < 3; i++ {
			swap(&c.d[i][0], &c.d[i][2])
		}
	}
}

func main() {
	var c cube
	c.init()
	c.printCube()
}

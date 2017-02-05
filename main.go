package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

const (
	red     = 0
	white   = 1
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
	initSide(&c.l, white)
	initSide(&c.r, yellow)
	initSide(&c.u, blue)
	initSide(&c.b, magenta)
	initSide(&c.d, cyan)
}

func colorSq() func(clr int) string {
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	return func(clr int) string {
		switch clr {
		case 0:
			return red("■")
		case 1:
			return white("■")
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
		for i := 0; i < 3; i++ {
			swap(&c.r[row][i], &c.f[row][i])
			swap(&c.f[row][i], &c.l[row][i])
			swap(&c.l[row][i], &c.b[row][i])
		}
	} else {
		for i := 0; i < 3; i++ {
			swap(&c.l[row][i], &c.b[row][i])
			swap(&c.f[row][i], &c.l[row][i])
			swap(&c.r[row][i], &c.f[row][i])
		}
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

func (c *cube) rotateCol(col int, inv bool) {
	if inv != true {
		for i := 0; i < 3; i++ {
			swap(&c.f[i][col], &c.u[i][col])
			swap(&c.d[i][col], &c.f[i][col])
			swap(&c.d[i][col], &c.b[2-i][2-col])
		}
	} else {
		for i := 0; i < 3; i++ {
			swap(&c.d[i][col], &c.b[2-i][2-col])
			swap(&c.d[i][col], &c.f[i][col])
			swap(&c.f[i][col], &c.u[i][col])
		}
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

func (c *cube) rotateFace(face int, inv bool) {
	if inv != true {
		for i := 0; i < 3; i++ {
			swap(&c.u[2-face][i], &c.r[i][face])
		}
		for i := 0; i < 3; i++ {
			swap(&c.l[i][2-face], &c.u[2-face][2-i])
			swap(&c.l[i][2-face], &c.d[face][i])
		}
	} else {
		for i := 0; i < 3; i++ {
			swap(&c.l[i][2-face], &c.d[face][i])
			swap(&c.l[i][2-face], &c.u[2-face][2-i])
		}
		for i := 0; i < 3; i++ {
			swap(&c.u[2-face][i], &c.r[i][face])
		}
	}

	switch face {
	case 0:
		transpose(&c.f)
		rotateHelper(&c.f, &c.f, true != inv)
	case 2:
		transpose(&c.b)
		rotateHelper(&c.b, &c.b, false || inv)
	}

}

func (c *cube) shuffle(times int, debug bool) {
	for i := 0; i < times; i++ {
		mov := rand.Intn(2)
		switch mov {
		case 0:
			val := rand.Intn(3)
			inv := bool(0 == rand.Intn(2))
			c.rotateRow(val, inv)
			if debug {
				fmt.Println("Row, value: ", val, "inv: ", inv)
				c.printCube()
			}
		case 1:
			val := rand.Intn(3)
			inv := bool(0 == rand.Intn(2))
			c.rotateCol(val, inv)
			if debug {
				fmt.Println("Col, value: ", val, "inv: ", inv)
				c.printCube()
			}
		case 2:
			val := rand.Intn(3)
			inv := bool(0 == rand.Intn(2))
			c.rotateFace(val, inv)
			if debug {
				fmt.Println("Face, value: ", val, "inv: ", inv)
				c.printCube()
			}
		}
	}
}

func main() {
	var c cube
	c.init()
	c.printCube()
	c.shuffle(20, false)
	c.printCube()
}

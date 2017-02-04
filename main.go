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

type sq struct {
	right, top *sq
	color      int
}

func (s *sq) r(x int) *sq {
	var tmp *sq = s
	for i := 0; i < x; i++ {
		tmp = tmp.right
	}
	return tmp
}

func (s *sq) t(x int) *sq {
	var tmp *sq = s
	for i := 0; i < x; i++ {
		tmp = tmp.top
	}
	return tmp
}

type cube struct {
	r1, r2, r3, c1, c2, c3 *sq
}

func colorIt() func(clr int) string {
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
	clr := colorIt()

	fmt.Println("        ---------")
	for i := 3; i >= 1; i-- {
		fmt.Printf("        | %s %s %s |\n", clr(c.c1.t(i).color),
			clr(c.c2.t(i).color), clr(c.c3.t(i).color))
	}

	fmt.Println("---------------------------------")

	fmt.Printf("| %s %s %s |", clr(c.r1.r(9).color), clr(c.r1.r(10).color), clr(c.r1.r(11).color))
	for i := 0; i <= 6; i = i + 3 {
		fmt.Printf(" %s %s %s |", clr(c.r1.r(i).color), clr(c.r1.r(i+1).color), clr(c.r1.r(i+2).color))
	}
	fmt.Printf("\n")

	fmt.Printf("| %s %s %s |", clr(c.r2.r(9).color), clr(c.r2.r(10).color), clr(c.r2.r(11).color))
	for i := 0; i <= 6; i = i + 3 {
		fmt.Printf(" %s %s %s |", clr(c.r2.r(i).color), clr(c.r2.r(i+1).color), clr(c.r2.r(i+2).color))
	}
	fmt.Printf("\n")

	fmt.Printf("| %s %s %s |", clr(c.r3.r(9).color), clr(c.r3.r(10).color), clr(c.r3.r(11).color))
	for i := 0; i <= 6; i = i + 3 {
		fmt.Printf(" %s %s %s |", clr(c.r3.r(i).color), clr(c.r3.r(i+1).color), clr(c.r3.r(i+2).color))
	}
	fmt.Printf("\n")

	fmt.Println("---------------------------------")

	for i := 9; i >= 7; i-- {
		fmt.Printf("        | %s %s %s |\n", clr(c.c1.t(i).color),
			clr(c.c2.t(i).color), clr(c.c3.t(i).color))
	}
	fmt.Println("        ---------")
}

func (c *cube) fillColor() {
	clr := []int{0, 1, 2, 3}

	for i := 0; i <= 3; i++ {
		c.r1.r(i*3).color, c.r1.r(i*3+1).color, c.r1.r(i*3+2).color = clr[i], clr[i], clr[i]
		c.r2.r(i*3).color, c.r2.r(i*3+1).color, c.r2.r(i*3+2).color = clr[i], clr[i], clr[i]
		c.r3.r(i*3).color, c.r3.r(i*3+1).color, c.r3.r(i*3+2).color = clr[i], clr[i], clr[i]
	}

	c.c1.t(1).color, c.c1.t(2).color, c.c1.t(3).color = 4, 4, 4
	c.c2.t(1).color, c.c2.t(2).color, c.c2.t(3).color = 4, 4, 4
	c.c3.t(1).color, c.c3.t(2).color, c.c3.t(3).color = 4, 4, 4

	c.c1.t(9).color, c.c1.t(8).color, c.c1.t(7).color = 5, 5, 5
	c.c2.t(9).color, c.c2.t(8).color, c.c2.t(7).color = 5, 5, 5
	c.c3.t(9).color, c.c3.t(8).color, c.c3.t(7).color = 5, 5, 5
}

func initRow() (r1 *sq) {
	var next, prev, first *sq

	prev = new(sq)
	first = prev

	for i := 0; i < 11; i++ {
		next = new(sq)
		prev.right, prev = next, next
	}
	next.right = first
	return first
}

func initRows() (r1, r2, r3 *sq) {
	return initRow(), initRow(), initRow()
}

func initCols(r1, r2, r3 *sq) (c1, c2, c3 *sq) {
	c1, c2, c3 = r1, r1.r(1), r1.r(2)
	return c1, c2, c3
}

func createLink(c1, c2 *sq) {
	tmp := c1
	for i := 0; i < 3; i++ {
		next := new(sq)
		tmp.top = next
		tmp = next
	}
	tmp.top = c2
}

func linkRowsCols(r1, r2, r3, c1, c2, c3 *sq) {
	tmp1, tmp2, tmp3 := r1, r2, r3

	for i := 0; i < 3; i++ {
		tmp3.top = tmp2
		tmp2.top = tmp1
		tmp1, tmp2, tmp3 = tmp1.r(1), tmp2.r(1), tmp3.r(1)
	}
	tmp1, tmp2, tmp3 = tmp1.r(3), tmp2.r(3), tmp3.r(3)

	for i := 0; i < 3; i++ {
		tmp1.top = tmp2
		tmp2.top = tmp3
		tmp1, tmp2, tmp3 = tmp1.r(1), tmp2.r(1), tmp3.r(1)
	}

	createLink(c1, c1.r(8))
	createLink(c2, c2.r(6))
	createLink(c3, c3.r(4))

	b1, b2, b3 := r3.r(6), r3.r(7), r3.r(8)
	createLink(b1, b1.r(8))
	createLink(b2, b2.r(6))
	createLink(b3, b3.r(4))
}

func initCube() (cb *cube) {
	r1, r2, r3 := initRows()
	c1, c2, c3 := initCols(r1, r2, r3)
	linkRowsCols(r1, r2, r3, c1, c2, c3)
	return &cube{r1, r2, r3, c1, c2, c3}
}

func main() {
	c1 := initCube()
	c1.fillColor()
	c1.printCube()
}

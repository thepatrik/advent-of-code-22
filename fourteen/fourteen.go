package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type pos struct {
	x int
	y int
}

type material int

const (
	rock material = iota
	water
)

type cave map[pos]material

func (cave cave) bounds() (int, int, int, int) {
	s, n, w, e := 0, math.MaxInt, math.MaxInt, 0
	for p := range cave {
		if p.y > s {
			s = p.y
		}
		if p.y < n {
			n = p.y
		}
		if p.x < w {
			w = p.x
		}
		if p.x > e {
			e = p.x
		}
	}

	n = 0

	return s, w, n, e
}

func (cave cave) Print() {
	s, w, n, e := cave.bounds()

	for y := n; y <= s; y++ {
		for x := w; x <= e; x++ {
			p := pos{x: x, y: y}
			if val, ok := cave[p]; ok {
				if val == rock {
					fmt.Print("#")
				} else {
					fmt.Print("o")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Draw(screen tcell.Screen, cave cave) {
	s, w, n, e := cave.bounds()

	row := 0
	for y := n; y <= s; y++ {
		col := 0
		for x := w; x <= e; x++ {
			p := pos{x: x, y: y}
			if val, ok := cave[p]; ok {
				if val == rock {
					screen.SetContent(col, row, ' ', nil, rockStyle)
				} else {
					screen.SetContent(col, row, ' ', nil, sandStyle)
				}
			} else {
				screen.SetContent(col, row, ' ', nil, defStyle)
			}
			col++
		}
		row++
	}
}

var (
	defStyle  = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorReset)
	rockStyle = defStyle.Background(tcell.ColorBlack)
	sandStyle = defStyle.Background(tcell.ColorRosyBrown)
)

func main() {
	cavem := parse("./data")
	south, _, _, _ := cavem.bounds()

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	s.SetStyle(defStyle)

	// Clear screen
	s.Clear()

	drip := pos{x: 500, y: 0}

	var dripper func(cave cave, currentPos pos, south int) (pos, bool)
	dripper = func(cave cave, currentPos pos, south int) (pos, bool) {
		if currentPos.y > south {
			return currentPos, false
		}

		// time.Sleep(1 * time.Microsecond)
		cavem[currentPos] = water
		Draw(s, cavem)
		s.Show()

		defer func() {
			delete(cavem, currentPos)
			s.Show()
		}()

		//down
		nextPos := pos{x: currentPos.x, y: currentPos.y + 1}
		if _, ok := cave[nextPos]; !ok {
			return dripper(cave, nextPos, south)
		}

		//down left
		nextPos.x = currentPos.x - 1
		if _, ok := cave[nextPos]; !ok {
			return dripper(cave, nextPos, south)
		}

		//down right
		nextPos.x = currentPos.x + 1
		if _, ok := cave[nextPos]; !ok {
			return dripper(cave, nextPos, south)
		}

		return currentPos, true
	}

	draw := func() bool {
		pos, b := dripper(cavem, drip, south)
		if !b {
			return false

		}
		cavem[pos] = water
		Draw(s, cavem)
		return true
	}

	Draw(s, cavem)

	start := func() {
		for draw() {
		}
	}

	go start()

	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}
}

func PartOne(filename string) int {
	cavem := parse(filename)
	south, _, _, _ := cavem.bounds()

	drip := pos{x: 500, y: 0}

	var dripper func(cave cave, currentPos pos, south int) (pos, bool)
	dripper = func(cave cave, currentPos pos, south int) (pos, bool) {
		if currentPos.y > south {
			return currentPos, false
		}

		//down
		nextPos := pos{x: currentPos.x, y: currentPos.y + 1}
		if _, ok := cave[nextPos]; !ok {
			return dripper(cave, nextPos, south)
		}

		//down left
		nextPos.x = currentPos.x - 1
		if _, ok := cave[nextPos]; !ok {
			return dripper(cave, nextPos, south)
		}

		//down right
		nextPos.x = currentPos.x + 1
		if _, ok := cave[nextPos]; !ok {
			return dripper(cave, nextPos, south)
		}

		return currentPos, true
	}

	drips := func() int {
		for i := 0; ; i++ {
			pos, b := dripper(cavem, drip, south)
			if !b {
				return i
			}
			cavem[pos] = water
		}
	}()

	return drips
}

func PartTwo(filename string) int {
	cavem := parse(filename)
	south, _, _, _ := cavem.bounds()

	drip := pos{x: 500, y: 0}

	var dripper func(cave cave, currentPos pos, floor int) (pos, bool)
	dripper = func(cave cave, currentPos pos, floor int) (pos, bool) {
		nextPos := pos{x: currentPos.x, y: currentPos.y + 1}
		if nextPos.y < floor {
			if _, ok := cave[nextPos]; nextPos.y < floor && !ok {
				return dripper(cave, nextPos, floor)
			}

			nextPos.x = currentPos.x - 1
			if _, ok := cave[nextPos]; !ok {
				return dripper(cave, nextPos, floor)
			}

			nextPos.x = currentPos.x + 1
			if _, ok := cave[nextPos]; !ok {
				return dripper(cave, nextPos, floor)
			}
		}

		_, ok := cave[currentPos]

		return currentPos, !ok
	}

	drips := func() int {
		for i := 0; ; i++ {
			pos, b := dripper(cavem, drip, south+2)
			if !b {
				return i
			}
			cavem[pos] = water
		}
	}()

	// cavem.Print()

	return drips
}

func parse(filename string) cave {
	rocks := [][]pos{}
	strslice := parser.ReadFile(filename)
	for _, line := range strslice {
		split := strings.Split(line, " -> ")
		positions := []pos{}
		for _, s := range split {
			spos := strings.Split(s, ",")
			x, _ := strconv.Atoi(spos[0])
			y, _ := strconv.Atoi(spos[1])
			pos := pos{x: x, y: y}
			positions = append(positions, pos)
		}
		rocks = append(rocks, positions)
	}

	cavem := make(cave)
	for _, r := range rocks {
		var prev pos
		for i, p := range r {
			if i == 0 {
				prev = p
				continue
			}

			for i := prev.x; i <= p.x; i++ {
				cavem[pos{x: i, y: p.y}] = rock
			}
			for i := prev.x; i >= p.x; i-- {
				cavem[pos{x: i, y: p.y}] = rock
			}
			for i := prev.y; i <= p.y; i++ {
				cavem[pos{x: p.x, y: i}] = rock
			}
			for i := prev.y; i >= p.y; i-- {
				cavem[pos{x: p.x, y: i}] = rock
			}
			prev = p
		}
	}

	return cavem
}

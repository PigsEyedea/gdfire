package main

import (
	"fmt"
	"charm.land/lipgloss/v2"
	"golang.org/x/term"
	"math/rand/v2"
	"time"
)




func updateFlames(mat [][]int) [][]int {


	for level := 0; level<len(mat) ; level++ {
		row := mat[level]
		for idx,_ := range row {
			cell := 0
			if level == 0 {
				cell = level	
			} else {
				ncell := mat[level-1][idx]
				ncell += rand.IntN(6)
				cell = ncell
				if cell > len(mat)-2 {
					cell = len(mat)-2
				}
				if cell < 0 {
					cell = 0
				} 
					
			}
			mat[level][idx]=cell
		}
	}

	return mat

	
	
}


func drawFlames(mat[][]int,height int){
	var styler = lipgloss.NewStyle()
	pal := lipgloss.Blend1D(
		height-1,
		lipgloss.Color("#FFFFFF"),//white
		lipgloss.Color("#FFFF00"),//yellow
		lipgloss.Color("#FFA500"),//orange
		lipgloss.Color("#FF0000"),//red
		lipgloss.Color("#000000"),//black
	)

	for level := len(mat)-1; level>(-1); level-=1 {
		row := mat[level]
		for idx,_ := range row {
			color := pal[row[idx]]
			lipgloss.Print(styler.Foreground(color).Render("@"))
		}
		fmt.Println()
	}

	
	
}

func main(){

	width, height, err := term.GetSize(0)
	if err != nil {
		return
	}



	matrix := make([][]int,height)
	for idx,_ := range matrix {
		matrix[idx] = make([]int, width)
	
	}

	for {

		width2, height2, err2 := term.GetSize(0)
		if err2 != nil {
			return
		}

		if width2 != width || height2 != height {
			width = width2
			height = height2
			matrix = make([][]int,height)
			for idx,_ := range matrix {
				matrix[idx] = make([]int, width)	
			}
			
		}
		
		fmt.Print("\033[H\033[2J")
		matrix = updateFlames(matrix)
		drawFlames(matrix,height)
		time.Sleep(80 * time.Millisecond)
	}
	
}

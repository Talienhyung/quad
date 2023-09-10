package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if (j == 1 || j == x) && i == 1 {
					z01.PrintRune('A')
				} else if (j == 1 || j == x) && i == y {
					z01.PrintRune('C')
				} else if (j != 1 && j != x) && (i != 1 && i != y) {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

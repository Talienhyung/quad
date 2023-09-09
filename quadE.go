package quad

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x > 1 && y > 1 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if (j == 1 && i == 1) || (j == x && i == y) {
					z01.PrintRune('A')
				} else if (i == 1 && j == x) || (i == y && j == 1) {
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
	if x == 1 && y == 1 {
		z01.PrintRune('A')
		z01.PrintRune('\n')
	}
	if x == 1 && y != 1 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == y {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		}

	}
	if y == 1 && x != 1 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if j == 1 {
					z01.PrintRune('A')
				} else if j == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

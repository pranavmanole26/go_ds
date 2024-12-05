package main

import "fmt"

func Pattern1() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(" * ")
		}
		fmt.Println("")
	}
}

func Pattern2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" * ")
		}
		fmt.Println("")
	}
}

func Pattern3() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf(" %d ", j+1)
		}
		fmt.Println("")
	}
}

func Pattern4() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf(" %d ", i+1)
		}
		fmt.Println("")
	}
}

func Pattern5() {
	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println("")
	}
}

func Pattern6() {
	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf(" %d ", j+1)
		}
		fmt.Println("")
	}
}

func Pattern7() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 9; j++ {
			if j >= (9/2-i) && j <= (9/2+i) {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("")
	}
}

func Pattern8() {
	for i := 4; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			if j >= (9/2-i) && j <= (9/2+i) {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("")
	}
}

func Pattern9() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if i < 10/2 && j >= (9/2-i) && j <= (9/2+i) {
				fmt.Print(" * ")
			} else if i >= 10/2 {
				newI := 5 - i%5
				if j > (9/2-newI) && j <= (9/2+newI-1) {
					fmt.Print(" * ")
				} else {
					fmt.Print("   ")
				}
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("")
	}
}

func Pattern9_1() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 9; j++ {
			if j >= 9/2-i && j <= 9/2+i {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
	for i := 4; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			if j >= 9/2-i && j <= 9/2+i {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func Pattern10() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 5; j++ {
			if i <= 9/2 {
				if j <= i {
					fmt.Print(" * ")
				} else {
					fmt.Print("   ")
				}
			} else {
				newI := 4 - i%5
				if j < newI {
					fmt.Print(" * ")
				} else {
					fmt.Print("   ")
				}
			}
		}
		fmt.Println()
	}
}

func Pattern11() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j <= i {
				fmt.Printf(" %d ", (j+i+1)%2)
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func Pattern12() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 8; j++ {
			if j <= 8/2 {
				if j <= i {
					fmt.Printf("%d", j)
				} else {
					fmt.Print(" ")
				}
			} else {
				if i >= (9 - j) {
					fmt.Printf("%d", 9-j)
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

func Pattern13() {
	c := 0
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j <= i {
				c++
				fmt.Printf(" %d ", c)
			} else {
				break
			}
		}
		fmt.Println()
	}
}

func Pattern14() {
	charMap := map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E",
	}
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j <= i {
				fmt.Printf(" %s ", charMap[j])
			} else {
				break
			}
		}
		fmt.Println()
	}
}

func Pattern15() {
	charMap := map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E",
	}
	for i := 5; i >= 1; i-- {
		for j := 1; j <= 5; j++ {
			if j <= i {
				fmt.Printf(" %s ", charMap[j])
			}
		}
		fmt.Println()
	}
}

func Pattern16() {
	charMap := map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E",
	}
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j <= i {
				fmt.Printf(" %s ", charMap[i])
			} else {
				break
			}
		}
		fmt.Println()
	}
}

func Pattern17() {
	charMap := map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E",
	}
	for i := 1; i <= 4; i++ {
		c := 0
		for j := 1; j <= 7; j++ {
			if j > (7/2+1-i) && j <= (7/2+i) {
				if j <= (7/2 + 1) {
					c++
				} else {
					c--
				}
				fmt.Printf("%s", charMap[c])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Pattern18() {
	charMap := map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E",
	}
	for i := 1; i <= 5; i++ {
		c := 6 - i
		for j := 1; j <= 5; j++ {
			if j <= i {
				fmt.Printf("%s", charMap[c])
				c++
			} else {
				break
			}
		}
		fmt.Println()
	}
}

func Pattern19() {
	newI := 0
	for i := 1; i <= 10; i++ {
		if i <= 10/2 {
			for j := 1; j <= 5; j++ {
				if j+i <= 6 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			for j := 1; j <= 5; j++ {
				if j >= i {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		} else {
			newI++
			for j := 1; j <= 5; j++ {
				if j <= newI {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			for j := 1; j <= 5; j++ {
				if j+newI >= 6 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

func Pattern20() {
	newI := 0
	for i := 1; i < 10; i++ {
		if i <= 10/2 {
			for j := 1; j <= 5; j++ {
				if j <= i {
					fmt.Print("*")
				} else {
					fmt.Print("-")
				}
			}
			for j := 1; j <= 5; j++ {
				if i+j >= 6 {
					fmt.Print("*")
				} else {
					fmt.Print("-")
				}
			}
		} else {
			newI++
			for j := 1; j <= 5; j++ {
				if j+newI <= 5 {
					fmt.Print("*")
				} else {
					fmt.Print("-")
				}
			}
			for j := 1; j <= 5; j++ {
				if j > newI {
					fmt.Print("*")
				} else {
					fmt.Print("-")
				}
			}
		}
		fmt.Println()
	}
}

func Pattern21() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if i == 1 || i == 4 ||
				j == 1 || j == 4 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Pattern22() {
	for i := 1; i <= 7; i++ {
		for j := 1; j <= 7; j++ {

		}
		fmt.Println()
	}
}

func main() {
	// Pattern1()
	// Pattern2()
	// Pattern3()
	// Pattern4()
	// Pattern5()
	// Pattern6()
	// Pattern7()
	// Pattern8()
	// Pattern9()
	// Pattern9_1()
	// Pattern10()
	// Pattern11()
	// Pattern12()
	// Pattern13()
	// Pattern14()
	// Pattern15()
	// Pattern16()
	// Pattern17()
	// Pattern18()
	// Pattern19()
	// Pattern20()
	Pattern21()
}

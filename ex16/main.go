package main

import (
	"ft"
	"piscine"
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(9223372036854775807, "0123456789ABCDEF")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-9223372036854775808, "0123456789ABCDEF")
	ft.PrintRune('\n')
}

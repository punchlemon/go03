package main

import (
	"ft"
	"piscine"
)

func main() {
	var a string
	ft.PrintRune(piscine.LastRune("Hello!"))
	ft.PrintRune(piscine.LastRune("Salut!"))
	ft.PrintRune(piscine.LastRune("Bye!"))
	ft.PrintRune(piscine.LastRune("Bye!"))
	ft.PrintRune(piscine.LastRune("Ola!"))
	ft.PrintRune(piscine.LastRune(""))
	ft.PrintRune(piscine.LastRune(a))
	ft.PrintRune('\n')
}

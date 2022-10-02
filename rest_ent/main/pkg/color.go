package pkg

import (
	"fmt"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func Colorize(color Color, message string) string {
	return string(color)+ message+ string(ColorReset)
} 


func PrintC(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}
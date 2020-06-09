package utils

import "fmt"

const (
	BLACK = iota
	GREEN
	RED
)

type Printer interface {
	LineBreak(a ...interface{})
	NoLineBreak(a ...interface{})
}

//黑色输出
type Black struct {
}

func (b *Black) LineBreak(a ...interface{}) {
	fmt.Println(a...)
}

func (b *Black) NoLineBreak(a ...interface{}) {
	fmt.Print(a...)
}

//绿色输出
type Green struct {
}

func (b *Green) LineBreak(a ...interface{}) {
	fmt.Print("\033[32m")
	fmt.Println(a...)
	fmt.Print("\033[0m")
}

func (b *Green) NoLineBreak(a ...interface{}) {
	fmt.Print("\033[32m")
	fmt.Print(a...)
	fmt.Print("\033[0m")
}

//红色输出
type Red struct {
}

func (b *Red) LineBreak(a ...interface{}) {
	fmt.Print("\033[31m")
	fmt.Println(a...)
	fmt.Print("\033[0m")
}

func (b *Red) NoLineBreak(a ...interface{}) {
	fmt.Print("\033[31m")
	fmt.Print(a...)
	fmt.Print("\033[0m")
}

func PrinterCreate(types int) Printer {
	switch types {
	case BLACK:
		return &Black{}
	case GREEN:
		return &Green{}
	case RED:
		return &Red{}
	default:
		return &Black{}
	}
}

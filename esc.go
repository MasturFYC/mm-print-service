package main

import (
	"fmt"
	"io"
)

type Esc struct {
	w io.Writer
}

func NewEscpos(w io.Writer) *Esc {
	return &Esc{
		w: w,
	}
}

func (e *Esc) Init() (int, error) {
	return fmt.Fprintf(e.w, "%c%c", 27, 64)
}

/*
length = lines
*/
func (e *Esc) PageLength(length byte) (int, error) {
	return fmt.Fprintf(e.w, "%c%c%c%c%c", 27, 50, 27, 67, length)
}

/*
0 = roman, 1 = sans-serif
*/
func (e *Esc) Typeface(face byte) (int, error) {
	return fmt.Fprintf(e.w, "%c%c%c%c", 27, 107, face, 1)
}

// ESC g
//
// 80 = 10-cpi, 77 = 12-cpi , 103 = 15-cpi
func (e *Esc) Pitch(cpi byte) (int, error) {
	return fmt.Fprintf(e.w, "%c%c", 27, cpi)
}

// POS 8
// 1 small
// 0 standard
func (e *Esc) Character(cpi byte) (int, error) {
	return fmt.Fprintf(e.w, "%c%c%c", 27, 77, cpi)
}

// POS 8
func (e *Esc) Feed(line byte) (int, error) {
	return fmt.Fprintf(e.w, "%c%c%c", 27, 100, line)
}

// POS 8
func (e *Esc) Cut() (int, error) {
	return fmt.Fprintf(e.w, "%cV%c", 29, 49)
}

func (e *Esc) Print(s string) (int, error) {
	return fmt.Fprint(e.w, s)
}

func (e *Esc) Println(s string) (int, error) {
	fmt.Fprint(e.w, s)
	return fmt.Fprint(e.w, "\n")
}

func (e *Esc) Eject() (int, error) {
	return fmt.Fprintf(e.w, "%c%c%c", 27, 25, 66)
}

func (e *Esc) Bold() (int, error) {
	return fmt.Fprintf(e.w, "%c%c", 27, 69)
}

func (e *Esc) Unbold() (int, error) {
	return fmt.Fprintf(e.w, "%c%c", 27, 70)
}

func (e *Esc) Italic() (int, error) {
	return fmt.Fprintf(e.w, "%c%c", 27, 52)
}

func (e *Esc) Unitalic() (int, error) {
	return fmt.Fprintf(e.w, "%c%c", 27, 53)
}

/*
ESC t n

0 = italic, 1 = PC437, 2 = User-defined, 4 = PC437
*/
func (e *Esc) TableChar(table byte) (int, error) {
	return fmt.Fprintf(e.w, "%c%c%c", 27, 116, table)
}

/*
ESC x

0 = draft, 1 = Lq
*/
func (e *Esc) LqMode(mode byte) (int, error) {
	return fmt.Fprintf(e.w, "%c%c%c", 27, 120, mode)
}

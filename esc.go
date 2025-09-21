package main

import (
	"fmt"
	"io"
)

type Esc struct{}

func (t Esc) Init(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "%c%c", 27, 64)
}

/*
length = lines
*/
func (t Esc) PageLength(w io.Writer, length byte) (int, error) {
	return fmt.Fprintf(w, "%c%c%c%c%c", 27, 50, 27, 67, length)
}

/*
0 = roman, 1 = sans-serif
*/
func (t Esc) Typeface(w io.Writer, face byte) (int, error) {
	return fmt.Fprintf(w, "%c%c%c%c", 27, 107, face, 1)
}

// ESC g
//
// 80 = 10-cpi, 77 = 12-cpi , 103 = 15-cpi
func (t Esc) Pitch(w io.Writer, cpi byte) (int, error) {
	return fmt.Fprintf(w, "%c%c", 27, cpi)
}

// POS 8
// 1 small
// 0 standard
func (t Esc) Character(w io.Writer, cpi byte) (int, error) {
	return fmt.Fprintf(w, "%c%c%c", 27, 77, cpi)
}

// POS 8
func (t Esc) Feed(w io.Writer, line byte) (int, error) {
	return fmt.Fprintf(w, "%c%c%c", 27, 100, line)
}

// POS 8
func (t Esc) Cut(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "%cV%c", 29, 49)
}

func (t Esc) Print(w io.Writer, s string) (int, error) {
	return fmt.Fprint(w, s)
}

func (t Esc) Println(w io.Writer, s string) (int, error) {
	fmt.Fprint(w, s)
	return fmt.Fprint(w, "\n")
}

func (t Esc) Eject(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "%c%c%c", 27, 25, 66)
}

func (t Esc) Bold(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "%c%c", 27, 69)
}

func (t Esc) Unbold(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "%c%c", 27, 70)
}

func (t Esc) Italic(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "%c%c", 27, 52)
}

func (t Esc) Unitalic(w io.Writer) (int, error) {
	return fmt.Fprintf(w, "%c%c", 27, 53)
}

/*
ESC t n

0 = italic, 1 = PC437, 2 = User-defined, 4 = PC437
*/
func (t Esc) TableChar(w io.Writer, table byte) (int, error) {
	return fmt.Fprintf(w, "%c%c%c", 27, 116, table)
}

/*
ESC x

0 = draft, 1 = Lq
*/
func (t Esc) LqMode(w io.Writer, mode byte) (int, error) {
	return fmt.Fprintf(w, "%c%c%c", 27, 120, mode)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package colour

import (
	"fmt"
	"strconv"
)

const colourFmt = "\x1b[%dm%s\x1b[0m"

// Paint describes a terminal colour.
type Paint int

// Defines basic ANSI colours.
const (
	Black     Paint = iota + 30 // 30
	Red                         // 31
	Green                       // 32
	Yellow                      // 33
	Blue                        // 34
	Magenta                     // 35
	Cyan                        // 36
	LightGray                   // 37
	DarkGray  = 90

	Bold = 1
)

// colourize returns an ASCII coloured string based on given colour.
func colourize(s string, c Paint) string {
	if c == 0 {
		return s
	}
	return fmt.Sprintf(colourFmt, c, s)
}

// ANSIcolourize colours a string.
func ANSIcolourize(text string, colour int) string {
	return "\033[38;5;" + strconv.Itoa(colour) + "m" + text + "\033[0m"
}

// Highlight colourize bytes at given indices.
func Highlight(bb []byte, ii []int, c int) []byte {
	b := make([]byte, 0, len(bb))
	for i, j := 0, 0; i < len(bb); i++ {
		if j < len(ii) && ii[j] == i {
			b = append(b, colourizeByte(bb[i], 209)...)
			j++
		} else {
			b = append(b, bb[i])
		}
	}

	return b
}

func colourizeByte(b byte, colour int) []byte {
	return []byte(ANSIcolourize(string(b), colour))
}

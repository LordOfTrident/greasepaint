package greasepaint

import (
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/LordOfTrident/ansi-go"
)

const (
	AutoWidth = 0
	NoBorder  = rune(0)
)

func TerminalSize() (int, int) {
	if w, h, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		return w, h
	} else {
		return -1, -1
	}
}

func TerminalWidth() int {
	w, _ := TerminalSize()
	return w
}

func TerminalHeight() int {
	_, h := TerminalSize()
	return h
}

func Gap(size int) string {
	if size <= 0 {
		return ""
	}

	return strings.Repeat(" ", size)
}

func Line(text string, length int) string {
	if length < 0 {
		length += TerminalWidth() + 1
	}

	return string([]rune(strings.Repeat(text, length))[:length])
}

func Size(text string) int {
	return ansi.SizeNoAnsi(text)
}

func Length(text string) int {
	return ansi.LengthNoAnsi(text)
}

func Chain(fns ...Applier) Applier {
	return func(args ...string) string {
		output := strings.Join(args, "")
		for _, fn := range fns {
			output = fn(output)
		}
		return output
	}
}

type Applier func(...string) string
type Padding [2]int
type Border  [2]rune

type Target int
const (
	TargetStdout = Target(iota)
	TargetStderr
	TargetTTY
	TargetOther
)

var targetIsTTY = map[Target]bool{
	TargetStdout: true,
	TargetStderr: true,
	TargetTTY:    true,
	TargetOther:  false,
}

func init() {
	targetIsTTY[TargetStdout] = term.IsTerminal(int(os.Stdout.Fd()))
	targetIsTTY[TargetStderr] = term.IsTerminal(int(os.Stderr.Fd()))
}

func (t Target) IsTTY() bool {
	return targetIsTTY[t]
}

type Align int
const (
	AlignLeft = Align(iota)
	AlignRight
	AlignCenter
)

type Attributes map[string]bool

func (a Attributes) Bold()      bool {return a["bold"]}
func (a Attributes) Dim()       bool {return a["dim"]}
func (a Attributes) Italic()    bool {return a["italic"]}
func (a Attributes) Underline() bool {return a["underline"]}
func (a Attributes) Blink()     bool {return a["blink"]}
func (a Attributes) Reverse()   bool {return a["reverse"]}
func (a Attributes) Invisible() bool {return a["invisible"]}
func (a Attributes) Crossout()  bool {return a["crossout"]}

type Style struct {
	fg, bg  Color
	padding Padding
	margin  Padding
	border  Border
	gridGap int
	attrs   Attributes
	w       int // TODO: Height
	target  Target
	align   Align
}

func NewStyle() Style {
	return Style{fg: ColorDefault{}, bg: ColorDefault{}, attrs: make(map[string]bool)}
}

func (s Style) Apply(args ...string) (output string) {
	var contentWidth int
	seqs, contentSeqs := s.getSequences()
	totalPadding      := s.padding[0] + s.padding[1]

	for i, arg := range args {
		if i > 0 {
			output       += seqs + Gap(s.gridGap)
			contentWidth += s.gridGap
		}

		arg = strings.ReplaceAll(arg, ansi.Reset, s.target.ifTTY(ansi.Reset) + seqs)

		output       += contentSeqs + arg + s.target.ifTTY(ansi.Reset)
		contentWidth += Length(arg)
	}

	w := s.w
	if w < 0 {
		w += TerminalWidth() + 1
	}

	if w > 0 {
		gapWidth := w - contentWidth - totalPadding

		if gapWidth > 0 { // Insert gaps
			switch s.align {
			case AlignLeft:  output += seqs + Gap(gapWidth)
			case AlignRight: output  = Gap(gapWidth) + output
			case AlignCenter:
				left  := Gap(int(float64(gapWidth) / 2 + 0.5))
				right := Gap(int(float64(gapWidth) / 2))
				output = left + output + seqs + right
			}
		} else if gapWidth < 0 { // Cut off text
			// TODO: Keep the proper alignment when cutting
			seqsLen := 0
			it      := output
			pos     := 0
			for {
				loc := ansi.FindNext(it)
				if loc == nil {
					break
				}

				// The width up to the current position
				currentWidth := pos + loc[0] - seqsLen + totalPadding

				if currentWidth >= w {
					break
				}

				seqsLen += loc[1] - loc[0]
				pos     += loc[1]
				it       = it[loc[1]:]
			}

			output = output[:w + seqsLen - totalPadding] + s.target.ifTTY(ansi.Reset)
		}
	}

	left  := Gap(s.margin[0]) + seqs + s.borderStringLeft() + Gap(s.padding[0])
	right := seqs + Gap(s.padding[1]) + s.borderStringRight() +
	         s.target.ifTTY(ansi.Reset) + Gap(s.margin[1])

	output = left + output + right
	return
}

func (t Target) ifTTY(text string) string {
	if t.IsTTY() {
		return text
	} else {
		return ""
	}
}

func (s Style) borderStringLeft() string {
	if s.border[0] > 0 {
		return string(s.border[0])
	} else {
		return ""
	}
}

func (s Style) borderStringRight() string {
	if s.border[1] > 0 {
		return string(s.border[1])
	} else {
		return ""
	}
}

func (s Style) getSequences() (seqs, contentSeqs string) {
	if !s.target.IsTTY() {
		return
	}

	for attr, enabled := range s.attrs {
		if !enabled {
			continue
		}

		var seq string
		switch attr {
		case "bold":      seq = ansi.Bold
		case "dim":       seq = ansi.Dim
		case "italic":    seq = ansi.Italic
		case "underline": seq = ansi.Underline
		case "blink":     seq = ansi.Blink
		case "reverse":   seq = ansi.Reverse
		case "invisible": seq = ansi.Invisible
		case "crossout":  seq = ansi.Crossout
		}

		if attr == "underline" || attr == "crossout" {
			contentSeqs += seq
		} else {
			seqs += seq
		}
	}

	seqs += s.fg.AnsiFg()
	seqs += s.bg.AnsiBg()
	return
}

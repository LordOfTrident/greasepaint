package greasepaint

import "github.com/LordOfTrident/ansi-go"

var (
	Black   = NewColor256(ansi.Black)
	Red     = NewColor256(ansi.Red)
	Green   = NewColor256(ansi.Green)
	Yellow  = NewColor256(ansi.Yellow)
	Blue    = NewColor256(ansi.Blue)
	Magenta = NewColor256(ansi.Magenta)
	Cyan    = NewColor256(ansi.Cyan)
	White   = NewColor256(ansi.White)

	LightBlack   = NewColor256(ansi.LightBlack)
	LightRed     = NewColor256(ansi.LightRed)
	LightGreen   = NewColor256(ansi.LightGreen)
	LightYellow  = NewColor256(ansi.LightYellow)
	LightBlue    = NewColor256(ansi.LightBlue)
	LightMagenta = NewColor256(ansi.LightMagenta)
	LightCyan    = NewColor256(ansi.LightCyan)
	LightWhite   = NewColor256(ansi.LightWhite)
)

type Color interface {
	AnsiFg() string
	AnsiBg() string
}

type ColorDefault struct {}
type Color256     struct {Value uint8}
type ColorRGB     struct {R, G, B uint8}
type ColorHex     struct {Value string}

func NewColor256(value uint8)   Color256 {return Color256{Value: value}}
func NewColorRGB(r, g, b uint8) ColorRGB {return ColorRGB{R: r, G: g, B: b}}
func NewColorHex(value string)  ColorHex {return ColorHex{Value: value}}

func (c ColorDefault) AnsiFg() string {return ""}
func (c ColorDefault) AnsiBg() string {return ""}
func (c Color256)     AnsiFg() string {return ansi.Fg(c.Value)}
func (c Color256)     AnsiBg() string {return ansi.Bg(c.Value)}
func (c ColorRGB)     AnsiFg() string {return ansi.FgRGB(c.R, c.G, c.B)}
func (c ColorRGB)     AnsiBg() string {return ansi.BgRGB(c.R, c.G, c.B)}
func (c ColorHex)     AnsiFg() string {return ansi.FgHex(c.Value)}
func (c ColorHex)     AnsiBg() string {return ansi.BgHex(c.Value)}

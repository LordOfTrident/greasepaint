package greasepaint

func (s Style) Width(w int) Style {
	s.w = w
	return s
}

func (s Style) Target(target Target) Style {
	s.target = target
	return s
}

func (s Style) Align(align Align) Style {
	s.align = align
	return s
}

func (s Style) GridGap(gridGap int) Style {
	s.gridGap = gridGap
	return s
}

func (s Style) Border(border rune) Style {
	s.border = Border{border, border}
	return s
}

func (s Style) BorderLeft(border rune) Style {
	s.border[0] = border
	return s
}

func (s Style) BorderRight(border rune) Style {
	s.border[1] = border
	return s
}

// Colors
func (s Style) DefaultFg() Style {
	s.fg = ColorDefault{}
	return s
}

func (s Style) DefaultBg() Style {
	s.bg = ColorDefault{}
	return s
}

func (s Style) Fg(color Color) Style {
	s.fg = color
	return s
}

func (s Style) Bg(color Color) Style {
	s.bg = color
	return s
}

func (s Style) Fg256(value uint8) Style {
	s.fg = NewColor256(value)
	return s
}

func (s Style) Bg256(value uint8) Style {
	s.bg = NewColor256(value)
	return s
}

func (s Style) FgRGB(r, g, b uint8) Style {
	s.fg = NewColorRGB(r, g, b)
	return s
}

func (s Style) BgRGB(r, g, b uint8) Style {
	s.bg = NewColorRGB(r, g, b)
	return s
}

func (s Style) FgHex(value string) Style {
	s.fg = NewColorHex(value)
	return s
}

func (s Style) BgHex(value string) Style {
	s.bg = NewColorHex(value)
	return s
}

// Attributes
func (s Style) Bold(enable bool) Style {
	s.attrs["bold"] = enable
	return s
}

func (s Style) Dim(enable bool) Style {
	s.attrs["dim"] = enable
	return s
}

func (s Style) Italic(enable bool) Style {
	s.attrs["italic"] = enable
	return s
}

func (s Style) Underline(enable bool) Style {
	s.attrs["underline"] = enable
	return s
}

func (s Style) Blink(enable bool) Style {
	s.attrs["blink"] = enable
	return s
}

func (s Style) Reverse(enable bool) Style {
	s.attrs["reverse"] = enable
	return s
}

func (s Style) Invisible(enable bool) Style {
	s.attrs["invisible"] = enable
	return s
}

func (s Style) Crossout(enable bool) Style {
	s.attrs["crossout"] = enable
	return s
}

// Padding
func (s Style) Padding(size int) Style {
	s.padding = Padding{size, size}
	return s
}

func (s Style) PaddingLeft(size int) Style {
	s.padding[0] = size
	return s
}

func (s Style) PaddingRight(size int) Style {
	s.padding[1] = size
	return s
}

// Margin
func (s Style) Margin(size int) Style {
	s.margin = Padding{size, size}
	return s
}

func (s Style) MarginLeft(size int) Style {
	s.margin[0] = size
	return s
}

func (s Style) MarginRight(size int) Style {
	s.margin[1] = size
	return s
}

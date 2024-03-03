package greasepaint

func (s Style) GetWidth() int {
	return s.w
}

func (s Style) GetTarget() Target {
	return s.target
}

func (s Style) GetAlign() Align {
	return s.align
}

func (s Style) GetGridGap() int {
	return s.gridGap
}

func (s Style) GetBorder() Border {
	return s.border
}

func (s Style) GetFg() Color {
	return s.fg
}

func (s Style) GetBg() Color {
	return s.bg
}

func (s Style) GetAttributes() map[string]bool {
	return s.attrs
}

func (s Style) GetPadding() Padding {
	return s.padding
}

func (s Style) GetMargin() Padding {
	return s.margin
}

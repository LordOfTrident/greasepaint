package greasepaint

var (
	Bold       = NewStyle().Bold(true).Apply
	Dim        = NewStyle().Dim(true).Apply
	Italic     = NewStyle().Italic(true).Apply
	Underline  = NewStyle().Underline(true).Apply
	Blink      = NewStyle().Blink(true).Apply
	Reverse    = NewStyle().Reverse(true).Apply
	Invisible  = NewStyle().Invisible(true).Apply
	Crossout   = NewStyle().Crossout(true).Apply

	BlackFg   = NewStyle().Fg(Black).Apply
	RedFg     = NewStyle().Fg(Red).Apply
	GreenFg   = NewStyle().Fg(Green).Apply
	YellowFg  = NewStyle().Fg(Yellow).Apply
	BlueFg    = NewStyle().Fg(Blue).Apply
	MagentaFg = NewStyle().Fg(Magenta).Apply
	CyanFg    = NewStyle().Fg(Cyan).Apply
	WhiteFg   = NewStyle().Fg(White).Apply

	BlackBg   = NewStyle().Bg(Black).Apply
	RedBg     = NewStyle().Bg(Red).Apply
	GreenBg   = NewStyle().Bg(Green).Apply
	YellowBg  = NewStyle().Bg(Yellow).Apply
	BlueBg    = NewStyle().Bg(Blue).Apply
	MagentaBg = NewStyle().Bg(Magenta).Apply
	CyanBg    = NewStyle().Bg(Cyan).Apply
	WhiteBg   = NewStyle().Bg(White).Apply

	LightBlackFg   = NewStyle().Fg(LightBlack).Apply
	LightRedFg     = NewStyle().Fg(LightRed).Apply
	LightGreenFg   = NewStyle().Fg(LightGreen).Apply
	LightYellowFg  = NewStyle().Fg(LightYellow).Apply
	LightBlueFg    = NewStyle().Fg(LightBlue).Apply
	LightMagentaFg = NewStyle().Fg(LightMagenta).Apply
	LightCyanFg    = NewStyle().Fg(LightCyan).Apply
	LightWhiteFg   = NewStyle().Fg(LightWhite).Apply

	LightBlackBg   = NewStyle().Bg(LightBlack).Apply
	LightRedBg     = NewStyle().Bg(LightRed).Apply
	LightGreenBg   = NewStyle().Bg(LightGreen).Apply
	LightYellowBg  = NewStyle().Bg(LightYellow).Apply
	LightBlueBg    = NewStyle().Bg(LightBlue).Apply
	LightMagentaBg = NewStyle().Bg(LightMagenta).Apply
	LightCyanBg    = NewStyle().Bg(LightCyan).Apply
	LightWhiteBg   = NewStyle().Bg(LightWhite).Apply

	Center = NewStyle().Width(-1).Align(AlignCenter).Apply
	Left   = NewStyle().Width(-1).Align(AlignLeft).Apply
	Right  = NewStyle().Width(-1).Align(AlignRight).Apply
)

package main

import gp "github.com/LordOfTrident/greasepaint"

var (
	accent = gp.NewColorHex("ee0")

	styleHeader = gp.NewStyle().
		FgHex("000").
		Bg   (accent).
		Width(-1).
		Bold (true).
		Align(gp.AlignCenter).
		Apply

	styleTabs = gp.NewStyle().
		BgHex  ("444").
		Padding(2).
		GridGap(1).
		Width  (-1).
		Apply

	styleTabInactive = gp.NewStyle().
		FgHex  ("aaa").
		Padding(1).
		Apply

	styleTabActive = gp.NewStyle().
		Bold       (true).
		Underline  (true).
		BgHex      ("666").
		FgHex      ("fff").
		Padding    (1).
		Apply

	styleCode = gp.NewStyle().
		FgHex  ("fff").
		BgHex  ("444").
		Padding(1).
		Apply

	styleTitle = gp.NewStyle().
		Bold       (true).
		Underline  (true).
		FgHex      ("fff").
		MarginLeft (2).
		BorderRight(':').
		Apply

	stylePageText = gp.NewStyle().
		MarginLeft(4).
		Apply

	styleKeyword = gp.NewStyle().
		Bold(true).
		Fg  (accent).
		Apply

	styleUsageTitle = gp.NewStyle().
		Bold      (true).
		Fg        (gp.LightWhite).
		MarginLeft(2).
		Apply

	styleError = gp.NewStyle().
		Bold      (true).
		Fg        (gp.LightRed).
		MarginLeft(2).
		Apply

	styleSeparator = gp.Chain(gp.NewStyle().
		FgHex("444").
		Apply,
		gp.Center)
)

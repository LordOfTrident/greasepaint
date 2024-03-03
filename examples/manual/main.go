package main

import (
	"fmt"
	"os"
	"strings"

	gp "github.com/LordOfTrident/greasepaint"
)

type page struct {
	name, library, codeblock, desc, example, output string
}

var pages = []page{
	page{
		name:      "add",
		library:   "math",
		codeblock: "int add(int a, int b)",
		desc:      "Add two integers 'a' and 'b' together and return the result.",
		example:   "math::add(5, 2)",
		output:    "7",
	},
	page{
		name:      "sub",
		library:   "math",
		codeblock: "int sub(int a, int b)",
		desc:      "Subtract integer 'b' from integer 'a' and return the result.",
		example:   "math::sub(5, 2)",
		output:    "3",
	},
	page{
		name:      "repeat",
		library:   "strings",
		codeblock: "string repeat(string str, int n)",
		desc:      "Repeat string 'str' 'n' times and return the result.",
		example:   "strings::repeat('Hello ', 3)",
		output:    "HelloHelloHello",
	},
	page{
		name:      "substr",
		library:   "strings",
		codeblock: "string substr(string str, int pos, int len)",
		desc:      "Returns a portion of 'str' starting from 'pos' with length of 'len'.",
		example:   "strings::substr('Hello, world!', 7, 5)",
		output:    "world",
	},
}

func printSeparator() {
	fmt.Println(styleSeparator(gp.Line("⎯ ", -3)))
}

func printTitle(title string) {
	fmt.Println()
	fmt.Println(styleTitle(title))
	printSeparator()
}

func highlight(text string, keywords []string) string {
	for _, keyword := range keywords {
		text = strings.ReplaceAll(text, keyword, styleKeyword(keyword))
	}

	return text
}

func printPage(page page) {
	codeblock := highlight(page.codeblock, []string{"int", "string"})

	printTitle(page.library + "/" + page.name)
	fmt.Println(stylePageText(styleCode(codeblock)))
	fmt.Println()
	fmt.Println(stylePageText(page.desc))

	printTitle("Example")
	fmt.Println(stylePageText(styleCode(page.example)))
	fmt.Println()
	fmt.Println(stylePageText("Result: ") + styleCode(page.output))
}

func printTabs(currentPage int) {
	var tabs []string
	for i, page := range pages {
		if i == currentPage {
			tabs = append(tabs, styleTabActive(page.name))
		} else {
			tabs = append(tabs, styleTabInactive(page.name))
		}
	}

	fmt.Println(styleTabs(tabs...))
}

func printHeader() {
	fmt.Println(styleHeader("Mylang manual"))
	fmt.Println()
	fmt.Println(gp.Center("Welcome to the Mylang programming language manual page!"))
	fmt.Println()
}

func printUsage() {
	printSeparator()
	fmt.Println(styleUsageTitle("Usage:") + " manual [PAGE]")
	fmt.Println(styleUsageTitle("Pages:"))
	for _, page := range pages {
		fmt.Println(gp.Gap(4) + page.name)
	}
}

func printError(msg string) {
	printSeparator()
	fmt.Println(styleError(msg))
}

func main() {
	printHeader()

	if len(os.Args) <= 1 {
		printUsage()
	} else {
		currentPage := -1
		for i, page := range pages {
			if page.name == os.Args[1] {
				currentPage = i
				break
			}
		}
		if currentPage == -1 {
			printError("Page '" + os.Args[1] + "' not found")
		} else {
			printTabs(currentPage)
			printPage(pages[currentPage])
		}
	}

	fmt.Println()
}

<a name="readme-top"></a>
<div align="center">
	<a href="./LICENSE">
		<img alt="License" src="https://img.shields.io/badge/license-MIT-e8415e?style=for-the-badge">
	</a>
	<a href="https://github.com/LordOfTrident/greasepaint/graphs/contributors">
		<img alt="Contributors" src="https://img.shields.io/github/contributors/LordOfTrident/greasepaint?style=for-the-badge&color=f36a3b">
	</a>
	<a href="https://github.com/LordOfTrident/greasepaint/stargazers">
		<img alt="Stars" src="https://img.shields.io/github/stars/LordOfTrident/greasepaint?style=for-the-badge&color=efb300">
	</a>
	<a href="https://github.com/LordOfTrident/greasepaint/issues">
		<img alt="Issues" src="https://img.shields.io/github/issues/LordOfTrident/greasepaint?style=for-the-badge&color=0fae5e">
	</a>
	<a href="https://github.com/LordOfTrident/greasepaint/pulls">
		<img alt="Pull requests" src="https://img.shields.io/github/issues-pr/LordOfTrident/greasepaint?style=for-the-badge&color=4f79e4">
	</a>
	<br><br><br>
	<img src="./res/logo.png" width="70%">
	<h1 align="center">greasepaint</h1>
	<p align="center">💅 Go library for convenient CLI styling 🤡</p>
	<p align="center">
		<a href="#documentation">Documentation</a>
		·
		<a href="https://github.com/LordOfTrident/greasepaint/issues">Report Bug</a>
		·
		<a href="https://github.com/LordOfTrident/greasepaint/issues">Request Feature</a>
	</p>
	<br>
</div>

<details>
	<summary>Table of contents</summary>
	<ul>
		<li><a href="#introduction">Introduction</a></li>
		<li><a href="#usage">Usage</a></li>
		<li><a href="#example">Example</a></li>
		<li><a href="#documentation">Documentation</a></li>
		<li><a href="#bugs">Bugs</a></li>
	</ul>
</details>

## Introduction
A convenient CLI output styling library for Go which also handles omitting ansi sequences when
redirected to a file for sane output.

## Usage
Add this package to your project
```
$ go get github.com/LordOfTrident/greasepaint
```

And import it
```go
package main

import (
	"fmt"

	gp "github.com/LordOfTrident/greasepaint"
)

func main() {
	fmt.Println(gp.Center("Hello, world!"))
}
```

## Example
You can find examples in the [examples](./examples) folder.

## Documentation
Coming soon.

## Bugs
If you find any bugs, please, [create an issue and report them](https://github.com/LordOfTrident/greasepaint/issues).

<br>
<h1></h1>
<br>

<div align="center">
	<a href="https://go.dev/">
		<img alt="Go" src="https://img.shields.io/badge/Go-007d9c?style=for-the-badge&logo=go&logoColor=white">
	</a>
	<a href="https://github.com/lordoftrident/ansi-go">
		<img alt="ansi-go" src="https://img.shields.io/badge/ANSI, Go!-9f7279?style=for-the-badge&logoColor=white">
	</a>
	<p align="center">Made with ❤️ love</p>
</div>

<p align="right">(<a href="#readme-top">Back to top</a>)</p>

package citeproc_js_go

import (
	_ "embed"
)

//go:embed ieee.csl
var ieeeCsl string

//go:embed locales-en-US.xml
var enUsLocale string

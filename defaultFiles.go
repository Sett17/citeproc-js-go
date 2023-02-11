package citeproc_js_go

import "os"

// ieeeCslContent returns the contents of the file "ieee.csl" as a string.
// If there is an error opening the file, an error is returned.
func ieeeCslContent() (string, error) {
	buf, err := os.ReadFile("ieee.csl")
	return string(buf), err
}

// localesEnUsXmlContent returns the contents of the file "locales-en-US.xml" as a string.
// If there is an error opening the file, an error is returned.
func localesEnUsXmlContent() (string, error) {
	buf, err := os.ReadFile("locales-en-US.xml")
	return string(buf), err
}

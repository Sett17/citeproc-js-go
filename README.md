# citeproc-js-go

citeproc-js-go is a wrapper for selected features of the citeproc-js library, running on Goja, a JavaScript interpreter for Go. The project provides a convenient and efficient way to use citeproc-js in a Go environment.

## Getting Started


You can use `go get` to download citeproc-js-go and its dependencies:

```bash
go get github.com/sett17/citeproc-js-go
```

## Usage

To use citeproc-js-go, you will need to provide it with your own citeproc-js library and style file. You can find a comprehensive list of citation styles and the corresponding style files [here](https://github.com/citation-style-language/styles).

Once you have your citeproc-js library and style file, you can use the following code to initialize citeproc-js-go:

```go
package main

import (
	"fmt"
	"github.com/sett17/citeproc-js-go/csljson"
	citeproc "github.com/sett17/citeproc-js-go"
	"io"
)

func main() {
	// Create a new Citeproc session
	session := citeproc.NewSession()

	// If no files or content is set manually, ieee.csl and locale-en-US.xml will be used by default
	// Load the CSL file and the locale file
	err := session.SetCslFile("path/to/csl/file.csl")
	if err != nil {
		fmt.Println("Error loading CSL file:", err)
		return
	}
	err = session.SetLocaleFile("path/to/locale/file.xml")
	if err != nil {
		fmt.Println("Error loading locale file:", err)
		return
	}

	// Initialize the Citeproc session
	err = session.Init()
	if err != nil {
		fmt.Println("Error initializing Citeproc session:", err)
		return
	}

	// Define the citation items
	var items []csljson.Item

	//TODO: Add items to the session

	// Add all items to the session
	err = session.AddCitation(items...)

	cluster := make([]csljson.Item, 0)
	//TODO add items to the cluster
	
	// Process a citation cluster
	citationCluster, err := session.ProcessCitationCluster(cluster...)
	if err != nil {
		fmt.Println("Error processing citation cluster:", err)
		return
	}

	// Print the resulting citation (this is what goes into the text, e.g. [1], [4])
	fmt.Println(citationCluster)

	// Make bibliography
	bibliography, err := session.MakeBibliography()
	if err != nil {
		fmt.Println("Error making bibliography:", err)
		return
	}

	// Print the resulting bibliography
	fmt.Println(bibliography)
}
```

Replace `path/to/citeproc.js` and `path/to/style.csl` with the actual paths to your citeproc-js library and style file, respectively.


## Contributions

If you would like to contribute to citeproc-js-go, please fork the repository and create a pull request with your changes. Your contributions are always welcome!

## License

citeproc-js-go is licensed under the [MIT License](https://github.com/sett17/citeproc-js-go/blob/main/LICENSE).
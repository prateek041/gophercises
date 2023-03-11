// we will have a section for page creation.
// Then we will have a section for parsing json in the form of a struct
// eventually we will render the pages using page.execute.
// After the templates have been created, create a web handler for handling the requests.
// Initially the starting of the page will be shown and as the user selects a link, render the new page and route the
// user to the new page.

package main

import (
	"github.com/prateek041/gophercises/coya/page"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("template.html")
	if err != nil {
		log.Fatal("error in opening the template file")
	}

	// convert the byte data of the file into string
	str := string(f)
	// create a template
	template := page.CreateTemplate(str)
	_ = template
}

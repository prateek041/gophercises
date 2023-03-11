package page

import (
	"html/template"
	"log"
)

func CreateTemplate(tempDef string) *template.Template {
	temp, err := template.New("storyTemplate").Parse(tempDef)

	if err != nil {
		log.Fatal("error in creating page")
	}

	return temp
}

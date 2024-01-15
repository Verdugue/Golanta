package temps

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func IniTemps() {

	temp, errTemp := template.ParseGlob("./temps/*.html")
	if errTemp != nil {
		fmt.Println("Error template:", errTemp)
		os.Exit(1)
	}
	Temp = temp

}

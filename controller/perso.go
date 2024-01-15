package backend

import (
	"fmt"
	"net/http"
	"os"
	Read "piscine/backend"
	InitStruct "piscine/database"
	InitTemp "piscine/temps"
)

var err error

func Index(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "index", nil)
}

func Save(w http.ResponseWriter, r *http.Request) {

	InitStruct.LstPersonne, err = Read.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	//Prend les valeurs demandés
	InitStruct.Section.Nom = r.FormValue("Nom")
	InitStruct.Section.Prenom = r.FormValue("Prenom")
	InitStruct.Section.Surnom = r.FormValue("Surnom")
	InitStruct.Section.Categorie = r.FormValue("Categorie")
	InitStruct.Section.DateNaissance = r.FormValue("DateNaissance")
	InitStruct.Section.Sexe = r.FormValue("Sexe")

	InitStruct.LstPersonne = append(InitStruct.LstPersonne, InitStruct.Section)
	Read.EditJSON(InitStruct.LstPersonne) //Met les données dans le JSON
	http.Redirect(w, r, "/display", http.StatusMovedPermanently)
}

func Display(w http.ResponseWriter, r *http.Request) {
	InitStruct.LstPersonne, err = Read.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	InitTemp.Temp.ExecuteTemplate(w, "display", InitStruct.LstPersonne)
}

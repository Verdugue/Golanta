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
	InitStruct.Section.Pouvoir = r.FormValue("Pouvoir")
	InitStruct.Section.DateNaissance = r.FormValue("DateNaissance")
	InitStruct.Section.Sexe = r.FormValue("Sexe")
	InitStruct.Section.ImageProfil = r.FormValue("ImageProfil")

	InitStruct.LstPersonne = append(InitStruct.LstPersonne, InitStruct.Section)
	Read.EditJSON(InitStruct.LstPersonne) //Met les données dans le JSON
	http.Redirect(w, r, "/display", http.StatusMovedPermanently)
}

func Display(w http.ResponseWriter, r *http.Request) {
	// Chargement des données
	InitStruct.LstPersonne, err = Read.ReadJSON()
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	// Assurez-vous que les données sont passées ici
	InitTemp.Temp.ExecuteTemplate(w, "display", InitStruct.LstPersonne)
}

func DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	// Récupérez le nom du personnage à supprimer depuis la requête
	nomPersonnage := r.FormValue("nomPersonnage")

	// Supprimez le personnage de la liste
	for i, personnage := range InitStruct.LstPersonne {
		if personnage.Nom == nomPersonnage {
			// Supprimez le personnage de la liste
			InitStruct.LstPersonne = append(InitStruct.LstPersonne[:i], InitStruct.LstPersonne[i+1:]...)
			break
		}
	}

	// Mettez à jour le fichier JSON
	Read.EditJSON(InitStruct.LstPersonne)

	// Redirigez l'utilisateur vers la page "display"
	http.Redirect(w, r, "/display", http.StatusSeeOther)
}

func Main(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "main", nil)
}

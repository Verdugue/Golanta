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

func Modifier(w http.ResponseWriter, r *http.Request) {
	// Récupérez l'ID du personnage à partir des paramètres de l'URL
	nom := r.URL.Query().Get("Nom")
	// Recherchez le personnage correspondant dans votre liste
	var personnageAModifier InitStruct.Personne

	fmt.Println(nom)

	for _, p := range InitStruct.LstPersonne {
		if p.Nom == nom {
			personnageAModifier = p
			break
		}
	}
	fmt.Println(nom)
	// Vérifiez si le personnage a été trouvé
	if personnageAModifier.Nom == "" {

		http.Error(w, "Personnage introuvable", http.StatusNotFound)
		return
	}

	tmplData := struct {
		Nom           string
		Prenom        string
		DateNaissance string
		// Ajoutez d'autres champs ici en fonction de votre structure de données
	}{
		Nom:           personnageAModifier.Nom,
		Prenom:        personnageAModifier.Prenom,
		DateNaissance: personnageAModifier.DateNaissance,
		// Remplacez les données fictives par les données réelles du personnage
	}
	fmt.Println(tmplData)
	// Utilisez votre moteur de templates existant pour rendre la page de modification (modif.html)
	InitTemp.Temp.ExecuteTemplate(w, "modif", tmplData)
}

func ModifierDonneesPersonnage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupérez les données du formulaire
	nom := r.FormValue("nom")
	prenom := r.FormValue("prenom")
	dateNaissance := r.FormValue("dateNaissance")

	// Mettez à jour les données du personnage dans votre liste en utilisant l'ID
	for i, p := range InitStruct.LstPersonne {
		if p.Nom == nom { // Vous pouvez utiliser un identifiant unique à la place du nom
			InitStruct.LstPersonne[i].Nom = nom
			InitStruct.LstPersonne[i].Prenom = prenom
			InitStruct.LstPersonne[i].DateNaissance = dateNaissance
			// Mettez à jour les autres champs ici en fonction de votre structure de données
			break
		}
	}

	// Mettez à jour le fichier JSON avec les modifications
	Read.EditJSON(InitStruct.LstPersonne)

	// Redirigez l'utilisateur vers la page d'affichage après avoir enregistré les modifications
	http.Redirect(w, r, "/display", http.StatusSeeOther)
}

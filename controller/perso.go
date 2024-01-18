package backend

import (
	"fmt"
	"net/http"
	"os"
	"github.com/google/uuid"
	Read "piscine/backend"
	InitStruct "piscine/database"
	InitTemp "piscine/temps"
)

var err error
var personnageAModifier InitStruct.Personne

func Index(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "index", nil)
}

func Save(w http.ResponseWriter, r *http.Request) {

	InitStruct.LstPersonne, err = Read.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	newID := uuid.New().String() // Assurez-vous d'avoir importé "github.com/google/uuid"
    InitStruct.Section.ID = newID
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
    id := r.FormValue("ID")
    fmt.Println("ID reçu dans Modifier:", id) 

	if id == "" {
        fmt.Println("Aucun ID reçu")
        http.Error(w, "Aucun ID reçu", http.StatusBadRequest)
        return
    }

    var personnageAModifier InitStruct.Personne
    trouve := false
    for _, p := range InitStruct.LstPersonne {
        if p.ID == id {
            personnageAModifier = p
            trouve = true
            break
        }
    }

    if !trouve {
        fmt.Println("Personnage introuvable avec l'ID:", id)
        http.Error(w, "Personnage introuvable", http.StatusNotFound)
        return
    }

    // Envoyez personnageAModifier directement au template
    InitTemp.Temp.ExecuteTemplate(w, "modif", personnageAModifier)
}


func ModifierDonneesPersonnage(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    // Récupérez les données du formulaire
	id := r.FormValue("ID")
    nom := r.FormValue("nom")
    prenom := r.FormValue("prenom")
    dateNaissance := r.FormValue("dateNaissance")
    Pouvoir := r.FormValue("Pouvoir")
    Sexe := r.FormValue("Sexe")          
    ImageProfil := r.FormValue("ImageProfil")  

	fmt.Println("Nom reçu:", nom)
    fmt.Println("Prénom reçu:", prenom)
    fmt.Println("Date de Naissance reçue:", dateNaissance)
	fmt.Println("Pouvoir reçu:", Pouvoir)
	fmt.Println("Sexe reçu:", Sexe)
	fmt.Println("Image de profil reçu:", ImageProfil)
	// Vérifiez si les données sont correctes


    trouve := false
    for i, p := range InitStruct.LstPersonne {
        if p.ID == id {
            trouve = true
            InitStruct.LstPersonne[i].Prenom = prenom
            InitStruct.LstPersonne[i].DateNaissance = dateNaissance
            InitStruct.LstPersonne[i].Pouvoir = Pouvoir
            InitStruct.LstPersonne[i].Sexe = Sexe
            InitStruct.LstPersonne[i].ImageProfil = ImageProfil

            // Mettez à jour le fichier JSON et redirigez
            Read.EditJSON(InitStruct.LstPersonne)
            http.Redirect(w, r, "/display", http.StatusSeeOther)
            return
        }
    }

    if !trouve {
        http.Error(w, "Personnage non trouvé", http.StatusNotFound)
    }
}


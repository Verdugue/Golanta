package backend

import (
	"encoding/json"
	"fmt"
	"os"
	data "piscine/database"
)

// Fonction pour modifié le JSON
func EditJSON(ModifiedArticle []data.Personne) {

	modifiedJSON, errMarshal := json.Marshal(ModifiedArticle)
	if errMarshal != nil {
		fmt.Println("Error encodage ", errMarshal.Error())
		return
	}

	// Écrire le JSON modifié dans le fichier
	err := os.WriteFile("JSON/bdd.json", modifiedJSON, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié:", err)
		return
	}

}

// Fonction pour mettre le JSON dans une struct
func ReadJSON() ([]data.Personne, error) {
	jsonFile, err := os.ReadFile("JSON/bdd.json")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}

	var jsonData []data.Personne
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}

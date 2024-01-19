package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func UpdateJSON(modifiedPersonne data.Personne) {
	jsonFilePath := "JSON/bdd.json"

	// Read the JSON file
	fileData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a slice of Personne structs
	var personnes []data.Personne
	err = json.Unmarshal(fileData, &personnes)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Find and update the Personne with the given ID
	updated := false
	for i, p := range personnes {
		if p.ID == modifiedPersonne.ID {
			// Update the Personne record
			personnes[i] = modifiedPersonne
			updated = true
			break
		}
	}

	if !updated {
		fmt.Println("No record found with ID:", modifiedPersonne.ID)
		return
	}

	// Marshal the updated slice back to JSON
	updatedData, err := json.Marshal(personnes)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Write the updated JSON data back to the file
	err = ioutil.WriteFile(jsonFilePath, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Record updated successfully.")

}

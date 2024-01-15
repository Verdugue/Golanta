package database

type Personne struct {
	Nom           string `json:"Nom"`
	Prenom        string `json:"Prenom"`
	Surnom        string `json:"Surnom"`
	Categorie     string `json:"Categorie"`
	DateNaissance string `json:"DateNaissance"`
	Pouvoir       string `json:"Pouvoir"`
	Sexe          string `json:"Sexe"`
	ImageProfil   string `json:"ImageProfil"`
}

var bddFileName = "bdd.json"
var Section Personne
var LstPersonne []Personne

package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"html/template"
	"net/http"
)

// Index page pour acceuil du site//
func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var offre entities.OffresGenerales
	var entreprise entities.Entreprises
	var evenement entities.EvenementEntrepriseBd
	if session.Values["loggedIn"] != true {
		offres, _ := userModel.OfrresGeneralesIndex(&offre, &entreprise)
		entreprises, _, _ := userModel.GetEntreprises(&entreprise)
		evenements, _ := userModel.GetEvenements(&evenement)

		data := map[string]interface{}{
			"Offres":      offres,
			"Entreprises": entreprises,
			"Evenements":  evenements,
		}
		temp, _ := template.ParseFiles("views/index.html")
		temp.Execute(w, data)
	} else {
		http.Redirect(w, r, "/acceuil", http.StatusSeeOther)

	}
}

package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"html/template"
	"net/http"
	"strconv"
)

func AllEntreprises(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {

			var entrepises entities.Entreprises
			res, nbEntreprise, _ := userModel.GetAllEntreprises(&entrepises)
			data := map[string]interface{}{
				"Entreprises":   res,
				"NbEntreprises": nbEntreprise,
			}

			temp, _ := template.ParseFiles("views/entreprises.html")
			// fmt.Println(res, "tab")
			temp.Execute(w, data)
		}

	}

}

func DetailEntreprises(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {
			emp := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(emp)
			var detailEntreprise entities.Entreprises
			var entrepises entities.Entreprises
			var offre entities.OffresGenerales
			res, _ := userModel.Entreprise(&detailEntreprise, id)
			offres, _ := userModel.AllOfrresGeneralEntreprise(&offre, &entrepises, id)
			data := map[string]interface{}{
				"Entreprises": res,
				"Offres":      offres,
			}
			temp, _ := template.ParseFiles("views/entreprise-more.html")
			temp.Execute(w, data)
		}

	}

}

//Formulaire connexion recruteur
/*func ContactRecruteur(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/connexion-recruteur.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}*/

// dashboard recruteur après connexion
func DashboardRecruteur(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/dashboard-recruteur.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

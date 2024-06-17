package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"html/template"
	"net/http"
	"strconv"
)

func Offres(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {

			var offres entities.OffresGenerales
			var entreprise entities.Entreprises
			res, _ := userModel.AllOfrresGenerales(&offres, &entreprise)
			temp, _ := template.ParseFiles("views/Offres.html")
			// fmt.Println(res, "tab")
			temp.Execute(w, res)
		}

	}

}

func DetailOffres(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {

			if r.Method == http.MethodGet {
				emp := r.URL.Query().Get("id")
				id, _ := strconv.Atoi(emp)
				var offres entities.OffresEntreprises

				// var entreprise entities.Entreprises
				res, _ := userModel.DetailOffresPoster(&offres, id)
				temp, _ := template.ParseFiles("views/details-offre.html")
				// fmt.Println(res, "tab")
				temp.Execute(w, res)
			}
			//  else if r.Method == http.MethodPost {
			// 	r.ParseForm()
			// 	v := r.Form.Get("offres_id")
			// 	fmt.Println(v)
			// }

		}

	}

}

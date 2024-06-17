package recruteurscontroller

import (
	"digiJob/config"
	recruteursentities "digiJob/entities/recruteursEntities"
	"net/http"
	"text/template"
)

// Après connexion du recruteur

func DashboardRecruteur(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginRecruteurs", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginRecruteurs", http.StatusSeeOther)

		} else {

			var allDigi recruteursentities.UsersDigi
			res, cmpt, _ := userModel.AllDigiCodeuses(&allDigi)
			data := map[string]interface{}{
				"Entreprises": res,
				"Nombre":      cmpt,
			}
			tmpl, err := template.ParseFiles("views/dashboard-recruteur.html")
			if err != nil {
				panic(err)
			}
			tmpl.Execute(w, data)
		}

	}
}

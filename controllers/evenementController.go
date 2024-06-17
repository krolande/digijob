package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func AllEvenement(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {

			var evenement entities.EvenementEntrepriseBd
			res, _ := userModel.GetAllEvenements(&evenement)

			temp, _ := template.ParseFiles("views/evenement.html")
			// fmt.Println(res, "tab")
			temp.Execute(w, res)
		}

	}

}

func DetailEvenements(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {
			emp := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(emp)
			var detailEntreprise entities.EvenementEntrepriseBd
			var entrepises entities.Entreprises
			var offre entities.OffresGenerales
			if r.Method == http.MethodGet {
				res, _ := userModel.DetailEvenements(&detailEntreprise, id)
				offres, _ := userModel.AllOfrresGeneralEntreprise(&offre, &entrepises, id)
				data := map[string]interface{}{
					"Entreprises": res,
					"Offres":      offres,
				}
				temp, _ := template.ParseFiles("views/details-evenement.html")
				temp.Execute(w, data)
			} else {
				if r.Method == http.MethodPost {
					r.ParseForm()
					id_Digi := session.Values["id"]
					//convertir un type interface en string
					uid := fmt.Sprintf("%v", id_Digi)
					uuid, _ := strconv.Atoi(uid)
					event_id := r.Form.Get("evenements_id")
					evenements_id, _ := strconv.Atoi(event_id)
					participant := entities.ParticipantEvent{
						Evenements_id:        evenements_id,
						UsersDigicodeuses_id: uuid,
					}
					_, err := userModel.InsParticipantEvent(participant)
					var message string
					if err != nil {
						fmt.Println(err, "insert participant")
						message = "Echec d'enregistrement"
					} else {
						message = "Enregistrement effectué avec succès"

					}
					//Affichage de page
					res, _ := userModel.DetailEvenements(&detailEntreprise, evenements_id)
					data := map[string]interface{}{
						"Entreprises": res,
						"Retour":      message,
					}
					temp, _ := template.ParseFiles("views/details-evenement.html")
					temp.Execute(w, data)
				}
			}

		}

	}

}

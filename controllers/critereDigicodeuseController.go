package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func CritereDigicodeuse(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var userDataAll entities.UsersDigi
	var competenceDataAll entities.CompetenceCles
	var critereDigicodeuse entities.CritereDigicodeuses
	var barreProgInfoPerso entities.BarreProgressionInfoPerso
	var barreProgCv entities.BarreProgressionCv
	var barreProgCritere entities.BarreProgressionCritere

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			id := session.Values["id"]
			//convertir un type interface en string
			uid := fmt.Sprintf("%v", id)
			uuid, _ := strconv.Atoi(uid)
			if r.Method == http.MethodGet {

				res, _ := userModel.GetUser(&userDataAll, uuid)
				userCompetencesCles, _ := userModel.GetAllCompetenceClesOfUser(&competenceDataAll, uuid)
				userCritere, comptCritere, _ := userModel.GetCritereDigi(&critereDigicodeuse, uuid)
				nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
				nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
				nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
				barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere

				m := map[string]interface{}{
					"Profil":             res,
					"UserCompetenceCles": userCompetencesCles,
					"Critere":            userCritere,
					"ProgressBarre":      barreProgressionTotal,
				}
				v := entities.BarreProgressionCritere{
					NbreCritere: comptCritere,
				}
				_, er := userModel.InsertBarreProgresCritere(v, uuid)
				if er != nil {
					fmt.Println(er)
				}
				temp, _ := template.ParseFiles("views/candidates-mes-criteres.html")
				temp.Execute(w, m)
			} else if r.Method == http.MethodPost {
				r.ParseForm()
				IsnsertCritere := entities.CritereDigicodeuses{
					TypeContrat:         r.Form.Get("type_contrat"),
					Localite:            r.Form.Get("localite"),
					DureeContrat:        r.Form.Get("duree_contrat"),
					Modalite:            r.Form.Get("modalite"),
					MetierRecherche:     r.Form.Get("metier_recherche"),
					UsersDigicodeuse_id: uuid,
				}
				_, err := userModel.InsertOrUpdateCritereDigi(IsnsertCritere, uuid)
				// fmt.Println(err)
				var message string
				if err != nil {
					message = "Echec d'enregitrement:" + message
				} else {
					message = "Critères enregistrés avec succès"
				}

				res, _ := userModel.GetUser(&userDataAll, uuid)
				userCompetencesCles, _ := userModel.GetAllCompetenceClesOfUser(&competenceDataAll, uuid)
				userCritere, _, _ := userModel.GetCritereDigi(&critereDigicodeuse, uuid)
				nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
				nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
				nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
				barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere

				data := map[string]interface{}{
					"Profil":             res,
					"UserCompetenceCles": userCompetencesCles,
					"Critere":            userCritere,
					"ProgressBarre":      barreProgressionTotal,
					"retour":             message,
				}
				// fmt.Println(data)
				temp, _ := template.ParseFiles("views/candidates-mes-criteres.html")
				temp.Execute(w, data)
				// http.Redirect(w, r, "/critereDigicodeuse", http.StatusSeeOther)
			}

		}
	}
}

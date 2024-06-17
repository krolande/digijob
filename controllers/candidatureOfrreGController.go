package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func InsCandidature(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	var userDataAll entities.UsersDigi
	var competenceDataAll entities.CompetenceCles
	var candit entities.CandidatureOffreGenerales
	// var compt entities.CompetenceCles
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {
			id := session.Values["id"]
			//convertir un type interface en string
			uid := fmt.Sprintf("%v", id)
			uuid, _ := strconv.Atoi(uid)
			if r.Method == http.MethodGet {
				res, _ := userModel.GetUser(&userDataAll, uuid)
				//connaitre le lien que nous recuperons
				// fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)
				emp := r.URL.Query().Get("id")
				data := map[string]interface{}{
					"Profil":    res,
					"Offres_id": emp,
				}
				temp, _ := template.ParseFiles("views/formulaire-candidature.html")
				temp.Execute(w, data)

			} else {
				if r.Method == http.MethodPost {
					r.ParseForm()
					cv := uploadFile("cvCandidature", w, r)
					offres := r.Form.Get("offres_id")
					offres_id, _ := strconv.Atoi(offres)
					candidature := entities.CandidatureOffreGenerales{
						Cv:                   cv,
						Motivation:           r.Form.Get("motivation"),
						OffresGenerales_id:   offres_id,
						UsersDigicodeuses_id: uuid,
					}

					_, erreur := userModel.InsertCandidatureOffreG(candidature)

					// res, _:= userModel.GetAllCompetenceClesOfUser(&compt, uuid)
					// fmt.Println()

					//Filtrage automatique des candidates pour une offre en fonction des competences clées avant l'enregistrement de la statut de la candidature
					libelleDigi, _ := userModel.GetLibelleCompetenceCle(&competenceDataAll, uuid)
					libelleOffreG, _ := userModel.GetCompetenceClesOffreG(&competenceDataAll, offres_id)
					//Comparaison des CompetenceCles recherchées
					candidature_id, _ := userModel.GetIdCandidature(&candit, offres_id, uuid)
					if len(libelleOffreG) != 0 {
						nbreCompetenceCles := 0
						for a := 0; a < len(libelleOffreG); a++ {
							for b := 0; b < len(libelleDigi); b++ {
								if libelleOffreG[a] == libelleDigi[b] {
									nbreCompetenceCles++
								}
							}
						}

						if nbreCompetenceCles == len(libelleOffreG) {
							//on fait enregistrement de statutAuto
							satutCandidature := entities.StatutCandidature{
								SatutAuto:                    "Approuvée",
								SatutTalentManager:           "En attente",
								SatutEntreprise:              "En attente",
								CandidatureOffreGenerales_id: candidature_id,
							}
							_, err := userModel.InsStatutCandidatureOG(satutCandidature)

							if err != nil {
								fmt.Println(err)
							}
						} else {
							//on fait enregistrement de statutAuto, talentManager et Recruteur
							//Pour pouvoir enregistrer les statuts il faux que je recupère les données de la candidature
							satutCandidature := entities.StatutCandidature{
								SatutAuto:                    "Réjetée",
								SatutTalentManager:           "Réjetée",
								SatutEntreprise:              "Réjetée",
								CandidatureOffreGenerales_id: candidature_id,
							}
							_, err := userModel.InsStatutCandidatureOG(satutCandidature)

							if err != nil {
								fmt.Println(err)
							}
						}
					} else {
						satutCandidature := entities.StatutCandidature{
							SatutAuto:                    "Approuvée",
							SatutTalentManager:           "En attente",
							SatutEntreprise:              "En attente",
							CandidatureOffreGenerales_id: candidature_id,
						}
						_, err := userModel.InsStatutCandidatureOG(satutCandidature)

						if err != nil {
							fmt.Println(err)
						}
					}

					if erreur != nil {
						fmt.Println(erreur)
					}
					// var message string
					// if erreur != nil {
					// 	message = "Echec d'enregitrement:" + message
					// } else {
					// 	message = "Enregistrés effectuées avec succès"
					// }

					// id := session.Values["id"]
					// //convertir un type interface en string
					// uid := fmt.Sprintf("%v", id)
					// uuid, _ := strconv.Atoi(uid)
					// var candidatures entities.AllCandidaturesDigi

					// res, nbre, _ := userModel.GetAllCandidatureDigi(&candidatures, uuid)
					// _, nbrecandSucces, _ := userModel.AllCandidatureDigiPresel(&candidatures, uuid)
					// _, nbrecandFailed, _ := userModel.AllCandidatureDigiRejet(&candidatures, uuid)

					// data := map[string]interface{}{
					// 	"AllCandidature": nbre,
					// 	"NbrecandSucces": nbrecandSucces,
					// 	"NbrecandFailed": nbrecandFailed,
					// 	"Candidature":    res,
					// 	"Retour":         message,
					// }
					// temp, _ := template.ParseFiles("views/candidatures.html")
					// temp.Execute(w, data)
					http.Redirect(w, r, "/allCandidature", http.StatusSeeOther)
				}
			}

		}

	}

}

//Lorsque je fais la redirection vers le tableau de candidature c'est cette fonction j'appelle
func AllCandidatureDigi(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var candidatures entities.AllCandidaturesDigi
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {
			id := session.Values["id"]
			//convertir un type interface en string
			uid := fmt.Sprintf("%v", id)
			uuid, _ := strconv.Atoi(uid)
			if r.Method == http.MethodGet {
				res, nbre, _ := userModel.GetAllCandidatureDigi(&candidatures, uuid)
				_, nbrecandSucces, _ := userModel.AllCandidatureDigiPresel(&candidatures, uuid)
				_, nbrecandFailed, _ := userModel.AllCandidatureDigiRejet(&candidatures, uuid)

				temp, _ := template.ParseFiles("views/candidatures.html")
				data := map[string]interface{}{
					"AllCandidature": nbre,
					"NbrecandSucces": nbrecandSucces,
					"NbrecandFailed": nbrecandFailed,
					"Candidature":    res,
				}
				temp.Execute(w, data)
			} else {

				if r.Method == http.MethodPost {
					r.ParseForm()
					cv := uploadFile("cvCandidature", w, r)
					offres := r.Form.Get("offres_id")
					offres_id, _ := strconv.Atoi(offres)
					candidature := entities.CandidatureOffreGenerales{
						Cv:                   cv,
						Motivation:           r.Form.Get("motivation"),
						OffresGenerales_id:   offres_id,
						UsersDigicodeuses_id: uuid,
					}
					_, err := userModel.UpdateCandidature(candidature)
					if err != nil {
						fmt.Println(err)
					}

					res, nbre, _ := userModel.GetAllCandidatureDigi(&candidatures, uuid)
					_, nbrecandSucces, _ := userModel.AllCandidatureDigiPresel(&candidatures, uuid)
					_, nbrecandFailed, _ := userModel.AllCandidatureDigiRejet(&candidatures, uuid)

					temp, _ := template.ParseFiles("views/candidatures.html")
					data := map[string]interface{}{
						"AllCandidature": nbre,
						"NbrecandSucces": nbrecandSucces,
						"NbrecandFailed": nbrecandFailed,
						"Candidature":    res,
					}
					temp.Execute(w, data)
				}
			}

		}

	}

}

func DeleteCandidature(w http.ResponseWriter, r *http.Request) {
	emp := r.URL.Query().Get("idC")
	id, _ := strconv.Atoi(emp)
	userModel.DeleteCandidatureOffreG(id)
	userModel.DeleteStatutCandidature(id)
	// userModel.DeleteExperience(uuid)
	http.Redirect(w, r, "/allCandidature", http.StatusSeeOther)
}

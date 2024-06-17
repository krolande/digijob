package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

///////////////****************** MON CV ******************** ///////////////

func MonCV(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var userDataAll entities.UsersDigi
	var competenceDataAll entities.CompetenceCles
	var langues entities.Langues
	var formation entities.Formation
	var experience entities.ExperienceProfessionnelle
	var competenceProfessionnelle entities.CompetenceProfessionnelle
	var prixCertifications entities.PrixCertification
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
				langueAll, _ := userModel.GetAllLangues(&langues)
				langueUser, comptLangue, _ := userModel.GetLanguesUser(&langues, uuid)
				resFormation, comptFormation, _ := userModel.ReadForamtion(&formation, uuid)
				resExperience, comptExperience, _ := userModel.ReadExperience(&experience, uuid)
				resCompetencePro, comptCompetencePro, _ := userModel.GetCompetenceProfessionnelle(&competenceProfessionnelle, uuid)
				resPrixCertification, comptPrixCertificat, _ := userModel.ReadPrixCertification(&prixCertifications, uuid)
				nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
				nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
				nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
				barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere
				data := map[string]interface{}{
					"Profil":                    res,
					"UserCompetenceCles":        userCompetencesCles,
					"Langues":                   langueAll,
					"LanguesUser":               langueUser,
					"Formations":                resFormation,
					"ExperienceProfessionnelle": resExperience,
					"CompetenceProfessionnelle": resCompetencePro,
					"PrixCertification":         resPrixCertification,
					"ProgressBarre":             barreProgressionTotal,
				}
				v := entities.BarreProgressionCv{
					NbreLangue:         comptLangue,
					NbreExpeProf:       comptExperience,
					NbrePrixCertificat: comptPrixCertificat,
					NbreComptProf:      comptCompetencePro,
					NbreFormation:      comptFormation,
				}
				_, er := userModel.InsertBarreProgressionCv(v, uuid)
				fmt.Println(er)
				if er != nil {
					fmt.Println(er)
				}
				temp, _ := template.ParseFiles("views/candidates-mon-cv.html")
				temp.Execute(w, data)
			} else if r.Method == http.MethodPost {
				r.ParseForm()
				var message string
				/// ******************* LANGUES ******************* ///
				if r.Form.Get("typeForm") == "langues" {
					langues := entities.UserLangues{
						Libelle: r.Form["langue"],
					}
					var langueDb entities.LangueUsersDigicodeuses

					_, err := userModel.InsertOrUpdateLangue(langues.Libelle, &langueDb, uuid)
					if err != nil {
						fmt.Println(err)
						message = "Echec d'enregitrement:"
					} else {
						message = "Langue(s) enregistrée(s) avec succès"
					}

				}

				/// ******************* FORMATION ******************* ///
				if r.Form.Get("typeForm") == "formationsInsert" {
					// formationId, _ := strconv.Atoi(r.Form.Get("id"))
					insertFormation := entities.Formation{
						MoisDebut:      r.Form.Get("mois_debut"),
						MoisFin:        r.Form.Get("mois_fin"),
						AnneeDebut:     r.Form.Get("annee_debut"),
						AnneeFin:       r.Form.Get("annee_fin"),
						Now:            r.Form.Get("now"),
						TitreFormation: r.Form.Get("titre_formation"),
						NomEcole:       r.Form.Get("nom_ecole"),
						Description:    r.Form.Get("description"),
					}
					_, err := userModel.InsertFormation(insertFormation, uuid)
					if err != nil {
						fmt.Println(err)
						message = "Echec d'enregitrement"
					} else {
						message = "Formation ajoutée avec succès"
					}

				}

				if r.Form.Get("typeForm") == "formationsEdit" {
					formationId, _ := strconv.Atoi(r.Form.Get("id"))
					insertFormation := entities.Formation{
						Id:             formationId,
						MoisDebut:      r.Form.Get("mois_debut"),
						MoisFin:        r.Form.Get("mois_fin"),
						AnneeDebut:     r.Form.Get("annee_debut"),
						AnneeFin:       r.Form.Get("annee_fin"),
						Now:            r.Form.Get("now"),
						TitreFormation: r.Form.Get("titre_formation"),
						NomEcole:       r.Form.Get("nom_ecole"),
						Description:    r.Form.Get("description"),
					}
					_, err := userModel.UpdateFormation(insertFormation)
					// fmt.Println(err, "update")
					if err != nil {
						fmt.Println(err)
						message = "Echec d'enregitrement"
					} else {
						message = "Formation modifiée avec succès"
					}
				}

				/// ******************* EXPERIENCE PROFESSIONNELLES ******************* ///

				if r.Form.Get("typeForm") == "experienceInsert" {
					// formationId, _ := strconv.Atoi(r.Form.Get("id"))
					insertExperience := entities.ExperienceProfessionnelle{
						MoisDebut:       r.Form.Get("mois_debut"),
						MoisFin:         r.Form.Get("mois_fin"),
						AnneeDebut:      r.Form.Get("annee_debut"),
						AnneeFin:        r.Form.Get("annee_fin"),
						Now:             r.Form.Get("now"),
						TitreExperience: r.Form.Get("titre_experience"),
						NomEntreprise:   r.Form.Get("nom_entreprise"),
						TypeContrat:     r.Form.Get("type_contrat"),
						Description:     r.Form.Get("description"),
					}
					_, err := userModel.InsertExpProfessionnelle(insertExperience, uuid)
					if err != nil {
						fmt.Println(err)
						message = "Echec d'enregistrement"
					} else {
						message = "Experience professionnelle ajoutée avec succès"
					}

				}

				if r.Form.Get("typeForm") == "experienceUpdate" {
					formationId, _ := strconv.Atoi(r.Form.Get("id"))
					updateExperience := entities.ExperienceProfessionnelle{
						Id:              formationId,
						MoisDebut:       r.Form.Get("mois_debut"),
						MoisFin:         r.Form.Get("mois_fin"),
						AnneeDebut:      r.Form.Get("annee_debut"),
						AnneeFin:        r.Form.Get("annee_fin"),
						Now:             r.Form.Get("now"),
						TitreExperience: r.Form.Get("titre_experience"),
						NomEntreprise:   r.Form.Get("nom_entreprise"),
						TypeContrat:     r.Form.Get("type_contrat"),
						Description:     r.Form.Get("description"),
					}
					_, err := userModel.UpdateExperience(updateExperience)
					if err != nil {
						fmt.Println(err)
						message = "Echec de modification"
					} else {
						message = "Experience professionnelle modifiée avec succès"
					}
				}

				// ************ COMPETENCE PRofessionnelles ************ //
				if r.Form.Get("typeForm") == "competencePro" {
					tab := r.Form["id"]
					tabPourcentage := r.Form["pourcentage"]
					var tabInt []int
					for i := 0; i < len(tab); i++ {
						m, _ := strconv.Atoi(tab[i])
						tabInt = append(tabInt, m)
					}
					var tabPourcentageInt []int
					for a := 0; a < len(tabPourcentage); a++ {
						v, _ := strconv.Atoi(tabPourcentage[a])
						tabPourcentageInt = append(tabPourcentageInt, v)
					}
					insertCompetencePro := entities.CompetenceProfessionnelles{
						Id:          tabInt,
						Titre:       r.Form["titre"],
						Pourcentage: tabPourcentageInt,
					}
					for i := 0; i < len(insertCompetencePro.Titre); i++ {
						if insertCompetencePro.Pourcentage[i] <= 100 {
							competencePro := entities.CompetenceProfessionnelle{
								Id:          insertCompetencePro.Id[i],
								Titre:       insertCompetencePro.Titre[i],
								Pourcentage: insertCompetencePro.Pourcentage[i],
							}
							_, err := userModel.InsertCompetenceProfessionnelle(competencePro, &competencePro, uuid, competencePro.Id)
							if err != nil {
								fmt.Println(err)
								message = "Echec d'enregistrement"
							} else {
								message = "Compétence(s) professionnelle(s) enregistrée avec succès"
							}
						}

					}

				}

				// ********************* PRIX ERTIFICAT ***************** //
				if r.Form.Get("typeForm") == "insertPrixCertificat" {
					insertPrixCertificat := entities.PrixCertification{
						TitrePrix:           r.Form.Get("titre_prix"),
						NomEtablissement:    r.Form.Get("nom_etablissement"),
						DateObtention:       r.Form.Get("date_obtention"),
						Description:         r.Form.Get("description"),
						UsersDigicodeuse_id: uuid,
					}
					_, err := userModel.InsertPrixCertification(insertPrixCertificat)
					if err != nil {
						fmt.Println(err)
						message = "Echec d'engregistrement"
					} else {
						message = "Prix et certificat avec succès"
					}

				}
				if r.Form.Get("typeForm") == "updatePrixCertificat" {
					titrePrixId, _ := strconv.Atoi(r.Form.Get("id"))
					updatePrixCertificat := entities.PrixCertification{
						Id:               titrePrixId,
						TitrePrix:        r.Form.Get("titre_prix"),
						NomEtablissement: r.Form.Get("nom_etablissement"),
						DateObtention:    r.Form.Get("date_obtention"),
						Description:      r.Form.Get("description"),
					}
					_, err := userModel.UpdatePrixCertification(updatePrixCertificat)
					if err != nil {
						fmt.Println(err)
						message = "Echec de modification"
					} else {
						message = "Prix et certificat modifié avec succès"
					}

				}

				res, _ := userModel.GetUser(&userDataAll, uuid)
				userCompetencesCles, _ := userModel.GetAllCompetenceClesOfUser(&competenceDataAll, uuid)
				langueAll, _ := userModel.GetAllLangues(&langues)
				langueUser, _, _ := userModel.GetLanguesUser(&langues, uuid)
				resFormation, _, _ := userModel.ReadForamtion(&formation, uuid)
				resExperience, _, _ := userModel.ReadExperience(&experience, uuid)
				resCompetencePro, _, _ := userModel.GetCompetenceProfessionnelle(&competenceProfessionnelle, uuid)
				resPrixCertification, _, _ := userModel.ReadPrixCertification(&prixCertifications, uuid)
				nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
				nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
				nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
				barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere
				data := map[string]interface{}{
					"Profil":                    res,
					"UserCompetenceCles":        userCompetencesCles,
					"Langues":                   langueAll,
					"LanguesUser":               langueUser,
					"Formations":                resFormation,
					"ExperienceProfessionnelle": resExperience,
					"CompetenceProfessionnelle": resCompetencePro,
					"PrixCertification":         resPrixCertification,
					"ProgressBarre":             barreProgressionTotal,
					"retour":                    message,
				}
				temp, _ := template.ParseFiles("views/candidates-mon-cv.html")
				temp.Execute(w, data)
				// http.Redirect(w, r, "/monCv", http.StatusSeeOther)
			}

		}
	}
}

func DeleteFormationCv(w http.ResponseWriter, r *http.Request) {
	emp := r.URL.Query().Get("id")
	uuid, _ := strconv.Atoi(emp)
	userModel.DeleteFormation(uuid)
	http.Redirect(w, r, "/monCv", http.StatusSeeOther)
}

func DeleteExperienceCv(w http.ResponseWriter, r *http.Request) {
	emp := r.URL.Query().Get("id")
	uuid, _ := strconv.Atoi(emp)
	userModel.DeleteExperience(uuid)
	http.Redirect(w, r, "/monCv", http.StatusSeeOther)
}

func DeleteCompetenceProfessionnelle(w http.ResponseWriter, r *http.Request) {
	emp := r.URL.Query().Get("id")
	uuid, _ := strconv.Atoi(emp)
	userModel.DeleteCompetencePro(uuid)
	http.Redirect(w, r, "/monCv", http.StatusSeeOther)
}

func DeletePrixCertifications(w http.ResponseWriter, r *http.Request) {
	emp := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(emp)
	userModel.DeletePrixCertification(id)
	http.Redirect(w, r, "/monCv", http.StatusSeeOther)
}

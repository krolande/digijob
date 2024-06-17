package controllers

import (
	"crypto/rand"
	"digiJob/config"
	"digiJob/entities"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var nbInfoPerso int

// Mettre ajout les informations personnelles de la digicodeuses <<<<<<<<<<<<<< MODIFIER LE PROFIL >>>>>>>>>>>>>>> //
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var userDataAll entities.UsersDigi
	var competenceDataAll entities.CompetenceCles
	var UserCompetenceCles entities.CompetenceCleUsersDigicodeuses
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
				competencesAll, _ := userModel.AllCompetenceCles(&competenceDataAll)
				userCompetencesCles, _ := userModel.GetAllCompetenceClesOfUser(&competenceDataAll, uuid)
				nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
				nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
				nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
				barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere
				m := map[string]interface{}{
					"Profil":             res,
					"Competences":        competencesAll,
					"UserCompetenceCles": userCompetencesCles,
					"ProgressBarre":      barreProgressionTotal,
				}
				temp, _ := template.ParseFiles("views/candidates-edit-mon-profil.html")
				// fmt.Println(err, "resultController")
				temp.Execute(w, m)
			} else if r.Method == http.MethodPost {
				r.ParseForm()
				// fmt.Println(r.PostFormValue("nom"))
				// fmt.Println(r.FormValue("nom"))
				// id, _ := strconv.Atoi(r.Form.Get("uid"))
				cheminCv := uploadFile("cv", w, r)
				cheminPhoto := uploadFile("photo", w, r)
				if cheminPhoto == "" {
					cheminPhoto = r.Form.Get("photoProfil")
				}
				if cheminCv == "" {
					cheminCv = r.Form.Get("pdfCv")
				}
				user := entities.UsersDigicodeuses{
					Nom:             r.Form.Get("nom"),
					Prenoms:         r.Form.Get("prenoms"),
					Email:           r.Form.Get("email"),
					Ville:           r.Form.Get("ville"),
					Pays:            r.Form.Get("pays"),
					DateNaissance:   r.Form.Get("date_naissance"),
					Indicatif:       r.Form.Get("indicatif"),
					NumeroTelephone: r.Form.Get("numero_telephone"),
					Nationalite:     r.Form.Get("nationalite"),
					Cv:              cheminCv,
					Photo:           cheminPhoto,
					TitrePosteCv:    r.Form.Get("titrePosteCv"),
					Porfolio:        r.Form.Get("portfolio"),
					Facebook:        r.Form.Get("facebook"),
					Twitter:         r.Form.Get("twitter"),
					Linkedin:        r.Form.Get("linkedin"),
					Temoignage:      r.Form.Get("temoignage"),
					Description:     r.Form.Get("description"),
					NiveauEtude:     r.Form.Get("niveau_etude"),
					Adresse:         r.Form.Get("adresse"),
					Statut:          r.Form.Get("statut"),
					AnneeExperience: r.Form.Get("annee_experience"),
					// Password:        r.Form.Get("password"),
					Id: uuid,
				}
				userCompetenceCles := entities.CompetencesCles{
					Libelle: r.Form["competenceCles"],
				}

				_, erreur := userModel.CreateUpdateCompetenceCles(userCompetenceCles.Libelle, &UserCompetenceCles, uuid)
				// fmt.Println(valCompetenceCles, "competences")
				if erreur != nil {
					fmt.Println(erreur)
				}

				//cryptage de password
				hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
				user.Password = string(hashPassword)
				//insertion des nouveau données
				_, err := userModel.Update(user)
				// fmt.Println(val, err)
				if err != nil {
					fmt.Println(err)
				}

				// ************************* Pourcentage nbInfoPerso **********************
				if r.Form.Get("nom") != "" {
					nbInfoPerso++
				}

				if r.Form.Get("prenoms") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("email") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("ville") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("pays") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("date_naissance") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("indicatif") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("numero_telephone") != "" {
					nbInfoPerso++
				}
				if cheminCv != "Error Retrieving the File" {
					nbInfoPerso++
				}
				if cheminPhoto != "Error Retrieving the File" {
					nbInfoPerso++
				}
				if r.Form.Get("titrePosteCv") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("portfolio") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("facebook") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("twitter") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("linkedin") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("description") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("niveau_etude") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("adresse") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("statut") != "" {
					nbInfoPerso++
				}
				if r.Form.Get("annee_experience") != "" {
					nbInfoPerso++
				}
				if len(userCompetenceCles.Libelle) != 0 {
					nbInfoPerso++
				}

				v := entities.BarreProgressionInfoPerso{
					NbreInfoPerso: nbInfoPerso,
				}
				_, er := userModel.InsertBarreProgresInfoPerso(v, uuid)
				if err != nil {
					fmt.Println(er)
				}
				nbInfoPerso = 0

				var message string

				if erreur != nil && err != nil {
					message = "Echec d'enregitrement:" + message
				} else {
					message = "Enregistrement réussi"
				}

				res, _ := userModel.GetUser(&userDataAll, uuid)
				competencesAll, _ := userModel.AllCompetenceCles(&competenceDataAll)
				userCompetencesCles, _ := userModel.GetAllCompetenceClesOfUser(&competenceDataAll, uuid)
				nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
				nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
				nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
				barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere
				data := map[string]interface{}{
					"Profil":             res,
					"Competences":        competencesAll,
					"UserCompetenceCles": userCompetencesCles,
					"ProgressBarre":      barreProgressionTotal,
					"retour":             message,
				}
				temp, _ := template.ParseFiles("views/candidates-edit-mon-profil.html")
				temp.Execute(w, data)
				// http.Redirect(w, r, "/Updateprofile", http.StatusSeeOther)

			}
		}

	}

}

///////////////////////////////////<<<<<<<<<<<<<<<<<<<       MON PROFIL 	>>>>>>>>>>>>>>>>>>>>//////////////////////////

func ViewsProfil(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var userDataAll entities.UsersDigi
	var competenceDataAll entities.CompetenceCles
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
			res, _ := userModel.GetUser(&userDataAll, uuid)
			userCompetencesCles, _ := userModel.GetAllCompetenceClesOfUser(&competenceDataAll, uuid)
			nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
			nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
			nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
			barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere

			m := map[string]interface{}{
				"Profil":             res,
				"UserCompetenceCles": userCompetencesCles,
				"ProgressBarre":      barreProgressionTotal,
			}
			temp, _ := template.ParseFiles("views/candidates-mon-profil.html")
			temp.Execute(w, m)

		}
	}
}

func ModifPassword(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var userDataAll entities.UsersDigi
	var competenceDataAll entities.CompetenceCles
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
			res, _ := userModel.GetUser(&userDataAll, uuid)
			if r.Method == http.MethodGet {
				userCompetencesCles, _ := userModel.GetAllCompetenceClesOfUser(&competenceDataAll, uuid)
				nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
				nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
				nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
				barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere

				m := map[string]interface{}{
					"Profil":             res,
					"UserCompetenceCles": userCompetencesCles,
					"ProgressBarre":      barreProgressionTotal,
				}
				temp, _ := template.ParseFiles("views/candidates-change-password.html")
				temp.Execute(w, m)
			} else if r.Method == http.MethodPost {
				r.ParseForm()
				var message error
				// UserInput est dans le controller de authDigicodeusesController.go
				UserInput := &UserInput{
					PasswordActuel: r.Form.Get("password_actuel"),
					Cpassword:      r.Form.Get("cpassword"),
					Password:       r.Form.Get("password"),
				}
				// Vérifie le password actuelle avec l'ancien password
				errPassword := bcrypt.CompareHashAndPassword([]byte(userDataAll.Password), []byte(UserInput.PasswordActuel))
				if errPassword != nil {
					message = errors.New("password incorrect!")
					fmt.Println(message)
				} else {

					hashPassword, _ := bcrypt.GenerateFromPassword([]byte(UserInput.Password), bcrypt.DefaultCost)
					UserInput.Password = string(hashPassword)
					// fmt.Println(UserInput.Password, "UserInput.Password")
					//Verifie le password entrée avec le confirm password
					errPassword := bcrypt.CompareHashAndPassword([]byte(UserInput.Password), []byte(UserInput.Cpassword))
					if errPassword != nil {
						message = errors.New("confirmation de password incorrect!")
						fmt.Println(message)
					} else {
						_, err := userModel.UpdatePassword(UserInput.Password, userDataAll.Id)
						if err != nil {
							fmt.Println(err)
						} else {
							message = errors.New("Password modifié avec succès!")
						}
						userCompetencesCles, _ := userModel.GetAllCompetenceClesOfUser(&competenceDataAll, uuid)
						nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
						nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
						nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
						barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere

						m := map[string]interface{}{
							"Profil":             res,
							"UserCompetenceCles": userCompetencesCles,
							"ProgressBarre":      barreProgressionTotal,
							"retour":             message,
						}
						temp, _ := template.ParseFiles("views/candidates-change-password.html")
						temp.Execute(w, m)
					}
				}
			}

		}
	}
}

/////////// *********************CETTE FONCTION PERMET D'AJOUTER UN FICHIER********************* ///////////////

const maxUploadSize = 2 * 1024 * 1024 // 2 mb

func uploadFile(fichier string, w http.ResponseWriter, r *http.Request) string {

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Println("Could not parse multipart form: n", err)
		renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
		// return
		return "ok"
	}

	// parse and validate file and post parameters
	file, fileHeader, err := r.FormFile(fichier)

	var newFileName string
	if err == nil {
		// renderError(w, "INVALID_FILE", http.StatusBadRequest)
		// return "ok"
		defer file.Close()
		// Get and print out file size
		fileSize := fileHeader.Size
		fmt.Println("File size (bytes): ", fileSize)
		// validate file size
		if fileSize > maxUploadSize {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return "ok"
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return "ok"
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		detectedFileType := http.DetectContentType(fileBytes)
		switch detectedFileType {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return "ok"
		}

		fileName := randToken(12)
		fileEndings, err := mime.ExtensionsByType(detectedFileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return "ok"
		}

		newFileName = fileName + fileEndings[0]
		var newPath string
		if fichier == "cv" {
			newPath = filepath.Join("public/cv", newFileName)
			fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)
		} else if fichier == "photo" {
			newPath = filepath.Join("public/profil", newFileName)
			fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

		} else if fichier == "cvCandidature" {
			newPath = filepath.Join("public/cvCandidatureG", newFileName)
			fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)
		}

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return "ok"
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return "ok"
		}
	}

	// w.Write([]byte(fmt.Sprintf("SUCCESS - use /files/%v to access the file", newFileName)))
	return newFileName
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

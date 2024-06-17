package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"digiJob/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// gestion d'erreur
// func erreur(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

type UserInput struct {
	Email          string
	Password       string
	PasswordActuel string
	Cpassword      string
}

var userModel = models.NewUserModel()

func Acceuil(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var barreProgInfoPerso entities.BarreProgressionInfoPerso
	var barreProgCv entities.BarreProgressionCv
	var barreProgCritere entities.BarreProgressionCritere
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/loginDigicodeuses", http.StatusSeeOther)

		} else {
			// data := map[string]interface{}{
			// 	"nom": session.Values["nom"],
			// }

			// temp, _ := template.ParseFiles("views/page-digi.html")

			// temp.Execute(w, data)
			id := session.Values["id"]
			//convertir un type interface en string
			uid := fmt.Sprintf("%v", id)
			uuid, _ := strconv.Atoi(uid)
			var offres entities.OffresGenerales
			var entreprise entities.Entreprises
			var critere entities.CritereDigicodeuses
			nbreProgressionInfoPerso, _ := userModel.GetBarreProgressionInfoPerso(&barreProgInfoPerso, uuid)
			nbreProgressionCv, _ := userModel.GetBarreProgressionCv(&barreProgCv, uuid)
			nbreProgressionCritere, _ := userModel.GetBarreProgressionCritere(&barreProgCritere, uuid)
			barreProgressionTotal := nbreProgressionInfoPerso + nbreProgressionCv + nbreProgressionCritere
			typeContrat, localite, _ := userModel.GetCriteresDigi(&critere, uuid)
			//Trie en fonction des critères
			res, _ := userModel.OfrresDigicodeuses(&offres, &entreprise, uuid, &critere, typeContrat, localite)
			data := map[string]interface{}{
				"Offres":        res,
				"ProgressBarre": barreProgressionTotal,
			}
			// fmt.Println(data, "tab")
			temp, _ := template.ParseFiles("views/page-digi.html")
			temp.Execute(w, data)
		}

	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/login.html")
		// erreur(err)
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// fmt.Println("Processus de connexion")
		r.ParseForm()
		UserInput := &UserInput{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		var user entities.UsersDigicodeuses
		userModel.Where(&user, "email", UserInput.Email)
		var message error
		if user.Email == "" {
			message = errors.New("Email incorrect!")
		} else {
			// fmt.Println(user.Email, "= user.Email")
			// fmt.Println(user.Password, "= user.Password")
			// fmt.Println(user.Pays)
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			if errPassword != nil {
				message = errors.New("password incorrect!")
			}

		}

		if message != nil {
			// renvoyer le message d'erreur
			data := map[string]interface{}{
				"error": message,
			}
			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		} else {

			session, _ := config.Store.Get(r, config.SESSION_ID)
			session.Values["loggedIn"] = true
			session.Values["id"] = user.Id
			session.Values["email"] = user.Email
			session.Values["username"] = user.Username
			session.Values["nom"] = user.Nom
			session.Values["prenoms"] = user.Prenoms
			session.Values["pays"] = user.Pays
			session.Save(r, w)
			// fmt.Println(session)
			http.Redirect(w, r, "/acceuil", http.StatusSeeOther)
			// if len(session.Values) < 10 {
			// 	temp, _ := template.ParseFiles("views/profileDigicodeuses.html")
			// 	temp.Execute(w, session.Values)
			// } else {
			// 	http.Redirect(w, r, "/acceuil", http.StatusSeeOther)
			// }

		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	// supprimier la session
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/inscription-digi.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		user := entities.UsersDigicodeuses{
			Nom:             r.Form.Get("nom"),
			Prenoms:         r.Form.Get("prenoms"),
			Email:           r.Form.Get("email"),
			NumeroTelephone: r.Form.Get("numero"),
			Indicatif:       r.Form.Get("indicatif"),
			Ville:           r.Form.Get("ville"),
			Pays:            r.Form.Get("pays"),
			// Username:            r.Form.Get("username"),
			Autorisation: r.Form.Get("autorisation"),
			Password:     r.Form.Get("password"),
			Cpassword:    r.Form.Get("cpassword"),
		}
		// fmt.Println(r.Form.Get("autorisation"))
		errorMessages := make(map[string]interface{})

		if user.Nom == "" {
			errorMessages["Nom"] = "Renseignez le nom"
		}

		if user.Prenoms == "" {
			errorMessages["Prenoms"] = "Renseignez le prenoms"
		}

		if user.Email == "" {
			errorMessages["Email"] = "Renseignez l'email"
		}

		// if user.Username == "" {
		// 	errorMessages["Username"] = "Renseignez le username"
		// }

		if user.Password == "" {
			errorMessages["Password"] = "Renseignez le password"
		}

		if user.Cpassword == "" {
			errorMessages["Cpassword"] = "Renseignez la confirmation password"
		} else {
			if user.Password != user.Cpassword {
				errorMessages["Cpassword"] = "la cofirmation password ne semble pas correspondre"
			}
		}
		if len(errorMessages) > 0 {
			// on va afficher cela sur le formulaire
			data := map[string]interface{}{
				"validation": errorMessages,
			}
			temp, _ := template.ParseFiles("views/inscription-digi.html")
			temp.Execute(w, data)
		} else {
			// processus d'insertion dans la base de données
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			//inserer les données dans la base de données
			_, err := userModel.Create(user)
			var message string
			if err != nil {
				message = "Echec d'enregitrement:" + message
			} else {
				message = "Processus de connexion réussi"
			}

			data := map[string]interface{}{
				"retour": message,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		}
	}
}

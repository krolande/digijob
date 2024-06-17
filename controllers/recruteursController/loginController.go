package recruteurscontroller

import (
	"digiJob/config"
	recruteursentities "digiJob/entities/recruteursEntities"
	recruteursmodel "digiJob/models/recruteursModel"
	"errors"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

type EntrepriseInput struct {
	Nom      string
	Email    string
	Password string
}

var userModel = recruteursmodel.NewUserModel()

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/connexion-recruteur.html")
		// erreur(err)
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// fmt.Println("Processus de connexion")
		r.ParseForm()
		EntrepriseInput := &EntrepriseInput{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		var entreprise recruteursentities.Entreprise
		userModel.WhereEntreprise(&entreprise, "email", EntrepriseInput.Email)
		var message error
		if entreprise.Email == "" {
			message = errors.New("Email incorrect!")
		} else {
			// fmt.Println(user.Email, "= user.Email")
			// fmt.Println(user.Password, "= user.Password")
			// fmt.Println(user.Pays)
			errPassword := bcrypt.CompareHashAndPassword([]byte(entreprise.Password), []byte(EntrepriseInput.Password))
			if errPassword != nil {
				message = errors.New("password incorrect!")
			}

		}

		if message != nil {
			// renvoyer le message d'erreur
			data := map[string]interface{}{
				"error": message,
			}
			temp, _ := template.ParseFiles("views/connexion-recruteur.html")
			temp.Execute(w, data)
		} else {

			session, _ := config.Store.Get(r, config.SESSION_ID)
			session.Values["loggedIn"] = true
			session.Values["id"] = entreprise.Id
			session.Values["email"] = entreprise.Email
			session.Save(r, w)
			// fmt.Println(session)
			http.Redirect(w, r, "/dashboardRecruteur", http.StatusSeeOther)

		}
	}
}

// deconnexion recruteur

func LogoutRecruteurs(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	// supprimier la session
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

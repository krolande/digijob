package controllers

import (
	"digiJob/config"
	"digiJob/entities"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func ApercuCv(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var userDataAll entities.UsersDigi
	var competenceDataAll entities.CompetenceCles
	var langues entities.Langues
	var formation entities.Formation
	var experience entities.ExperienceProfessionnelle
	var prixCertifications entities.PrixCertification
	var competenceProfessionnelle entities.CompetenceProfessionnelle
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
				langueUser, _, _ := userModel.GetLanguesUser(&langues, uuid)
				resFormation, _, _ := userModel.ReadForamtion(&formation, uuid)
				resExperience, _, _ := userModel.ReadExperience(&experience, uuid)
				resCompetencePro, _, _ := userModel.GetCompetenceProfessionnelle(&competenceProfessionnelle, uuid)
				resPrixCertification, _, _ := userModel.ReadPrixCertification(&prixCertifications, uuid)

				data := map[string]interface{}{
					"Profil":                    res,
					"UserCompetenceCles":        userCompetencesCles,
					"Langues":                   langueAll,
					"LanguesUser":               langueUser,
					"Formations":                resFormation,
					"ExperienceProfessionnelle": resExperience,
					"CompetenceProfessionnelle": resCompetencePro,
					"PrixCertification":         resPrixCertification,
				}
				// GeneratePDF(data)
				temp, _ := template.ParseFiles("views/apercu-mon-cv.html")
				temp.Execute(w, data)

			}
		}
	}
}

// func GeneratePDF(data map[string]interface{}) ([]byte, error) {

// 	fmt.Println("ok")
// 	var templ *template.Template
// 	var err error

// 	// use Go's default HTML template generation tools to generate your HTML
// 	if templ, err = template.ParseFiles("apercu-mon-cv.html"); err != nil {
// 		return nil, err
// 	}

// 	// apply the parsed HTML template data and keep the result in a Buffer
// 	var body bytes.Buffer
// 	if err = templ.Execute(&body, data); err != nil {
// 		return nil, err
// 	}

// 	// initalize a wkhtmltopdf generator
// 	pdfg, err := wkhtmltopdf.NewPDFGenerator()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// read the HTML page as a PDF page
// 	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

// 	// enable this if the HTML file contains local references such as images, CSS, etc.
// 	page.EnableLocalFileAccess.Set(true)

// 	// add the page to your generator
// 	pdfg.AddPage(page)

// 	// manipulate page attributes as needed
// 	pdfg.MarginLeft.Set(0)
// 	pdfg.MarginRight.Set(0)
// 	pdfg.Dpi.Set(300)
// 	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
// 	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

// 	// magic
// 	err = pdfg.Create()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return pdfg.Bytes(), nil
// }

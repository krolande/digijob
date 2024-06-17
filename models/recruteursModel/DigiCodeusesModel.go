package recruteursmodel

import (
	"digiJob/entities"
	recruteursentities "digiJob/entities/recruteursEntities"
	"fmt"
)

func (u UserModel) AllDigiCodeuses(user *recruteursentities.UsersDigi) (interface{}, int, error) {
	res := recruteursentities.UsersDigicodeuses{}
	allResultat := []recruteursentities.UsersDigicodeuses{}
	// result, err := u.db.Query("SELECT id, nom, prenoms, email, password, pays, ville, date_naissance, indicatif, numero_telephone, nationalite, cv, photo, titrePosteCv, portfolio, facebook, twitter, linkedin, temoignage, autorisation, description, niveau_etude, adresse, created_at, updated_at FROM usersDigicodeuses WHERE id=?", id)
	result, err := u.db.Query("SELECT * FROM `usersDigicodeuses`")

	if err != nil {
		return res, 0, err
	}
	competence := entities.CompetenceCles{}
	cmpt := 0
	for result.Next() {
		err = result.Scan(&user.Id, &user.Nom, &user.Prenoms, &user.Email, &user.Password, &user.Username, &user.Pays, &user.Ville, &user.DateNaissance, &user.Indicatif, &user.NumeroTelephone, &user.Nationalite, &user.Cv, &user.Photo, &user.TitrePosteCv, &user.Porfolio, &user.Facebook, &user.Twitter, &user.Linkedin, &user.Temoignage, &user.Autorisation, &user.Description, &user.NiveauEtude, &user.AnneeExperience, &user.Adresse, &user.Statut, &user.CreatedAt, &user.UpdatedAt)
		// err = result.Scan(&user.Id, &user.Nom, &user.Prenoms, &user.Email, &user.Password)
		if err != nil {
			fmt.Println(err)
			return res, 0, err
		}
		// fmt.Println(user)
		res.Id = user.Id
		res.Nom = user.Nom
		res.Prenoms = user.Prenoms
		res.Email = user.Email
		res.Pays = user.Pays
		res.Ville = user.Ville
		res.DateNaissance = user.DateNaissance.String
		res.Indicatif = user.Indicatif.String
		res.NumeroTelephone = user.NumeroTelephone.String
		res.Nationalite = user.Nationalite.String
		res.Cv = user.Cv.String
		res.Photo = user.Photo.String
		res.TitrePosteCv = user.TitrePosteCv.String
		res.Porfolio = user.Porfolio.String
		res.Facebook = user.Facebook.String
		res.Twitter = user.Twitter.String
		res.Linkedin = user.Linkedin.String
		res.Temoignage = user.Temoignage.String
		res.Autorisation = user.Autorisation.String
		res.Description = user.Description.String
		res.NiveauEtude = user.NiveauEtude.String
		res.Adresse = user.Adresse.String
		res.Statut = user.Statut.String
		res.Password = user.Password
		res.AnneeExperience = user.AnneeExperience.String
		res.Username = user.Username.String
		comp, errC := u.db.Query("SELECT cc.* FROM competenceCles cc JOIN competenceCle_usersDigicodeuses co ON cc.id= co.	competenceCles_id WHERE co.usersDigicodeuses_id = ?", user.Id)
		if errC != nil {
			return comp, 0, errC
		}
		for comp.Next() {
			comp.Scan(&user.IdC, &user.Libelle, &user.CreatedAtC, &user.UpdatedAtC)
			competence.Id = user.IdC
			competence.Libelle = user.Libelle
			res.TabLibelle = append(res.TabLibelle, competence)
		}
		allResultat = append(allResultat, res)
		res.TabLibelle = []entities.CompetenceCles{}
		cmpt++
	}
	return allResultat, cmpt, err
}

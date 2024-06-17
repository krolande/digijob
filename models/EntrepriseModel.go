package models

import (
	"digiJob/entities"
)

// recupere toutes les entreprises de la BD
func (u UserModel) GetAllEntreprises(entreprise *entities.Entreprises) (interface{}, int, error) {
	row, err := u.db.Query("SELECT * FROM entreprises ORDER BY id DESC")
	if err != nil {
		return row, 0, err
	}
	defer row.Close()
	allEntreprises := entities.Entreprise{}
	res := []entities.Entreprise{}

	for row.Next() {
		row.Scan(&entreprise.Id, &entreprise.NomEntreprise, &entreprise.SecteurActivite, &entreprise.Adresse, &entreprise.Logo, &entreprise.SiteWeb, &entreprise.DescriptionEntreprise, &entreprise.NumeroEntreprise, &entreprise.VideoEntreprise, &entreprise.TailleEntreprise, &entreprise.Facebook, &entreprise.Twitter, &entreprise.Linkedin, &entreprise.Email, &entreprise.Slogan, &entreprise.Password, &entreprise.CreatedAt, &entreprise.UpdateAt)
		allEntreprises.Id = entreprise.Id
		allEntreprises.NomEntreprise = entreprise.NomEntreprise
		allEntreprises.Logo = entreprise.Logo
		allEntreprises.DescriptionEntreprise = entreprise.DescriptionEntreprise.String
		res = append(res, allEntreprises)
	}
	return res, len(res), nil
}

//Recupere les information d'une entreprises

func (u UserModel) Entreprise(entreprise *entities.Entreprises, id int) (interface{}, error) {
	row, err := u.db.Query("SELECT * FROM entreprises WHERE id = ?", id)
	if err != nil {
		return row, err
	}
	defer row.Close()
	allEntreprises := entities.Entreprise{}
	allcontenue := entities.ContenueEntreprises{}
	for row.Next() {
		row.Scan(&entreprise.Id, &entreprise.NomEntreprise, &entreprise.SecteurActivite, &entreprise.Adresse, &entreprise.Logo, &entreprise.SiteWeb, &entreprise.DescriptionEntreprise, &entreprise.NumeroEntreprise, &entreprise.VideoEntreprise, &entreprise.TailleEntreprise, &entreprise.Facebook, &entreprise.Twitter, &entreprise.Linkedin, &entreprise.Email, &entreprise.Slogan, &entreprise.Password, &entreprise.CreatedAt, &entreprise.UpdateAt)
		allEntreprises.Id = entreprise.Id
		allEntreprises.NomEntreprise = entreprise.NomEntreprise
		allEntreprises.Logo = entreprise.Logo
		allEntreprises.DescriptionEntreprise = entreprise.DescriptionEntreprise.String
		allEntreprises.Slogan = entreprise.Slogan.String
		allEntreprises.TailleEntreprise = entreprise.TailleEntreprise.String
		allEntreprises.SecteurActivite = entreprise.SecteurActivite
		allEntreprises.Adresse = entreprise.Adresse
		allEntreprises.SiteWeb = entreprise.SiteWeb.String
		allEntreprises.Email = entreprise.Email.String
		allEntreprises.NumeroEntreprise = entreprise.NumeroEntreprise.String
		allEntreprises.Facebook = entreprise.Facebook.String
		allEntreprises.Twitter = entreprise.Twitter.String
		allEntreprises.VideoEntreprise = entreprise.VideoEntreprise.String
		//renvoyer les contenues de l'entreprise
		cont, erreur := u.db.Query("SELECT * FROM `photoEntreprises` WHERE entreprises_id=?", allEntreprises.Id)
		if erreur != nil {
			return cont, erreur
		}
		for cont.Next() {
			cont.Scan(&entreprise.IdC, &entreprise.NomPhoto, &entreprise.PhotoUpload, &entreprise.Entreprises_id, &entreprise.Created_At, &entreprise.Update_At)
			allcontenue.Id = entreprise.IdC
			allcontenue.PhotoUpload = entreprise.PhotoUpload
			allcontenue.NomPhoto = entreprise.NomPhoto.String
			allEntreprises.ContenueEntreprise = append(allEntreprises.ContenueEntreprise, allcontenue)
		}
	}
	return allEntreprises, nil
}

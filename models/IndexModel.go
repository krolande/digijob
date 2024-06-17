package models

import (
	"digiJob/entities"
)

//recuperer les 4 derniers offres
func (u UserModel) OfrresGeneralesIndex(offres *entities.OffresGenerales, entreprise *entities.Entreprises) (interface{}, error) {

	offresAll := entities.OffresEntreprises{}

	res := []entities.OffresEntreprises{}
	row, err := u.db.Query("select p.*, lu.* from offreGenerales p JOIN entreprises lu ON lu.id = p.entreprises_id ORDER BY p.id DESC LIMIT 4")
	if err != nil {
		return res, err
	}

	defer row.Close()

	//recupère les données et les saisir dans les entités utilisaturs
	for row.Next() {
		row.Scan(&offres.Id, &offres.TitrePoste, &offres.DescriptionPoste, &offres.NombrePoste, &offres.AnneeExperience, &offres.Formation, &offres.Ville, &offres.DateLimite, &offres.Entreprise_id, &offres.CreatedAt, &offres.UpdateAt, &offres.TypeContrat, &entreprise.Id, &entreprise.NomEntreprise, &entreprise.SecteurActivite, &entreprise.Adresse, &entreprise.Logo, &entreprise.SiteWeb, &entreprise.DescriptionEntreprise, &entreprise.NumeroEntreprise, &entreprise.VideoEntreprise, &entreprise.TailleEntreprise, &entreprise.Facebook, &entreprise.Twitter, &entreprise.Linkedin, &entreprise.Email, &entreprise.Slogan, &entreprise.Password, &entreprise.CreatedAt, &entreprise.UpdateAt)
		offresAll.Id = offres.Id
		offresAll.TitrePoste = offres.TitrePoste
		offresAll.DescriptionPoste = offres.DescriptionPoste
		offresAll.NombrePoste = offres.NombrePoste
		offresAll.AnneeExperience = offres.AnneeExperience
		offresAll.TypeContrat = offres.TypeContrat
		offresAll.Ville = offres.Ville
		offresAll.DateLimite = offres.DateLimite
		offresAll.CreatedAt = offres.CreatedAt
		offresAll.LogoEntreprise = entreprise.Logo
		offresAll.NomEntreprise = entreprise.NomEntreprise
		res = append(res, offresAll)
		offres.Id = 0
		offres.TitrePoste = ""
		offres.DescriptionPoste = ""
		offres.NombrePoste = 0
		offres.AnneeExperience = ""
		offres.TypeContrat = ""
		offres.Ville = ""
		offresAll.DateLimite = ""
	}
	// fmt.Println(res)
	return res, nil
}

// recupere 5 entreprises de la BD
func (u UserModel) GetEntreprises(entreprise *entities.Entreprises) (interface{}, int, error) {
	row, err := u.db.Query("SELECT * FROM entreprises ORDER BY id DESC LIMIT 5")
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

//recuperer 3 évènement

func (u UserModel) GetEvenements(evenement *entities.EvenementEntrepriseBd) (interface{}, error) {
	allEvenements := entities.EvenementEntreprises{}
	res := []entities.EvenementEntreprises{}
	row, err := u.db.Query("SELECT ev.*, en.* FROM evenements ev JOIN entreprises en ON en.id = ev.entreprises_id ORDER BY ev.id DESC LIMIT 3")
	if err != nil {
		return res, err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&evenement.Id, &evenement.TitreEvent, &evenement.DescriptionEvent, &evenement.FileUpload, &evenement.DateEvent, &evenement.TypeEvent, &evenement.Heure, &evenement.FraisInscriptionEvent, &evenement.ModaliteEvent, &evenement.Lien, &evenement.LieuEvent, &evenement.DateLimite, &evenement.Entreprises_id, &evenement.CreatedAt, &evenement.UpdateAt, &evenement.IdE, &evenement.NomEntreprise, &evenement.SecteurActivite, &evenement.Adresse, &evenement.Logo, &evenement.SiteWeb, &evenement.DescriptionEntreprise, &evenement.NumeroEntreprise, &evenement.VideoEntreprise, &evenement.TailleEntreprise, &evenement.Facebook, &evenement.Twitter, &evenement.Linkedin, &evenement.Email, &evenement.Slogan, &evenement.Password, &evenement.CreatedAtE, &evenement.UpdateAtE)
		allEvenements.Id = evenement.Id
		allEvenements.TitreEvent = evenement.TitreEvent
		allEvenements.DescriptionEvent = evenement.DescriptionEvent
		allEvenements.FileUpload = evenement.FileUpload
		allEvenements.DateEvent = evenement.DateEvent
		allEvenements.DateLimites = evenement.DateLimite.String
		allEvenements.FraisInscriptionEvent = evenement.FraisInscriptionEvent.String
		allEvenements.LieuEvent = evenement.LieuEvent

		//Entreprises
		allEvenements.IdE = evenement.IdE
		allEvenements.NomEntreprise = evenement.NomEntreprise
		allEvenements.Logo = evenement.Logo
		res = append(res, allEvenements)
	}

	return res, nil
}

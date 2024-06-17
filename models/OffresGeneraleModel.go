package models

import (
	"digiJob/entities"
)

type OffresGenerales struct {
	Id                 int
	TitrePoste         string
	DescriptionPoste   string
	NombrePoste        int
	AnneeExperience    string
	Formation          string
	Ville              string
	DateLimite         string
	UsersEntreprise_id int
	TypeContrat        string
	CreatedAt          string
}

//All offre
func (u UserModel) AllOfrresGenerales(offres *entities.OffresGenerales, entreprise *entities.Entreprises) (interface{}, error) {

	offresAll := entities.OffresEntreprises{}

	res := []entities.OffresEntreprises{}
	row, err := u.db.Query("select p.*, lu.* from offreGenerales p JOIN entreprises lu ON lu.id = p.entreprises_id ORDER BY p.id DESC")
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

// recuperer les criteres et les fait passer dans cette fonction
func (u UserModel) GetCriteresDigi(critere *entities.CritereDigicodeuses, uuid int) (string, string, error) {
	row, err := u.db.Query("SELECT * FROM critereDigicodeuses WHERE usersDigicodeuses_id = ?", uuid)
	if err != nil {
		return "res1", "res2", err
	}
	defer row.Close()

	critereDigi := entities.CritereDigicodeuses{}

	for row.Next() {
		row.Scan(&critere.Id, &critere.TypeContrat, &critere.Localite, &critere.DureeContrat, &critere.Modalite, &critere.MetierRecherche, &critere.UsersDigicodeuse_id, &critere.CreatedAt, &critere.UpdatedAt)
		critereDigi.TypeContrat = critere.TypeContrat
		critereDigi.Localite = critere.Localite
		critereDigi.MetierRecherche = critere.MetierRecherche

	}

	return critereDigi.TypeContrat, critereDigi.Localite, nil
}

////////////////////////************************/////////////////////////
// lien de ressource: https://www.sqlshack.com/sql-order-by-clause-overview-and-examples/
//afficher les offres en fonction des criteres
func (u UserModel) OfrresDigicodeuses(offres *entities.OffresGenerales, entreprise *entities.Entreprises, uuid int, critere *entities.CritereDigicodeuses, typeContrat string, localite string) (interface{}, error) {
	offresAll := entities.OffresEntreprises{}

	res := []entities.OffresEntreprises{}
	row, err := u.db.Query("select p.*, lu.*, cr.* from offreGenerales p JOIN entreprises lu ON lu.id = p.entreprises_id JOIN critereDigicodeuses cr ON cr.usersDigicodeuses_id = ? WHERE p.typeContrat = ?", uuid, typeContrat)
	if err != nil {
		return res, err
	}

	defer row.Close()

	//recupère les données et les saisir dans les entités utilisaturs
	for row.Next() {
		row.Scan(&offres.Id, &offres.TitrePoste, &offres.DescriptionPoste, &offres.NombrePoste, &offres.AnneeExperience, &offres.Formation, &offres.Ville, &offres.DateLimite, &offres.Entreprise_id, &offres.CreatedAt, &offres.UpdateAt, &offres.TypeContrat, &entreprise.Id, &entreprise.NomEntreprise, &entreprise.SecteurActivite, &entreprise.Adresse, &entreprise.Logo, &entreprise.SiteWeb, &entreprise.DescriptionEntreprise, &entreprise.NumeroEntreprise, &entreprise.VideoEntreprise, &entreprise.TailleEntreprise, &entreprise.Facebook, &entreprise.Twitter, &entreprise.Linkedin, &entreprise.Email, &entreprise.Slogan, &entreprise.Password, &entreprise.CreatedAt, &entreprise.UpdateAt, &critere.Id, &critere.TypeContrat, &critere.Localite, &critere.DureeContrat, &critere.Modalite, &critere.MetierRecherche, &critere.UsersDigicodeuse_id, &critere.CreatedAt, &critere.UpdatedAt)
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
		offresAll.TypeContratC = critere.TypeContrat
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

// Renvoie les détails d'une offre
func (u UserModel) DetailOffresPoster(detailoffre *entities.OffresEntreprises, id int) (interface{}, error) {
	row, err := u.db.Query("SELECT o.*, e.* FROM offreGenerales o JOIN entreprises e ON e.id = o.entreprises_id WHERE o.id =? ", id)
	offreDetail := entities.OffresEntreprisesRec{}
	if err != nil {
		return offreDetail, err
	}
	defer row.Close()

	//Les missions du poste
	missionOG := entities.MissionOffreGenerale{}

	//Profil recherché
	profilOG := entities.ProfilRechercheOffreG{}

	//CompetenceClées du poste
	competence := entities.CompetenceCles{}

	for row.Next() {
		row.Scan(&detailoffre.Id, &detailoffre.TitrePoste, &detailoffre.DescriptionPoste, &detailoffre.NombrePoste, &detailoffre.AnneeExperience, &detailoffre.Formation, &detailoffre.Ville, &detailoffre.DateLimite, &detailoffre.Entreprise_id, &detailoffre.CreatedAt, &detailoffre.UpdateAt, &detailoffre.TypeContrat, &detailoffre.IdE, &detailoffre.NomEntreprise, &detailoffre.SecteurActivite, &detailoffre.Adresse, &detailoffre.LogoEntreprise, &detailoffre.SiteWeb, &detailoffre.DescriptionEntreprise, &detailoffre.NumeroEntreprise, &detailoffre.VideoEntreprise, &detailoffre.TailleEntreprise, &detailoffre.Facebook, &detailoffre.Twitter, &detailoffre.Linkedin, &detailoffre.Email, &detailoffre.Slogan, &detailoffre.Password, &detailoffre.CreatedAtE, &detailoffre.UpdateAtE)
		offreDetail.Id = detailoffre.Id
		offreDetail.TitrePoste = detailoffre.TitrePoste
		offreDetail.DescriptionPoste = detailoffre.DescriptionPoste
		offreDetail.NombrePoste = detailoffre.NombrePoste
		offreDetail.AnneeExperience = detailoffre.AnneeExperience
		offreDetail.TypeContrat = detailoffre.TypeContrat
		offreDetail.Ville = detailoffre.Ville
		offreDetail.DateLimite = detailoffre.DateLimite
		offreDetail.CreatedAt = detailoffre.CreatedAt
		//////////////     *********************Entreprise***********************//
		offreDetail.IdE = detailoffre.IdE
		offreDetail.LogoEntreprise = detailoffre.LogoEntreprise
		offreDetail.TypeContratC = detailoffre.TypeContrat
		offreDetail.NomEntreprise = detailoffre.NomEntreprise
		offreDetail.SecteurActivite = detailoffre.SecteurActivite
		offreDetail.Adresse = detailoffre.Adresse
		offreDetail.SiteWeb = detailoffre.SiteWeb.String
		offreDetail.DescriptionEntreprise = detailoffre.DescriptionEntreprise.String
		offreDetail.NumeroEntreprise = detailoffre.NumeroEntreprise.String
		offreDetail.Facebook = detailoffre.Facebook.String
		offreDetail.Twitter = detailoffre.Twitter.String
		offreDetail.Linkedin = detailoffre.Linkedin.String
		offreDetail.Email = detailoffre.Email.String
		offreDetail.CreatedAtE = detailoffre.CreatedAtE
		///////// ****************************** MISSION DE L'OFFRE *******************************
		req, erreur := u.db.Query("SELECT * FROM `missionOffreGenerales` WHERE offreGenerales_id = ?", offreDetail.Id)
		if erreur != nil {
			return req, erreur
		}
		for req.Next() {
			req.Scan(&detailoffre.IdM, &detailoffre.MissionPoste, &detailoffre.OffresGenerales_id, &detailoffre.CreatedAtM, &detailoffre.UpdateAtM)
			missionOG.IdM = detailoffre.IdM
			missionOG.MissionPoste = detailoffre.MissionPoste
			offreDetail.TabMissionPoste = append(offreDetail.TabMissionPoste, missionOG)
		}

		//Les profil recherchés du poste
		profil, errP := u.db.Query("SELECT * FROM `profilRechercherOffreGenerales` WHERE offreGenerales_id = ?", offreDetail.Id)
		if errP != nil {
			return profil, errP
		}
		for profil.Next() {
			profil.Scan(&detailoffre.IdP, &detailoffre.ProfilPoste, &detailoffre.OffresGenerales_idP, &detailoffre.CreatedAtP, &detailoffre.UpdateAtP)
			profilOG.IdP = detailoffre.IdP
			profilOG.ProfilPoste = detailoffre.ProfilPoste
			offreDetail.TabProfilRecherche = append(offreDetail.TabProfilRecherche, profilOG)
		}

		//Competences clées du poste
		comp, errC := u.db.Query("SELECT cc.* FROM competenceCles cc JOIN competenceCle_offreGenerales co ON cc.id= co.	competenceCles_id WHERE co.offreGenerales_id = ?", offreDetail.Id)
		if errC != nil {
			return comp, errC
		}
		for comp.Next() {
			comp.Scan(&offreDetail.IdCc, &offreDetail.Libelle, &offreDetail.CreatedAtCc, &offreDetail.UpdatedAtCc)
			competence.Id = offreDetail.IdCc
			competence.Libelle = offreDetail.Libelle
			offreDetail.TabLibelle = append(offreDetail.TabLibelle, competence)
		}
	}
	// fmt.Println(offreDetail.TabLibelle)
	return offreDetail, nil
}

// Renvoie les offres d'une entreprise
func (u UserModel) AllOfrresGeneralEntreprise(offres *entities.OffresGenerales, entreprise *entities.Entreprises, id int) (interface{}, error) {

	offresAll := entities.OffresEntreprises{}

	res := []entities.OffresEntreprises{}
	row, err := u.db.Query("select p.*, lu.* from offreGenerales p JOIN entreprises lu ON lu.id = p.entreprises_id WHERE lu.id = ? ORDER BY p.id DESC", id)
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
	return res, nil
}

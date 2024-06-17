package models

import (
	"digiJob/entities"
	"fmt"
)

// Inserer les  candidatures offres Generales

func (u UserModel) InsertCandidatureOffreG(candidature entities.CandidatureOffreGenerales) (int64, error) {
	insForm, err := u.db.Exec("INSERT INTO `candidatureOffreGenerales`(`motivation`, `cv`, `offreGenerales_id`, `usersDigicodeuses_id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE motivation = ?, cv = ?",
		candidature.Motivation, candidature.Cv, candidature.OffresGenerales_id, candidature.UsersDigicodeuses_id, candidature.Motivation, candidature.Cv)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insForm.LastInsertId()

	return lastInsertId, nil
}

//inserer les statuts de la candidature

func (u UserModel) InsStatutCandidatureOG(statut entities.StatutCandidature) (int64, error) {
	insForm, err := u.db.Exec("INSERT INTO `statutCandidatures`(`satutAuto`, `satutTalentManager`, `satutEntreprise`, `candidatureOffreGenerales_id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE satutTalentManager = ?, satutEntreprise = ?",
		statut.SatutAuto, statut.SatutTalentManager, statut.SatutEntreprise, statut.CandidatureOffreGenerales_id, statut.SatutTalentManager, statut.SatutEntreprise)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insForm.LastInsertId()

	return lastInsertId, nil
}

// recuprere l' Id de la candidature

func (u UserModel) GetIdCandidature(candidature *entities.CandidatureOffreGenerales, offres_id int, uuid int) (int, error) {
	row, err := u.db.Query("SELECT * FROM candidatureOffreGenerales WHERE usersDigicodeuses_id = ? AND offreGenerales_id = ?", uuid, offres_id)
	if err != nil {
		return 0, err
	}
	defer row.Close()
	getCand := entities.CandidatureOffreGenerales{}
	for row.Next() {
		row.Scan(&candidature.Id, &candidature.Motivation, &candidature.Cv, &candidature.OffresGenerales_id, &candidature.UsersDigicodeuses_id, &candidature.Created_at, &candidature.Updated_at)
		getCand.Id = candidature.Id
	}
	return getCand.Id, nil
}

// Renvoyer toutes les candidatures

func (u UserModel) GetAllCandidatureDigi(candidatureTable *entities.AllCandidaturesDigi, uuid int) (interface{}, int, error) {
	row, err := u.db.Query("SELECT ca.*, of.*, en.*, st.* FROM candidatureOffreGenerales ca JOIN offreGenerales of ON of.id= ca.offreGenerales_id JOIN entreprises en ON en.id= of.entreprises_id JOIN statutCandidatures st ON st.candidatureOffreGenerales_id = ca.id WHERE ca.usersDigicodeuses_id = ? ORDER BY ca.id DESC", uuid)
	if err != nil {
		return row, 0, err
	}

	defer row.Close()
	res := []entities.AllCandidaturesDigi{}
	allCandidatures := entities.AllCandidaturesDigi{}
	// les variables que j'ai crées pour recuperer les competences clées en fonction de l'offres_id
	// resultCompClees: tableau qui regroupe l'ensemble des competences clées
	// j'utilise allCompetenceCles pour recuperer chaque competence clées et les enregistrer parceque
	// j'ai besoin d'une variable qui a les caracteristiques de la structure "CompetenceCles" pour recuperer les données de mêmes types
	resultCompClees := []entities.CompetenceCles{}
	allCompetenceCles := entities.CompetenceCles{}
	for row.Next() {
		row.Scan(&candidatureTable.IdC, &candidatureTable.Motivation, &candidatureTable.Cv, &candidatureTable.OffresGenerales_id, &candidatureTable.UsersDigicodeuses_id, &candidatureTable.Created_atC, &candidatureTable.Updated_atC, &candidatureTable.Id, &candidatureTable.TitrePoste, &candidatureTable.DescriptionPoste, &candidatureTable.NombrePoste, &candidatureTable.AnneeExperience, &candidatureTable.Formation, &candidatureTable.Ville, &candidatureTable.DateLimite, &candidatureTable.Entreprise_id, &candidatureTable.CreatedAt, &candidatureTable.UpdateAt, &candidatureTable.TypeContrat, &candidatureTable.IdE, &candidatureTable.NomEntreprise, &candidatureTable.SecteurActivite, &candidatureTable.Adresse, &candidatureTable.LogoEntreprise, &candidatureTable.SiteWeb, &candidatureTable.DescriptionEntreprise, &candidatureTable.NumeroEntreprise, &candidatureTable.VideoEntreprise, &candidatureTable.TailleEntreprise, &candidatureTable.Facebook, &candidatureTable.Twitter, &candidatureTable.Linkedin, &candidatureTable.Email, &candidatureTable.Slogan, &candidatureTable.Password, &candidatureTable.CreatedAtE, &candidatureTable.UpdateAtE, &candidatureTable.IdS, &candidatureTable.SatutAuto, &candidatureTable.SatutTalentManager, &candidatureTable.SatutEntreprise, &candidatureTable.CandidatureOffreGenerales_id, &candidatureTable.Created_atS, &candidatureTable.Updated_atS)
		//offres
		allCandidatures.Id = candidatureTable.Id
		allCandidatures.TitrePoste = candidatureTable.TitrePoste
		allCandidatures.TypeContrat = candidatureTable.TypeContrat
		allCandidatures.AnneeExperience = candidatureTable.AnneeExperience
		allCandidatures.DescriptionPoste = candidatureTable.DescriptionPoste
		allCandidatures.OffresGenerales_id = candidatureTable.OffresGenerales_id
		allCandidatures.NomEntreprise = candidatureTable.NomEntreprise
		allCandidatures.LogoEntreprise = candidatureTable.LogoEntreprise
		allCandidatures.Ville = candidatureTable.Ville
		allCandidatures.IdC = candidatureTable.IdC
		allCandidatures.Created_atC = candidatureTable.Created_atC
		allCandidatures.SatutEntreprise = candidatureTable.SatutEntreprise
		allCandidatures.Motivation = candidatureTable.Motivation
		allCandidatures.IdS = candidatureTable.IdS

		//**************************************************************
		req, erreur := u.db.Query("SELECT cc.* FROM competenceCles cc JOIN competenceCle_offreGenerales co ON co.competenceCles_id = cc.id WHERE co.offreGenerales_id  = ?", allCandidatures.OffresGenerales_id)
		if err != nil {
			return req, 0, erreur
		}
		for req.Next() {
			req.Scan(&candidatureTable.IdCc, &candidatureTable.Libelle, &candidatureTable.Created_atCc, &candidatureTable.Updated_atCc)
			allCompetenceCles.Id = candidatureTable.IdCc
			allCompetenceCles.Libelle = candidatureTable.Libelle
			resultCompClees = append(resultCompClees, allCompetenceCles)

		}
		allCandidatures.LibelleTab = append(allCandidatures.LibelleTab, resultCompClees)
		res = append(res, allCandidatures)
		//Initialisation des 2 variables a 0 pour pouvoir recuperer de nouvelle données et les enregistrer
		resultCompClees = []entities.CompetenceCles{}
		allCandidatures.LibelleTab = [][]entities.CompetenceCles{}

	}
	nbreCandidature := len(res)
	return res, nbreCandidature, nil
}

// Candidaturess préselectionnées

func (u UserModel) AllCandidatureDigiPresel(candidatureTable *entities.AllCandidaturesDigi, uuid int) (interface{}, int, error) {
	row, err := u.db.Query("SELECT ca.*, of.*, en.*, st.* FROM candidatureOffreGenerales ca JOIN offreGenerales of ON of.id= ca.offreGenerales_id JOIN entreprises en ON en.id= of.entreprises_id JOIN statutCandidatures st ON st.candidatureOffreGenerales_id = ca.id WHERE ca.usersDigicodeuses_id = ? AND st.satutEntreprise = ? ORDER BY ca.id DESC", uuid, "Approuvée")
	if err != nil {
		return row, 0, err
	}

	defer row.Close()
	res := []entities.AllCandidaturesDigi{}
	allCandidatures := entities.AllCandidaturesDigi{}

	for row.Next() {
		row.Scan(&candidatureTable.IdC, &candidatureTable.Motivation, &candidatureTable.Cv, &candidatureTable.OffresGenerales_id, &candidatureTable.UsersDigicodeuses_id, &candidatureTable.Created_atC, &candidatureTable.Updated_atC, &candidatureTable.Id, &candidatureTable.TitrePoste, &candidatureTable.DescriptionPoste, &candidatureTable.NombrePoste, &candidatureTable.AnneeExperience, &candidatureTable.Formation, &candidatureTable.Ville, &candidatureTable.DateLimite, &candidatureTable.Entreprise_id, &candidatureTable.CreatedAt, &candidatureTable.UpdateAt, &candidatureTable.TypeContrat, &candidatureTable.IdE, &candidatureTable.NomEntreprise, &candidatureTable.SecteurActivite, &candidatureTable.Adresse, &candidatureTable.LogoEntreprise, &candidatureTable.SiteWeb, &candidatureTable.DescriptionEntreprise, &candidatureTable.NumeroEntreprise, &candidatureTable.VideoEntreprise, &candidatureTable.TailleEntreprise, &candidatureTable.Facebook, &candidatureTable.Twitter, &candidatureTable.Linkedin, &candidatureTable.Email, &candidatureTable.Slogan, &candidatureTable.Password, &candidatureTable.CreatedAtE, &candidatureTable.UpdateAtE, &candidatureTable.IdS, &candidatureTable.SatutAuto, &candidatureTable.SatutTalentManager, &candidatureTable.SatutEntreprise, &candidatureTable.CandidatureOffreGenerales_id, &candidatureTable.Created_atS, &candidatureTable.Updated_atS)
		//offres
		allCandidatures.Id = candidatureTable.Id
		allCandidatures.TitrePoste = candidatureTable.TitrePoste
		allCandidatures.TypeContrat = candidatureTable.TypeContrat
		allCandidatures.NomEntreprise = candidatureTable.NomEntreprise
		allCandidatures.LogoEntreprise = candidatureTable.LogoEntreprise
		allCandidatures.Ville = candidatureTable.Ville
		allCandidatures.Created_atC = candidatureTable.Created_atC
		allCandidatures.SatutEntreprise = candidatureTable.SatutEntreprise
		res = append(res, allCandidatures)
	}
	nbreCandidature := len(res)
	return res, nbreCandidature, nil
}

// Candidaturess Réjetée

func (u UserModel) AllCandidatureDigiRejet(candidatureTable *entities.AllCandidaturesDigi, uuid int) (interface{}, int, error) {
	row, err := u.db.Query("SELECT ca.*, of.*, en.*, st.* FROM candidatureOffreGenerales ca JOIN offreGenerales of ON of.id= ca.offreGenerales_id JOIN entreprises en ON en.id= of.entreprises_id JOIN statutCandidatures st ON st.candidatureOffreGenerales_id = ca.id WHERE ca.usersDigicodeuses_id = ? AND st.satutEntreprise =? ORDER BY ca.id DESC", uuid, "Réjetée")
	if err != nil {
		return row, 0, err
	}

	defer row.Close()
	res := []entities.AllCandidaturesDigi{}
	allCandidatures := entities.AllCandidaturesDigi{}

	for row.Next() {
		row.Scan(&candidatureTable.IdC, &candidatureTable.Motivation, &candidatureTable.Cv, &candidatureTable.OffresGenerales_id, &candidatureTable.UsersDigicodeuses_id, &candidatureTable.Created_atC, &candidatureTable.Updated_atC, &candidatureTable.Id, &candidatureTable.TitrePoste, &candidatureTable.DescriptionPoste, &candidatureTable.NombrePoste, &candidatureTable.AnneeExperience, &candidatureTable.Formation, &candidatureTable.Ville, &candidatureTable.DateLimite, &candidatureTable.Entreprise_id, &candidatureTable.CreatedAt, &candidatureTable.UpdateAt, &candidatureTable.TypeContrat, &candidatureTable.IdE, &candidatureTable.NomEntreprise, &candidatureTable.SecteurActivite, &candidatureTable.Adresse, &candidatureTable.LogoEntreprise, &candidatureTable.SiteWeb, &candidatureTable.DescriptionEntreprise, &candidatureTable.NumeroEntreprise, &candidatureTable.VideoEntreprise, &candidatureTable.TailleEntreprise, &candidatureTable.Facebook, &candidatureTable.Twitter, &candidatureTable.Linkedin, &candidatureTable.Email, &candidatureTable.Slogan, &candidatureTable.Password, &candidatureTable.CreatedAtE, &candidatureTable.UpdateAtE, &candidatureTable.IdS, &candidatureTable.SatutAuto, &candidatureTable.SatutTalentManager, &candidatureTable.SatutEntreprise, &candidatureTable.CandidatureOffreGenerales_id, &candidatureTable.Created_atS, &candidatureTable.Updated_atS)
		//offres
		allCandidatures.Id = candidatureTable.Id
		allCandidatures.TitrePoste = candidatureTable.TitrePoste
		allCandidatures.TypeContrat = candidatureTable.TypeContrat
		allCandidatures.NomEntreprise = candidatureTable.NomEntreprise
		allCandidatures.LogoEntreprise = candidatureTable.LogoEntreprise
		allCandidatures.Ville = candidatureTable.Ville
		allCandidatures.Created_atC = candidatureTable.Created_atC
		allCandidatures.SatutEntreprise = candidatureTable.SatutEntreprise
		res = append(res, allCandidatures)
	}
	nbreCandidature := len(res)
	return res, nbreCandidature, nil
}

//Lorsqu'on supprime une candidature on supprime aussi sont statuts
func (u UserModel) DeleteCandidatureOffreG(id int) (int64, error) {
	del, err := u.db.Exec("DELETE FROM `candidatureOffreGenerales` WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := del.LastInsertId()
	return lastInsertId, nil
}

func (u UserModel) DeleteStatutCandidature(id int) (int64, error) {
	del, err := u.db.Exec("DELETE FROM `statutCandidatures` WHERE candidatureOffreGenerales_id=?", id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := del.LastInsertId()
	return lastInsertId, nil
}

//Update Candidature

func (u UserModel) UpdateCandidature(candidature entities.CandidatureOffreGenerales) (int64, error) {
	upt, err := u.db.Exec("UPDATE `candidatureOffreGenerales` SET `motivation`=?,`cv`=? WHERE offreGenerales_id=? AND usersDigicodeuses_id=?",
		candidature.Motivation, candidature.Cv, candidature.OffresGenerales_id, candidature.UsersDigicodeuses_id)

	// fmt.Println(upt)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := upt.LastInsertId()
	// fmt.Println(err, "modl")
	return lastInsertId, nil
}

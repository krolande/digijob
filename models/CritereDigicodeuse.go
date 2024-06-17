package models

import (
	"digiJob/entities"
)

func (u UserModel) InsertOrUpdateCritereDigi(insCritere entities.CritereDigicodeuses, uuid int) (int64, error) {
	insForm, err := u.db.Exec("INSERT INTO `critereDigicodeuses`(`type_contrat`, `localite`, `duree_contrat`, `modalite`, `metier_recherche`, `usersDigicodeuses_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE type_contrat=?, localite=?, duree_contrat=?, modalite=?, metier_recherche=?",
		insCritere.TypeContrat, insCritere.Localite, insCritere.DureeContrat, insCritere.Modalite, insCritere.MetierRecherche, insCritere.UsersDigicodeuse_id, insCritere.TypeContrat, insCritere.Localite, insCritere.DureeContrat, insCritere.Modalite, insCritere.MetierRecherche)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insForm.LastInsertId()

	return lastInsertId, nil
}

func (u UserModel) GetCritereDigi(critere *entities.CritereDigicodeuses, uuid int) (interface{}, int, error) {
	res, err := u.db.Query("SELECT * FROM `critereDigicodeuses` WHERE usersDigicodeuses_id=?", uuid)
	if err != nil {
		return res, 0, err
	}
	defer res.Close()
	resAll := entities.CritereDigicodeuses{}
	var nbre int
	istrue := false
	for res.Next() {
		res.Scan(&critere.Id, &critere.TypeContrat, &critere.Localite, &critere.DureeContrat, &critere.Modalite, &critere.MetierRecherche, &critere.UsersDigicodeuse_id, &critere.CreatedAt, &critere.UpdatedAt)
		resAll.Id = critere.Id
		resAll.TypeContrat = critere.TypeContrat
		resAll.Localite = critere.Localite
		resAll.DureeContrat = critere.DureeContrat
		resAll.Modalite = critere.Modalite
		resAll.MetierRecherche = critere.MetierRecherche
		istrue = true
	}
	if istrue {
		nbre = 10
	} else {
		nbre = 0
	}
	return resAll, nbre, err
}

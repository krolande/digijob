package models

import (
	"digiJob/entities"
)

// INSERT OR UPTADE COMPETENCE PROFESSIONNELLE
func (u UserModel) InsertCompetenceProfessionnelle(competencePro entities.CompetenceProfessionnelle, c *entities.CompetenceProfessionnelle, uuid int, id int) (int64, error) {

	row, er := u.db.Query("SELECT id, usersDigicodeuses_id FROM cvCompetences")
	if er != nil {
		return 0, er
	}
	idUser := entities.CompetenceProfessionnelle{}

	defer row.Close()
	var lastInsertId int64
	for row.Next() {
		row.Scan(&c.Id, &c.UsersDigicodeuses_id)
		idUser.UsersDigicodeuses_id = c.UsersDigicodeuses_id
		idUser.Id = c.Id
		if id == idUser.Id {
			insForm, _ := u.db.Exec("UPDATE `cvCompetences` SET `titre`=?,`pourcentage`=? WHERE id=?", competencePro.Titre, competencePro.Pourcentage, competencePro.Id)
			if er != nil {
				return 0, er
			}
			lastInsertId, _ := insForm.LastInsertId()

			return lastInsertId, nil
		} else if id == 0 {

			insForm, err := u.db.Exec("INSERT INTO `cvCompetences`(`titre`, `pourcentage`, `usersDigicodeuses_id`) VALUES (?,?,?)",
				competencePro.Titre, competencePro.Pourcentage, uuid)

			if err != nil {
				return 0, err
			}
			lastInsertId, _ := insForm.LastInsertId()

			return lastInsertId, nil
		}
	}

	return lastInsertId, nil
}

// RENVOYER LES COMPETENCES CLES SOUS FORME EDIT DANS LE FORMULAIRE
func (u UserModel) GetCompetenceProfessionnelle(competencePro *entities.CompetenceProfessionnelle, uuid int) (interface{}, int, error) {
	row, err := u.db.Query("SELECT `id`, `titre`, `pourcentage` FROM `cvCompetences` WHERE usersDigicodeuses_id=?", uuid)
	if err != nil {
		return row, 0, err
	}
	defer row.Close()
	dataDb := entities.CompetenceProfessionnelle{}
	result := []entities.CompetenceProfessionnelle{}
	var nbre int
	istrue := false
	for row.Next() {
		row.Scan(&competencePro.Id, &competencePro.Titre, &competencePro.Pourcentage)
		dataDb.Id = competencePro.Id
		dataDb.Titre = competencePro.Titre
		dataDb.Pourcentage = competencePro.Pourcentage
		result = append(result, dataDb)
		istrue = true
	}
	if istrue {
		nbre = 10
	} else {
		nbre = 0
	}
	return result, nbre, err
}

//Delete ---->>>>>> supprimer

func (u UserModel) DeleteCompetencePro(id int) (int64, error) {
	del, err := u.db.Exec("DELETE FROM `cvCompetences` WHERE id=?", id)
	if err != nil {
		// fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := del.LastInsertId()
	return lastInsertId, nil
}

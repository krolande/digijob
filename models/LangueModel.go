package models

import (
	"digiJob/entities"
	"fmt"
	"strconv"
)

// /////////// Renvoyer les langues de la Langue por selection //////////
func (u UserModel) GetAllLangues(langue *entities.Langues) (interface{}, error) {
	row, err := u.db.Query("SELECT * FROM langues")

	langueAll := entities.Langues{}
	result := []entities.Langues{}
	if err != nil {
		return result, err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&langue.Id, &langue.Libelle, &langue.CreatedAt, langue.UpdatedAt)
		langueAll.Id = langue.Id
		langueAll.Libelle = langue.Libelle
		result = append(result, langueAll)
	}
	// fmt.Println(result)
	return result, err
}

/// ************** RENVOYER LES LANGUES SELECTIONNER PAR LA DIGICODEUSE SUR LA PLATEFORME************** ///*

func (u UserModel) GetLanguesUser(langues *entities.Langues, uuid int) (interface{}, int, error) {
	row, err := u.db.Query("SELECT p.* FROM langues p JOIN langue_usersDigicodeuses lu ON lu.langues_id = p.id WHERE lu.usersDigicodeuses_id = ?",
		uuid)
	if err != nil {
		return row, 0, err
	}
	defer row.Close()
	dataDb := entities.Langues{}
	result := []entities.Langues{}

	var nbre int
	istrue := false
	for row.Next() {
		row.Scan(&langues.Id, &langues.Libelle, &langues.CreatedAt, langues.UpdatedAt)
		dataDb.Id = langues.Id
		dataDb.Libelle = langues.Libelle
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

///////////// ***************Enregistrer et mettre a jour les langues*************** /////////////

func (u UserModel) InsertOrUpdateLangue(Libelle []string, langueDb *entities.LangueUsersDigicodeuses, uuid int) (int64, error) {
	// recuperer les langues
	row, err := u.db.Query("SELECT id, langues_id, usersDigicodeuses_id FROM langue_usersDigicodeuses")
	langueDbAll := entities.LangueUsersDigicodeuses{}
	resultLangue := []entities.LangueUsersDigicodeuses{}
	if err != nil {
		return 0, err
	}
	defer row.Close()

	isVerif := true
	for row.Next() {
		if isVerif {
			row.Scan(&langueDb.Id, &langueDb.Langues_id, &langueDb.UsersDigicodeuses_id)

			langueDbAll.UsersDigicodeuses_id = langueDb.UsersDigicodeuses_id
			langueDbAll.Id = langueDb.Id
			if uuid == langueDbAll.UsersDigicodeuses_id {
				insForm, err := u.db.Query("SELECT langues_id FROM langue_usersDigicodeuses WHERE usersDigicodeuses_id = ?", uuid)
				if err != nil {
					fmt.Println(err)
					return 0, err
				}
				row.Close()
				for insForm.Next() {
					insForm.Scan(&langueDb.Langues_id)
					langueDbAll.Langues_id = langueDb.Langues_id
					resultLangue = append(resultLangue, langueDbAll)
					// fmt.Println(resultLangue, "resultLangue")
				}

				isVerif = false
			}
		}

	}

	//////// Fonction modification et ajout
	for a := 0; a < len(Libelle); a++ {

		isVerification := false
		c, _ := strconv.Atoi(Libelle[a])
		for b := 0; b < len(resultLangue); b++ {
			d := resultLangue[b].Langues_id

			if c == d {
				isVerification = true
			}
		}
		if isVerification {
			fmt.Println("UPDATE SANS AJOUT, CompetenceClesModel")
		} else {
			fmt.Println("on doit l'enregistrer")
			_, err = u.db.Exec("INSERT INTO langue_usersDigicodeuses (langues_id, usersDigicodeuses_id) values(?,?)",
				c, uuid)
			// fmt.Println(err, "2")
			if err != nil {
				return 0, err
			}
		}
	}

	// ****************************** delete ***************************** //

	isEgale := false
	for b := 0; b < len(resultLangue); b++ {
		d := resultLangue[b].Langues_id

		for a := 0; a < len(Libelle); a++ {
			c, _ := strconv.Atoi(Libelle[a])
			if c == d {
				isEgale = true
				continue
			}
		}
		if isEgale {
			isEgale = false
			continue
		} else {
			_, err := u.db.Exec("DELETE FROM `langue_usersDigicodeuses` WHERE langues_id=? AND usersDigicodeuses_id =?", d, uuid)
			if err != nil {
				// fmt.Println(err)
				return 0, err
			}
		}
	}

	return 1, nil
}

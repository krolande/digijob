package models

import (
	"digiJob/entities"
	"fmt"
	"strconv"
)

// type CompetenceCles struct {
// 	Id      int
// 	Libelle string
// }
//////// PERMET DE RETOURNER TOUTES LES LIBELLES DE LA TABLE COMPETENCECLES AFIN DE LES AFFICHER DANS LE SELECT ////////
func (u UserModel) AllCompetenceCles(competencecles *entities.CompetenceCles) (interface{}, error) {
	row, err := u.db.Query("SELECT * FROM competenceCles")

	competencesAll := entities.CompetenceCles{}
	res := []entities.CompetenceCles{}
	if err != nil {
		return res, err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&competencecles.Id, &competencecles.Libelle, &competencecles.CreatedAt, &competencecles.UpdatedAt)
		competencesAll.Id = competencecles.Id
		competencesAll.Libelle = competencecles.Libelle
		res = append(res, competencesAll)
	}
	// fmt.Println(res)

	return res, err

}

// SELECTIONNER LES COMPETENCECLES DE LA DIGICODEUSE SUR LE PROFIL-EDIT // <<<<< MANY TO MANY >>>>>
func (u UserModel) GetAllCompetenceClesOfUser(competencecles *entities.CompetenceCles, uuid int) (interface{}, error) {
	row, err := u.db.Query("SELECT p.* FROM competenceCles p JOIN competenceCle_usersDigicodeuses cu ON cu.competenceCles_id = p.id WHERE cu.usersDigicodeuses_id = ?",
		uuid)

	competencesAll := entities.CompetenceCles{}
	res := []entities.CompetenceCles{}
	if err != nil {
		return res, err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&competencecles.Id, &competencecles.Libelle, &competencecles.CreatedAt, &competencecles.UpdatedAt)
		competencesAll.Id = competencecles.Id
		competencesAll.Libelle = competencecles.Libelle
		competencesAll.CreatedAt = competencecles.CreatedAt
		competencesAll.UpdatedAt = competencecles.UpdatedAt
		res = append(res, competencesAll)
	}
	// fmt.Println(res)

	return res, err
}

// obtenir un tableau de libelle des compétences clées de la Digi
func (u UserModel) GetLibelleCompetenceCle(competencecles *entities.CompetenceCles, uuid int) ([]string, error) {
	row, err := u.db.Query("SELECT p.* FROM competenceCles p JOIN competenceCle_usersDigicodeuses cu ON cu.competenceCles_id = p.id WHERE cu.usersDigicodeuses_id = ?",
		uuid)

	// competencesAll := entities.CompetenceCles{}
	var libelle string
	res := make([]string, 0)
	if err != nil {
		return res, err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&competencecles.Id, &competencecles.Libelle, &competencecles.CreatedAt, &competencecles.UpdatedAt)
		libelle = competencecles.Libelle
		res = append(res, libelle)
	}
	// fmt.Println(res)

	return res, err
}

/////// PERMET D'ENREGISTRER LES COMPETENCES DE LA DIGICODEUSE ET DE LES METTRENT A JOUR //////
func (u UserModel) CreateUpdateCompetenceCles(Libelle []string, c *entities.CompetenceCleUsersDigicodeuses, uuid int) (int64, error) {
	row, err := u.db.Query("SELECT usersDigicodeuses_id, competenceCles_id FROM competenceCle_usersDigicodeuses")
	idAll := entities.CompetenceCleUsersDigicodeuses{}
	resId := []entities.CompetenceCleUsersDigicodeuses{}
	if err != nil {
		// fmt.Println(err)
		return 0, err
	}

	defer row.Close()
	isVerif := true
	for row.Next() {
		if isVerif {
			row.Scan(&c.UsersDigicodeuses_id, &c.CompetenceCles_id)
			idAll.UsersDigicodeuses_id = c.UsersDigicodeuses_id
			// idAll.CompetenceCles_id = c.CompetenceCles_id
			// resId = append(resId, idAll)
			if uuid == idAll.UsersDigicodeuses_id {
				// fmt.Println("ok1")
				insForm, err := u.db.Query("SELECT competenceCles_id FROM competenceCle_usersDigicodeuses WHERE usersDigicodeuses_id = ?", uuid)
				if err != nil {
					return 0, err
				}
				for insForm.Next() {
					insForm.Scan(&c.CompetenceCles_id)
					idAll.CompetenceCles_id = c.CompetenceCles_id
					resId = append(resId, idAll)
				}
				isVerif = false
			}
		}

	}
	// ************************** Ajout or update **********************
	for a := 0; a < len(Libelle); a++ {
		isVerification := false
		c, _ := strconv.Atoi(Libelle[a])
		for b := 0; b < len(resId); b++ {
			d := resId[b].CompetenceCles_id

			if c == d {
				isVerification = true
			}
		}

		if isVerification {
			fmt.Println("UPDATE SANS AJOUT, CompetenceClesModel")
		} else {
			fmt.Println("on doit l'enregistrer")
			_, err := u.db.Exec("INSERT INTO competenceCle_usersDigicodeuses (usersDigicodeuses_id, competenceCles_id) values(?,?)",
				uuid, c)
			// fmt.Println(err, "2")
			if err != nil {
				return 0, err
			}

		}

	}

	///////////// ************************** Delete ******************** //////////////

	// ****************************** delete ***************************** //

	isEgale := false
	for b := 0; b < len(resId); b++ {
		d := resId[b].CompetenceCles_id

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
			_, err := u.db.Exec("DELETE FROM `competenceCle_usersDigicodeuses` WHERE competenceCles_id=? AND usersDigicodeuses_id =?", d, uuid)
			if err != nil {
				// fmt.Println(err)
				return 0, err
			}
		}
	}

	return 1, nil
}

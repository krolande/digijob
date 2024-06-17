package models

import (
	"digiJob/entities"
	"fmt"
)

// Create function

func (u UserModel) InsertPrixCertification(prixCertification entities.PrixCertification) (int64, error) {
	insForm, err := u.db.Exec("INSERT INTO `cvPrixCertifications`(`titrePrix`, `nomEtablissement`, `dateObtention`, `description`, `usersDigicodeuses_id`) VALUES (?,?,?,?,?)",
		prixCertification.TitrePrix, prixCertification.NomEtablissement, prixCertification.DateObtention, prixCertification.Description, prixCertification.UsersDigicodeuse_id)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insForm.LastInsertId()

	return lastInsertId, nil

}

// Read information function

func (u UserModel) ReadPrixCertification(prixCertifications *entities.PrixCertification, uuid int) (interface{}, int, error) {
	result, err := u.db.Query("SELECT * FROM `cvPrixCertifications` WHERE usersDigicodeuses_id=?", uuid)

	if err != nil {
		return result, 0, err
	}

	var nbre int
	istrue := false
	prixCertification := entities.PrixCertification{}
	AllPrixCertification := []entities.PrixCertification{}
	defer result.Close()

	for result.Next() {
		result.Scan(&prixCertifications.Id, &prixCertifications.TitrePrix, &prixCertifications.NomEtablissement, &prixCertifications.DateObtention, &prixCertifications.Description, &prixCertifications.UsersDigicodeuse_id, &prixCertifications.CreatedAt, &prixCertifications.UpdatedAt)
		prixCertification.Id = prixCertifications.Id
		prixCertification.TitrePrix = prixCertifications.TitrePrix
		prixCertification.NomEtablissement = prixCertifications.NomEtablissement
		prixCertification.DateObtention = prixCertifications.DateObtention
		prixCertification.Description = prixCertifications.Description
		AllPrixCertification = append(AllPrixCertification, prixCertification)
		istrue = true
	}
	if istrue {
		nbre = 10
	} else {
		nbre = 0
	}
	return AllPrixCertification, nbre, err
}

func (u UserModel) UpdatePrixCertification(prixCertification entities.PrixCertification) (int64, error) {
	insForm, err := u.db.Exec("UPDATE `cvPrixCertifications` SET `titrePrix`=?,`nomEtablissement`=?,`dateObtention`=?,`description`=? WHERE id=?",
		prixCertification.TitrePrix, prixCertification.NomEtablissement, prixCertification.DateObtention, prixCertification.Description, prixCertification.Id)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insForm.LastInsertId()

	return lastInsertId, nil

}

func (u UserModel) DeletePrixCertification(id int) (int64, error) {
	del, err := u.db.Exec("DELETE FROM `cvPrixCertifications` WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := del.LastInsertId()
	return lastInsertId, nil
}

package models

import (
	"digiJob/entities"
	"fmt"
)

func (u UserModel) InsertFormation(formation entities.Formation, uuid int) (int64, error) {
	result, err := u.db.Exec("INSERT INTO `cvFormations`(`moisDebut`, `moisFin`, `anneeDebut`, `anneeFin`, `now`, `titreFormation`, `nomEcole`, `descriptionFormation`, `usersDigicodeuses_id`) VALUES (?,?,?,?,?,?,?,?,?)",
		formation.MoisDebut, formation.MoisFin, formation.AnneeDebut, formation.AnneeFin, formation.Now, formation.TitreFormation, formation.NomEcole, formation.Description, uuid)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil
}

func (u UserModel) ReadForamtion(formations *entities.Formation, uuid int) (interface{}, int, error) {
	insForm, err := u.db.Query("SELECT * FROM `cvFormations` WHERE usersDigicodeuses_id=?", uuid)

	if err != nil {
		return insForm, 0, err
	}

	formation := entities.Formation{}
	AllFormation := []entities.Formation{}
	defer insForm.Close()

	var nbre int
	istrue := false
	for insForm.Next() {
		insForm.Scan(&formations.Id, &formations.MoisDebut, &formations.MoisFin, &formations.AnneeDebut, &formations.AnneeFin, &formations.Now, &formations.TitreFormation, &formations.NomEcole, &formations.Description, &formations.UsersDigicodeuses_id, &formations.Created_at, &formations.Updated_at)
		formation.Id = formations.Id
		formation.MoisDebut = formations.MoisDebut
		formation.MoisFin = formations.MoisFin
		formation.AnneeDebut = formations.AnneeDebut
		formation.AnneeFin = formations.AnneeFin
		formation.Now = formations.Now
		formation.TitreFormation = formations.TitreFormation
		formation.NomEcole = formations.NomEcole
		formation.Description = formations.Description
		AllFormation = append(AllFormation, formation)
		istrue = true
	}
	if istrue {
		nbre = 20
	} else {
		nbre = 0
	}
	return AllFormation, nbre, err

}

func (u UserModel) UpdateFormation(formation entities.Formation) (int64, error) {
	upt, err := u.db.Exec("UPDATE `cvFormations` SET `moisDebut`=?,`moisFin`=?,`anneeDebut`=?,`anneeFin`=?,`now`=?,`titreFormation`=?,`nomEcole`=?,`descriptionFormation`=? WHERE id=?",
		formation.MoisDebut, formation.MoisFin, formation.AnneeDebut, formation.AnneeFin, formation.Now, formation.TitreFormation, formation.NomEcole, formation.Description, formation.Id)

	// fmt.Println(upt)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := upt.LastInsertId()
	// fmt.Println(err, "modl")
	return lastInsertId, nil
}

func (u UserModel) DeleteFormation(uuid int) (int64, error) {
	del, err := u.db.Exec("DELETE FROM `cvFormations` WHERE id=?", uuid)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := del.LastInsertId()
	return lastInsertId, nil
}

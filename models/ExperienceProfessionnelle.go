package models

import (
	"digiJob/entities"
	"fmt"
)

func (u UserModel) InsertExpProfessionnelle(experience entities.ExperienceProfessionnelle, uuid int) (int64, error) {
	result, err := u.db.Exec("INSERT INTO `cvExperiences`(`moisDebut`, `moisFin`, `anneeDebut`, `anneeFin`, `now`, `titreExperience`, `nomEntreprise`, typeContrat, `descriptionExperience`, `usersDigicodeuses_id`) VALUES (?,?,?,?,?,?,?,?,?,?)",
		experience.MoisDebut, experience.MoisFin, experience.AnneeDebut, experience.AnneeFin, experience.Now, experience.TitreExperience, experience.NomEntreprise, experience.TypeContrat, experience.Description, uuid)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil
}

func (u UserModel) ReadExperience(experiences *entities.ExperienceProfessionnelle, uuid int) (interface{}, int, error) {
	insForm, err := u.db.Query("SELECT * FROM `cvExperiences` WHERE usersDigicodeuses_id=?", uuid)

	if err != nil {
		return insForm, 0, err
	}

	experience := entities.ExperienceProfessionnelle{}
	AllExperience := []entities.ExperienceProfessionnelle{}
	defer insForm.Close()

	var nbre int
	istrue := false
	for insForm.Next() {
		insForm.Scan(&experiences.Id, &experiences.MoisDebut, &experiences.MoisFin, &experiences.AnneeDebut, &experiences.AnneeFin, &experiences.Now, &experiences.TitreExperience, &experiences.NomEntreprise, &experiences.TypeContrat, &experiences.Description, &experiences.UsersDigicodeuses_id, &experiences.Created_at, &experiences.Updated_at)
		experience.Id = experiences.Id
		experience.MoisDebut = experiences.MoisDebut
		experience.MoisFin = experiences.MoisFin
		experience.AnneeDebut = experiences.AnneeDebut
		experience.AnneeFin = experiences.AnneeFin
		experience.Now = experiences.Now
		experience.TitreExperience = experiences.TitreExperience
		experience.NomEntreprise = experiences.NomEntreprise
		experience.TypeContrat = experiences.TypeContrat
		experience.Description = experiences.Description
		AllExperience = append(AllExperience, experience)
		istrue = true
	}

	if istrue {
		nbre = 18
	} else {
		nbre = 0
	}
	return AllExperience, nbre, err

}

func (u UserModel) UpdateExperience(experience entities.ExperienceProfessionnelle) (int64, error) {
	upt, err := u.db.Exec("UPDATE `cvExperiences` SET `moisDebut`=?,`moisFin`=?,`anneeDebut`=?,`anneeFin`=?,`now`=?,`titreExperience`=?,`nomEntreprise`=?,`typeContrat`=?,`descriptionExperience`=? WHERE id=?",
		experience.MoisDebut, experience.MoisFin, experience.AnneeDebut, experience.AnneeFin, experience.Now, experience.TitreExperience, experience.NomEntreprise, experience.TypeContrat, experience.Description, experience.Id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := upt.LastInsertId()
	// fmt.Println(err, "modl")
	return lastInsertId, nil
}

func (u UserModel) DeleteExperience(uuid int) (int64, error) {
	del, err := u.db.Exec("DELETE FROM `cvExperiences` WHERE id=?", uuid)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastInsertId, _ := del.LastInsertId()
	return lastInsertId, nil
}

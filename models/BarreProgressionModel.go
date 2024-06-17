package models

import (
	"digiJob/entities"
)

func (u UserModel) InsertBarreProgressionCv(barre entities.BarreProgressionCv, uuid int) (int64, error) {
	insInfo, err := u.db.Exec("INSERT INTO `barreprogressioncvs`(`nbreFormation`, `nbreExpeProf`, `nbreComptProf`, `nbrePrixCertificat`, `nbreLangue`, `usersDigicodeuses_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE nbreFormation=?, nbreExpeProf=?, nbreComptProf=?, nbrePrixCertificat=?, nbreLangue=?",
		barre.NbreFormation, barre.NbreExpeProf, barre.NbreComptProf, barre.NbrePrixCertificat, barre.NbreLangue, uuid, barre.NbreFormation, barre.NbreExpeProf, barre.NbreComptProf, barre.NbrePrixCertificat, barre.NbreLangue)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insInfo.LastInsertId()

	return lastInsertId, nil
}

func (u UserModel) InsertBarreProgresInfoPerso(barre entities.BarreProgressionInfoPerso, uuid int) (int64, error) {
	insInfo, err := u.db.Exec("INSERT INTO `barreprogressionprofils`(`nbreInfoPerso`, `usersDigicodeuses_id`) VALUES (?,?) ON DUPLICATE KEY UPDATE  nbreInfoPerso=?",
		barre.NbreInfoPerso, uuid, barre.NbreInfoPerso)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insInfo.LastInsertId()

	return lastInsertId, nil
}
func (u UserModel) InsertBarreProgresCritere(barre entities.BarreProgressionCritere, uuid int) (int64, error) {
	insInfo, err := u.db.Exec("INSERT INTO `barreprogressioncriteres`(`nbreCritere`, `usersDigicodeuses_id`) VALUES (?,?) ON DUPLICATE KEY UPDATE  nbreCritere=?",
		barre.NbreCritere, uuid, barre.NbreCritere)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insInfo.LastInsertId()

	return lastInsertId, nil
}

// get les données de la barre de progression

func (u UserModel) GetBarreProgressionInfoPerso(nbre *entities.BarreProgressionInfoPerso, uuid int) (int, error) {
	res, err := u.db.Query("SELECT 	nbreInfoPerso FROM `barreprogressionprofils` WHERE usersDigicodeuses_id=?", uuid)
	if err != nil {
		return 0, err
	}
	defer res.Close()
	resAll := entities.BarreProgressionInfoPerso{}

	for res.Next() {
		res.Scan(&nbre.NbreInfoPerso)
		resAll.NbreInfoPerso = nbre.NbreInfoPerso
	}
	return resAll.NbreInfoPerso, err
}

func (u UserModel) GetBarreProgressionCv(nbre *entities.BarreProgressionCv, uuid int) (int, error) {
	res, err := u.db.Query("SELECT nbreFormation, nbreExpeProf, nbreComptProf, nbrePrixCertificat, nbreLangue FROM `barreprogressioncvs` WHERE usersDigicodeuses_id=?", uuid)
	if err != nil {
		return 0, err
	}
	defer res.Close()
	resAll := entities.BarreProgressionCv{}

	for res.Next() {
		res.Scan(&nbre.NbreFormation, &nbre.NbreExpeProf, &nbre.NbreComptProf, &nbre.NbrePrixCertificat, &nbre.NbreLangue)
		resAll.NbreFormation = nbre.NbreFormation
		resAll.NbreExpeProf = nbre.NbreExpeProf
		resAll.NbreComptProf = nbre.NbreComptProf
		resAll.NbrePrixCertificat = nbre.NbrePrixCertificat
		resAll.NbreLangue = nbre.NbreLangue
	}
	totalNbreCv := resAll.NbreFormation + resAll.NbreExpeProf + resAll.NbreComptProf + resAll.NbrePrixCertificat + resAll.NbreLangue
	return totalNbreCv, err
}

func (u UserModel) GetBarreProgressionCritere(nbre *entities.BarreProgressionCritere, uuid int) (int, error) {
	res, err := u.db.Query("SELECT 	nbreCritere FROM `barreprogressioncriteres` WHERE usersDigicodeuses_id=?", uuid)
	if err != nil {
		return 0, err
	}
	defer res.Close()
	resAll := entities.BarreProgressionCritere{}

	for res.Next() {
		res.Scan(&nbre.NbreCritere)
		resAll.NbreCritere = nbre.NbreCritere
	}
	return resAll.NbreCritere, err
}

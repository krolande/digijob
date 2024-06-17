package models

import (
	"digiJob/entities"
)

// Obetenir les competences clées de l'entreprise en fonction de l'offre générale
func (u UserModel) GetCompetenceClesOffreG(competencecles *entities.CompetenceCles, offres_id int) ([]string, error) {
	row, err := u.db.Query("SELECT p.* FROM competenceCles p JOIN competenceCle_offreGenerales cu ON cu.competenceCles_id = p.id WHERE cu.offreGenerales_id = ?",
		offres_id)

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

package recruteursmodel

import (
	"database/sql"
	"digiJob/config"
	recruteursentities "digiJob/entities/recruteursEntities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}
	return &UserModel{
		db: conn,
	}
}

func (u UserModel) WhereEntreprise(entreprise *recruteursentities.Entreprise, fieldName, fieldValue string) error {
	//recupère les données utilisateurs
	row, err := u.db.Query("select id, nomEntreprise, email, password from entreprises where "+fieldName+" = ? limit 1", fieldValue)
	if err != nil {
		return err
	}

	defer row.Close()

	//recupère les données et les saisir dans les entités utilisaturs
	for row.Next() {
		row.Scan(&entreprise.Id, &entreprise.Nom, &entreprise.Email, &entreprise.Password)
	}
	return nil
}

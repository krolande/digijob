package entities

import (
	"database/sql"
)

type UsersDigicodeuses struct {
	Id              int
	Nom             string `validate:"required" label:"nom"`
	Prenoms         string `validate:"required,prenoms,isunique=users-prenoms"`
	Email           string `validate:"required,email,isunique=users-email"`
	Username        string `validate:"required,gte=3,isunique=users-username"`
	Pays            string `validate:"required,pays,isunique=users-pays"`
	Ville           string
	DateNaissance   string
	Indicatif       string `validate:"required,indicatif,isunique=users-indicatif"`
	NumeroTelephone string
	Nationalite     string
	Cv              string
	Photo           string
	TitrePosteCv    string
	Facebook        string
	Twitter         string
	Linkedin        string
	Autorisation    string
	Adresse         string
	NiveauEtude     string
	Description     string
	Porfolio        string
	AnneeExperience string
	Temoignage      string
	Statut          string
	CreatedAt       string
	UpdatedAt       string
	Password        string
	Cpassword       string `validate:"required,eqfield=Password" label:"Konfirmasi Password"`
}

//Permet de recuperer les données de la BD dans la fonction GetUser()
type UsersDigi struct {
	Id              int
	Nom             string `validate:"required" label:"nom"`
	Prenoms         string `validate:"required,prenoms,isunique=users-prenoms"`
	Email           string `validate:"required,email,isunique=users-email"`
	Username        string `validate:"required,gte=3,isunique=users-username"`
	Pays            string `validate:"required,pays,isunique=users-pays"`
	Ville           string
	DateNaissance   sql.NullString
	Indicatif       sql.NullString `validate:"required,indicatif,isunique=users-indicatif"`
	NumeroTelephone sql.NullString
	Nationalite     sql.NullString
	Cv              sql.NullString
	Photo           sql.NullString
	TitrePosteCv    sql.NullString
	Facebook        sql.NullString
	Twitter         sql.NullString
	Linkedin        sql.NullString
	Autorisation    sql.NullString
	Adresse         sql.NullString
	NiveauEtude     sql.NullString
	Description     sql.NullString
	Porfolio        sql.NullString
	Temoignage      sql.NullString
	Statut          sql.NullString
	AnneeExperience sql.NullString
	CreatedAt       string
	UpdatedAt       string
	Password        string
	Cpassword       string `validate:"required,eqfield=Password" label:"Konfirmasi Password"`
}

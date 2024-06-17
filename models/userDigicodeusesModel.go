package models

import (
	"database/sql"
	"fmt"

	"digiJob/config"
	"digiJob/entities"
)

type UserModel struct {
	db *sql.DB
}

//Permet la communication des données de l'utilisateur avec la base de donnée
func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}
	return &UserModel{
		db: conn,
	}
}

//////Pour la connexion
func (u UserModel) Where(user *entities.UsersDigicodeuses, fieldName, fieldValue string) error {
	//recupère les données utilisateurs
	row, err := u.db.Query("select id, nom, prenoms, username, email, password, pays from usersDigicodeuses where "+fieldName+" = ? limit 1", fieldValue)
	if err != nil {
		return err
	}

	defer row.Close()

	//recupère les données et les saisir dans les entités utilisaturs
	for row.Next() {
		row.Scan(&user.Id, &user.Nom, &user.Prenoms, &user.Username, &user.Email, &user.Password, &user.Pays)
	}
	return nil
}

func (u UserModel) Create(user entities.UsersDigicodeuses) (int64, error) {
	// processus de requete pour inserer les éléments
	result, err := u.db.Exec("insert into usersDigicodeuses (nom, prenoms, username, email, indicatif, numero_telephone, ville, pays, autorisation, password) values(?,?,?,?,?,?,?,?,?,?)",
		user.Nom, user.Prenoms, user.Username, user.Email, user.Indicatif, user.NumeroTelephone, user.Ville, user.Pays, user.Autorisation, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil

}

func (u UserModel) Update(user entities.UsersDigicodeuses) (int64, error) {
	result, err := u.db.Exec("UPDATE usersDigicodeuses SET nom=?, prenoms=?, username=?, email=?, pays=?, ville=?, date_naissance=?, indicatif=?, numero_telephone=?, nationalite=?, cv=?, photo=?, titrePosteCv=?, portfolio=?, facebook=?, twitter=?, linkedin=?, temoignage=?, description=?, niveau_etude=?, annee_experience=?, adresse=?, statut=? WHERE id=?",
		user.Nom, user.Prenoms, user.Username, user.Email, user.Pays, user.Ville, user.DateNaissance, user.Indicatif, user.NumeroTelephone, user.Nationalite, user.Cv, user.Photo, user.TitrePosteCv, user.Porfolio, user.Facebook, user.Twitter, user.Linkedin, user.Temoignage, user.Description, user.NiveauEtude, user.AnneeExperience, user.Adresse, user.Statut, user.Id)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	// fmt.Println(err, "modl")
	return lastInsertId, nil
}

func (u UserModel) GetUser(user *entities.UsersDigi, id int) (interface{}, error) {
	res := entities.UsersDigicodeuses{}
	// result, err := u.db.Query("SELECT id, nom, prenoms, email, password, pays, ville, date_naissance, indicatif, numero_telephone, nationalite, cv, photo, titrePosteCv, portfolio, facebook, twitter, linkedin, temoignage, autorisation, description, niveau_etude, adresse, created_at, updated_at FROM usersDigicodeuses WHERE id=?", id)
	result, err := u.db.Query("SELECT * FROM `usersDigicodeuses` WHERE `id`=?", id)

	if err != nil {
		return res, err
	}

	for result.Next() {
		err = result.Scan(&user.Id, &user.Nom, &user.Prenoms, &user.Username, &user.Email, &user.Password, &user.Username, &user.Pays, &user.Ville, &user.DateNaissance, &user.Indicatif, &user.NumeroTelephone, &user.Nationalite, &user.Cv, &user.Photo, &user.TitrePosteCv, &user.Porfolio, &user.Facebook, &user.Twitter, &user.Linkedin, &user.Temoignage, &user.Autorisation, &user.Description, &user.NiveauEtude, &user.AnneeExperience, &user.Adresse, &user.Statut, &user.CreatedAt, &user.UpdatedAt)
		// err = result.Scan(&user.Id, &user.Nom, &user.Prenoms, &user.Email, &user.Password)
		if err != nil {

			fmt.Println(err)
			return res, err
		}
		// fmt.Println(user)
		res.Id = user.Id
		res.Nom = user.Nom
		res.Prenoms = user.Prenoms
		res.Username = user.Username
		res.Email = user.Email
		res.Pays = user.Pays
		res.Ville = user.Ville
		res.DateNaissance = user.DateNaissance.String
		res.Indicatif = user.Indicatif.String
		res.NumeroTelephone = user.NumeroTelephone.String
		res.Nationalite = user.Nationalite.String
		res.Cv = user.Cv.String
		res.Photo = user.Photo.String
		res.TitrePosteCv = user.TitrePosteCv.String
		res.Porfolio = user.Porfolio.String
		res.Facebook = user.Facebook.String
		res.Twitter = user.Twitter.String
		res.Linkedin = user.Linkedin.String
		res.Temoignage = user.Temoignage.String
		res.Autorisation = user.Autorisation.String
		res.Description = user.Description.String
		res.NiveauEtude = user.NiveauEtude.String
		res.Adresse = user.Adresse.String
		res.Statut = user.Statut.String
		res.Password = user.Password
		res.AnneeExperience = user.AnneeExperience.String

	}
	// fmt.Println(res)
	return res, err
}

func (u UserModel) UpdatePassword(newPassword string, uuid int) (int64, error) {
	result, err := u.db.Exec("UPDATE usersDigicodeuses SET password=? WHERE id=?",
		newPassword, uuid)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	// fmt.Println(err, "modl")
	return lastInsertId, nil
}

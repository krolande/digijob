package entities

import "database/sql"

type Entreprises struct {
	Id                    int
	NomEntreprise         string
	SecteurActivite       string
	Adresse               string
	Logo                  string
	SiteWeb               sql.NullString
	DescriptionEntreprise sql.NullString
	NumeroEntreprise      sql.NullString
	VideoEntreprise       sql.NullString
	TailleEntreprise      sql.NullString
	Facebook              sql.NullString
	Twitter               sql.NullString
	Linkedin              sql.NullString
	Email                 sql.NullString
	Slogan                sql.NullString
	Password              sql.NullString
	CreatedAt             string
	UpdateAt              string
	//Contenue
	IdC            int
	NomPhoto       sql.NullString
	PhotoUpload    string
	Entreprises_id int
	Created_At     string
	Update_At      string
}

//Pour recuperer les entreprises avec des types normales
type Entreprise struct {
	Id                    int
	NomEntreprise         string
	SecteurActivite       string
	Adresse               string
	Logo                  string
	SiteWeb               string
	DescriptionEntreprise string
	NumeroEntreprise      string
	VideoEntreprise       string
	TailleEntreprise      string
	Facebook              string
	Twitter               string
	Linkedin              string
	Email                 string
	Slogan                string
	ContenueEntreprise    []ContenueEntreprises
	Password              string
	CreatedAt             string
	UpdateAt              string
	//Contenue
	IdC            int
	NomPhoto       string
	PhotoUpload    string
	Entreprises_id int
	Created_At     string
	Update_At      string
}

type ContenueEntreprises struct {
	Id             int
	NomPhoto       string
	PhotoUpload    string
	Entreprises_id int
	Created_At     string
	Update_At      string
}

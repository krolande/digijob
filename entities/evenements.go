package entities

import "database/sql"

type Evenements struct {
	Id                    int
	TitreEvent            string
	DescriptionEvent      string
	FileUpload            string
	DateEvent             string
	TypeEvent             string
	Heure                 string
	FraisInscriptionEvent string
	ModaliteEvent         string
	Lien                  string
	LieuEvent             string
	DateLimite            string
	Entreprises_id        int
	Created_At            string
	Update_At             string
}

type EvenementBd struct {
	Id                    int
	TitreEvent            string
	DescriptionEvent      string
	FileUpload            string
	DateEvent             string
	TypeEvent             sql.NullString
	Heure                 sql.NullString
	FraisInscriptionEvent string
	ModaliteEvent         sql.NullString
	Lien                  sql.NullString
	LieuEvent             sql.NullString
	DateLimite            sql.NullString
	Entreprises_id        int
	Created_At            string
	Update_At             string
}

// EvenementEntreprise
type EvenementEntreprises struct {
	Id                    int
	TitreEvent            string
	DescriptionEvent      string
	FileUpload            string
	DateEvent             string
	DateLimite            string
	DateLimites           string
	TypeEvent             string
	Heure                 string
	FraisInscriptionEvent string
	FileLogo              string
	ModaliteEvent         string
	Lien                  string
	LieuEvent             string
	Entreprises_id        int
	CreatedAt             string
	UpdateAt              string

	//Entreprise
	IdE                   int
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
	Password              sql.NullString
	CreatedAtE            string
	UpdateAtE             string
}

type EvenementEntrepriseBd struct {
	Id                    int
	TitreEvent            string
	DescriptionEvent      string
	FileUpload            string
	DateEvent             string
	DateLimite            sql.NullString
	TypeEvent             sql.NullString
	Heure                 sql.NullString
	FraisInscriptionEvent sql.NullString
	FileLogo              string
	ModaliteEvent         sql.NullString
	Lien                  sql.NullString
	LieuEvent             string
	Entreprises_id        int
	CreatedAt             string
	UpdateAt              string

	//Entreprise
	IdE                   int
	NomEntreprise         string
	SecteurActivite       string
	Adresse               string
	Logo                  string
	SiteWeb               sql.NullString
	DescriptionEntreprise string
	NumeroEntreprise      string
	VideoEntreprise       sql.NullString
	TailleEntreprise      sql.NullString
	Facebook              sql.NullString
	Twitter               sql.NullString
	Linkedin              sql.NullString
	Email                 sql.NullString
	Slogan                sql.NullString
	Password              sql.NullString
	CreatedAtE            string
	UpdateAtE             string
}

// Participant des évènements

type ParticipantEvent struct {
	Id                   int
	Evenements_id        int
	UsersDigicodeuses_id int
	CreatedAt            string
	UpdateAt             string
}

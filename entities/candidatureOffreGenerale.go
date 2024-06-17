package entities

import "database/sql"

// ************************* Renvoyer les candidatures ****************************
type AllCandidaturesDigi struct {
	//offreGenerale
	Id               int
	TitrePoste       string
	DescriptionPoste string
	NombrePoste      int
	AnneeExperience  string
	Formation        string
	Ville            string
	DateLimite       string
	Entreprise_id    int
	TypeContrat      string
	CreatedAt        string
	UpdateAt         string

	//Entreprises
	IdE                   int
	LogoEntreprise        string
	NomEntreprise         string
	SecteurActivite       string
	Adresse               string
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
	CreatedAtE            string
	UpdateAtE             string

	//Candidature
	IdC                  int
	Motivation           string
	Cv                   string
	OffresGenerales_id   int
	UsersDigicodeuses_id int
	Created_atC          string
	Updated_atC          string

	//Statut Candidature

	IdS                          int
	SatutAuto                    string
	SatutTalentManager           string
	SatutEntreprise              string
	CandidatureOffreGenerales_id int
	Created_atS                  string
	Updated_atS                  string

	//CopetenceCle

	IdCc         int
	Libelle      string
	LibelleTab   [][]CompetenceCles
	Created_atCc string
	Updated_atCc string
}

//******************* Candidatures ****************//

type CandidatureOffreGenerales struct {
	Id                   int
	Motivation           string
	Cv                   string
	OffresGenerales_id   int
	UsersDigicodeuses_id int
	Created_at           string
	Updated_at           string
}

type StatutCandidature struct {
	Id                           int
	SatutAuto                    string
	SatutTalentManager           string
	SatutEntreprise              string
	CandidatureOffreGenerales_id int
	Created_at                   string
	Updated_at                   string
}

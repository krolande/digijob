package entities

import "database/sql"

type OffresGenerales struct {
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
}

// fusion des deux entites
type OffresEntreprises struct {
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

	IdE                   int
	LogoEntreprise        string
	TypeContratC          string
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
	//// Mission

	//Mission du Poste
	IdM                int
	MissionPoste       string
	OffresGenerales_id int
	TabMissionPoste    []MissionOffreGenerale
	CreatedAtM         string
	UpdateAtM          string

	//Profil Recherché

	IdP                 int
	ProfilPoste         string
	TabProfilRecherche  []ProfilRechercheOffreG
	OffresGenerales_idP int
	CreatedAtP          string
	UpdateAtP           string

	//Competence Clees
	IdCc        int
	Libelle     string
	TabLibelle  []CompetenceCles
	CreatedAtCc string
	UpdatedAtCc string
}

///// Dans la fusion des 2 entités celle qui sera ici sera appelé pour recuperer les valeurs

type OffresEntreprisesRec struct {
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

	IdE                   int
	LogoEntreprise        string
	TypeContratC          string
	NomEntreprise         string
	SecteurActivite       string
	Adresse               string
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
	Password              string
	CreatedAtE            string
	UpdateAtE             string

	//Mission du Poste
	IdM                int
	MissionPoste       string
	OffresGenerales_id int
	TabMissionPoste    []MissionOffreGenerale
	CreatedAtM         string
	UpdateAtM          string

	//Profil Recherché

	IdP                 int
	ProfilPoste         string
	TabProfilRecherche  []ProfilRechercheOffreG
	OffresGenerales_idP int
	CreatedAtP          string
	UpdateAtP           string

	//Competence Clees
	IdCc        int
	Libelle     string
	TabLibelle  []CompetenceCles
	CreatedAtCc string
	UpdatedAtCc string
}

// Mission entites
type MissionOffreGenerale struct {
	IdM                int
	MissionPoste       string
	OffresGenerales_id int
	CreatedAtM         string
	UpdateAtM          string
}

// Profil recherché

type ProfilRechercheOffreG struct {
	IdP                 int
	ProfilPoste         string
	OffresGenerales_idP int
	CreatedAtP          string
	UpdateAtP           string
}

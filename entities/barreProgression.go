package entities

type BarreProgressionCv struct {
	NbreFormation        int
	NbreExpeProf         int
	NbreComptProf        int
	NbrePrixCertificat   int
	NbreLangue           int
	UsersDigicodeuses_id int
	Created_at           string
	Updated_at           string
}

type BarreProgressionInfoPerso struct {
	Id            int
	NbreInfoPerso int
	Created_at    string
	Updated_at    string
}

type BarreProgressionCritere struct {
	Id          int
	NbreCritere int
	Created_at  string
	Updated_at  string
}

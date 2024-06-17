package entities

type CompetenceProfessionnelle struct {
	Id                   int
	Titre                string
	Pourcentage          int
	UsersDigicodeuses_id int
	Created_at           string
	Updated_at           string
}

type CompetenceProfessionnelles struct {
	Id                   []int
	Titre                []string
	Pourcentage          []int
	UsersDigicodeuses_id int
	Created_at           string
	Updated_at           string
}

package entities

type CompetenceCles struct {
	Id        int
	Libelle   string
	CreatedAt string
	UpdatedAt string
}

type CompetencesCles struct {
	Id        int
	Libelle   []string
	CreatedAt string
	UpdatedAt string
}

type CompetenceCleUsersDigicodeuses struct {
	Id                   int
	CompetenceCles_id    int
	UsersDigicodeuses_id int
}

// type AllCompetenceCleDigicodeuses struct {
// 	Id                   int
// 	CompetenceCles_id    int
// 	UsersDigicodeuses_id int
// }

package entities

type Langues struct {
	Id        int
	Libelle   string
	CreatedAt string
	UpdatedAt string
}

type UserLangues struct {
	Id        int
	Libelle   []string
	CreatedAt string
	UpdatedAt string
}

type LangueUsersDigicodeuses struct {
	Id                   int
	Langues_id           int
	UsersDigicodeuses_id int
	CreatedAt            string
	UpdatedAt            string
}

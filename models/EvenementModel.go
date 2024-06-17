package models

import (
	"digiJob/entities"
)

func (u UserModel) GetAllEvenements(evenement *entities.EvenementEntrepriseBd) (interface{}, error) {
	allEvenements := entities.EvenementEntreprises{}
	res := []entities.EvenementEntreprises{}
	row, err := u.db.Query("SELECT ev.*, en.* FROM evenements ev JOIN entreprises en ON en.id = ev.entreprises_id ORDER BY ev.id DESC")
	if err != nil {
		return res, err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&evenement.Id, &evenement.TitreEvent, &evenement.DescriptionEvent, &evenement.FileUpload, &evenement.DateEvent, &evenement.TypeEvent, &evenement.Heure, &evenement.FraisInscriptionEvent, &evenement.ModaliteEvent, &evenement.Lien, &evenement.LieuEvent, &evenement.DateLimite, &evenement.Entreprises_id, &evenement.CreatedAt, &evenement.UpdateAt, &evenement.IdE, &evenement.NomEntreprise, &evenement.SecteurActivite, &evenement.Adresse, &evenement.Logo, &evenement.SiteWeb, &evenement.DescriptionEntreprise, &evenement.NumeroEntreprise, &evenement.VideoEntreprise, &evenement.TailleEntreprise, &evenement.Facebook, &evenement.Twitter, &evenement.Linkedin, &evenement.Email, &evenement.Slogan, &evenement.Password, &evenement.CreatedAtE, &evenement.UpdateAtE)
		allEvenements.Id = evenement.Id
		allEvenements.TitreEvent = evenement.TitreEvent
		allEvenements.DescriptionEvent = evenement.DescriptionEvent
		allEvenements.FileUpload = evenement.FileUpload
		allEvenements.DateEvent = evenement.DateEvent
		allEvenements.FraisInscriptionEvent = evenement.FraisInscriptionEvent.String
		allEvenements.LieuEvent = evenement.LieuEvent
		allEvenements.ModaliteEvent = evenement.ModaliteEvent.String
		allEvenements.Lien = evenement.Lien.String
		allEvenements.Heure = evenement.Heure.String

		//Entreprises
		allEvenements.IdE = evenement.IdE
		allEvenements.NomEntreprise = evenement.NomEntreprise
		allEvenements.Logo = evenement.Logo
		res = append(res, allEvenements)
	}

	return res, nil
}

func (u UserModel) DetailEvenements(evenement *entities.EvenementEntrepriseBd, id int) (interface{}, error) {
	allEvenements := entities.EvenementEntreprises{}
	row, err := u.db.Query("SELECT ev.*, en.* FROM evenements ev JOIN entreprises en ON en.id = ev.entreprises_id WHERE ev.id = ?", id)
	if err != nil {
		return row, err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&evenement.Id, &evenement.TitreEvent, &evenement.DescriptionEvent, &evenement.FileUpload, &evenement.DateEvent, &evenement.TypeEvent, &evenement.Heure, &evenement.FraisInscriptionEvent, &evenement.ModaliteEvent, &evenement.Lien, &evenement.LieuEvent, &evenement.DateLimite, &evenement.Entreprises_id, &evenement.CreatedAt, &evenement.UpdateAt, &evenement.IdE, &evenement.NomEntreprise, &evenement.SecteurActivite, &evenement.Adresse, &evenement.Logo, &evenement.SiteWeb, &evenement.DescriptionEntreprise, &evenement.NumeroEntreprise, &evenement.VideoEntreprise, &evenement.TailleEntreprise, &evenement.Facebook, &evenement.Twitter, &evenement.Linkedin, &evenement.Email, &evenement.Slogan, &evenement.Password, &evenement.CreatedAtE, &evenement.UpdateAtE)
		allEvenements.Id = evenement.Id
		allEvenements.TitreEvent = evenement.TitreEvent
		allEvenements.DescriptionEvent = evenement.DescriptionEvent
		allEvenements.FileUpload = evenement.FileUpload
		allEvenements.DateEvent = evenement.DateEvent
		allEvenements.FraisInscriptionEvent = evenement.FraisInscriptionEvent.String
		allEvenements.ModaliteEvent = evenement.ModaliteEvent.String
		allEvenements.LieuEvent = evenement.LieuEvent
		allEvenements.Heure = evenement.Heure.String
		allEvenements.Lien = evenement.Lien.String

		//Entreprises
		allEvenements.IdE = evenement.IdE
		allEvenements.NomEntreprise = evenement.NomEntreprise
		allEvenements.Logo = evenement.Logo
	}
	return allEvenements, nil
}

//Renvoie les évènements d'une entreprise
func (u UserModel) EvenementEntreprises(evenement *entities.EvenementBd, id int) (interface{}, error) {
	row, err := u.db.Query("SELECT * FROM evenements WHERE entreprises_id = ?", id)
	if err != nil {
		return row, err
	}
	defer row.Close()
	allEvenements := entities.Evenements{}
	res := []entities.Evenements{}
	for row.Next() {
		row.Scan(&evenement.Id, &evenement.TitreEvent, &evenement.DescriptionEvent, &evenement.FileUpload, &evenement.DateEvent, &evenement.TypeEvent, &evenement.Heure, &evenement.FraisInscriptionEvent, &evenement.ModaliteEvent, &evenement.Lien, &evenement.LieuEvent, &evenement.DateLimite, &evenement.Entreprises_id, &evenement.Created_At, &evenement.Update_At)
		allEvenements.Id = evenement.Id
		allEvenements.TitreEvent = evenement.TitreEvent
		allEvenements.DescriptionEvent = evenement.DescriptionEvent
		allEvenements.FileUpload = evenement.FileUpload
		allEvenements.DateEvent = evenement.DateEvent
		res = append(res, allEvenements)
	}

	return res, nil
}

//Insertion des participants a un evenement
func (u UserModel) InsParticipantEvent(participant entities.ParticipantEvent) (int64, error) {
	insForm, err := u.db.Exec("INSERT INTO `participantEvenements`(`evenements_id`, `usersDigicodeuses_id`) VALUES (?,?)", participant.Evenements_id, participant.UsersDigicodeuses_id)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := insForm.LastInsertId()

	return lastInsertId, nil
}

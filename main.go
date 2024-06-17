package main

import (
	"fmt"
	"net/http"

	authcontroller "digiJob/controllers"
	recruteurscontroller "digiJob/controllers/recruteursController"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/acceuil", authcontroller.Acceuil)
	//********** PARAMETRE DE CONNEXION ************//
	http.HandleFunc("/loginDigicodeuses", authcontroller.Login)
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/modifPassword", authcontroller.ModifPassword)
	http.HandleFunc("/Updateprofile", authcontroller.UpdateProfile)
	http.HandleFunc("/profile", authcontroller.ViewsProfil)
	http.HandleFunc("/monCv", authcontroller.MonCV)
	http.HandleFunc("/deleteFormationCv", authcontroller.DeleteFormationCv)
	http.HandleFunc("/deleteExperienceCv", authcontroller.DeleteExperienceCv)
	http.HandleFunc("/deleteExperiencePro", authcontroller.DeleteCompetenceProfessionnelle)
	http.HandleFunc("/deletePrixCertification", authcontroller.DeletePrixCertifications)
	http.HandleFunc("/critereDigicodeuse", authcontroller.CritereDigicodeuse)
	http.HandleFunc("/apercuCv", authcontroller.ApercuCv)

	//************ Evenement
	http.HandleFunc("/allEvenement", authcontroller.AllEvenement)
	http.HandleFunc("/detailEvenements", authcontroller.DetailEvenements)

	//**************** Entreprise
	http.HandleFunc("/allEntreprises", authcontroller.AllEntreprises)
	http.HandleFunc("/detailEntreprises", authcontroller.DetailEntreprises)

	//*************Formulaire connexion recruteur *************//
	//http.HandleFunc("/connexionRecruteur", authcontroller.ContactRecruteur)
	// ******************* Login recruteur ***************//
	http.HandleFunc("/loginRecruteurs", recruteurscontroller.Login)
	http.HandleFunc("/logoutRecruteurs", recruteurscontroller.LogoutRecruteurs)
	//dashboard recruteur après connexion
	http.HandleFunc("/dashboardRecruteur", recruteurscontroller.DashboardRecruteur)

	//********** OFFRES GENERALE ***********//
	http.HandleFunc("/allOffres", authcontroller.Offres)
	http.HandleFunc("/detailOffres", authcontroller.DetailOffres)

	//************* CANDIDATURE OFFRE GENERALE*************//
	http.HandleFunc("/candidatureOffreGenerale", authcontroller.InsCandidature)
	http.HandleFunc("/allCandidature", authcontroller.AllCandidatureDigi)
	http.HandleFunc("/deleteCandidature", authcontroller.DeleteCandidature)

	//************** File **********************
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/files/", http.StripPrefix("/files", fs))
	http.Handle("/views/assets/", http.StripPrefix("/views/assets/", http.FileServer(http.Dir("./views/assets/"))))
	fmt.Println("http://localhost:12000")
	http.ListenAndServe(":12000", nil)
}

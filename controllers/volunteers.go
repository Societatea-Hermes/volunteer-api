package controllers

import (
	"hermes-api/helpers"
	"hermes-api/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var models services.Models
var volunteer = models.Volunteer

func GetAllVolunteers(w http.ResponseWriter, r *http.Request) {
	var volunteers services.Volunteer
	all, err := volunteers.GetAllVolunteers()
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, all)
}

func CreateVolunteer(w http.ResponseWriter, r *http.Request) {
	var volunteerResp services.Volunteer

	// err := json.NewDecoder(r.Body).Decode(&volunteerResp)
	err := helpers.ReadJSON(w, r, &volunteerResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	createdVolunteer, err := volunteer.CreateVolunteer(volunteerResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, createdVolunteer)
}

func GetVolunteerByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	volunteer, err := volunteer.GetVolunteerByEmail(email)
	if err != nil {
		log.Println("EROARE DOAMNE")
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, volunteer)
}

func ActivateVolunteer(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	volunteer, err := volunteer.UpdateVolunteerActive(email, true)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, volunteer)
}

func DeactivateVolunteer(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	volunteer, err := volunteer.UpdateVolunteerActive(email, false)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, volunteer)
}

func UpdateVolunteerPersonalInformation(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	var volunteerResp services.Volunteer
	err := helpers.ReadJSON(w, r, &volunteerResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	updatedVolunteer, err := volunteer.UpdatePersonalInfo(email, volunteerResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, updatedVolunteer)
}

func ChangeDepartmentVolunteer(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	department := chi.URLParam(r, "department")
	volunteer, err := volunteer.ChangeDepartment(email, department)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, volunteer)
}

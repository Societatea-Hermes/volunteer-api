package controllers

import (
	"hermes-api/helpers"
	"hermes-api/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var candidate = models.Candidate

func GetAllCandidates(w http.ResponseWriter, r *http.Request) {
	recruitmentCampaignId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)

		return
	}
	all, err := candidate.GetAllCandidates(recruitmentCampaignId)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, all)
}

func CreateCandidate(w http.ResponseWriter, r *http.Request) {
	var candidateResp services.Candidate

	err := helpers.ReadJSON(w, r, &candidateResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	createdCandidate, err := candidate.CreateCandidate(candidateResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, createdCandidate)
}

func UpdateCandidateStatus(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "personal_email")
	status := chi.URLParam(r, "status")
	candidate, err := candidate.UpdateRecruitmentStatus(email, status)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, candidate)
}

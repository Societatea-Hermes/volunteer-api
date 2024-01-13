package controllers

import (
	"net/http"
	"strconv"
	"volunteer-api/helpers"
	"volunteer-api/models"

	"github.com/go-chi/chi/v5"
)

func GetAllCandidates(w http.ResponseWriter, r *http.Request) {
	recruitmentCampaignId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)

		return
	}
	all, err := candidateModel.GetAllCandidates(recruitmentCampaignId)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, all)
}

func CreateCandidate(w http.ResponseWriter, r *http.Request) {
	var candidateResp models.Candidate

	err := helpers.ReadJSON(w, r, &candidateResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	createdCandidate, err := candidateModel.CreateCandidate(&candidateResp)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, createdCandidate)
}

func UpdateCandidateStatus(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "personal_email")
	status := chi.URLParam(r, "status")
	candidate, err := candidateModel.UpdateRecruitmentStatus(email, status)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, candidate)
}

package controllers

import (
	"hermes-api/helpers"
	"hermes-api/services"
	"net/http"
)

var recruitmentCampaign = models.RecruitmentCampaign

func CreateRecruitmentCampaign(w http.ResponseWriter, r *http.Request) {
	var campaignResponse services.RecruitmentCampaign

	err := helpers.ReadJSON(w, r, &campaignResponse)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	campaign, err := recruitmentCampaign.CreateRecruitmentCampaign(campaignResponse)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, campaign)
}

package controllers

import (
	"volunteer-api/helpers"
	"volunteer-api/models"
	"net/http"
)

var recruitmentCampaign = models.RecruitmentCampaign

func CreateRecruitmentCampaign(w http.ResponseWriter, r *http.Request) {
	var campaignResponse models.RecruitmentCampaign

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

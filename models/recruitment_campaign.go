package models

import (
	"time"
)

type RecruitmentCampaign struct {
	gorm.Model
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (r *RecruitmentCampaign) CreateRecruitmentCampaign(campaign* RecruitmentCampaign) (*RecruitmentCampaign, error) {
	result := db.Create(&campaign)
	if result.Error != nil {
		return nil, result.Error
	}
	return campaign, nil
}

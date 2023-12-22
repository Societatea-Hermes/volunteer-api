package models

import (
	"time"

	"gorm.io/gorm"
)

type RecruitmentCampaign struct {
	gorm.Model
	Name       string      `json:"name"`
	StartDate  time.Time   `json:"start_date"`
	EndDate    time.Time   `json:"end_date"`
	Candidates []Candidate `gorm:"foreignKey:RecruitmentCampaignID"`
	Volunteers []Volunteer `gorm:"foreignKey:RecruitmentCampaignID"`
}

func (r *RecruitmentCampaign) CreateRecruitmentCampaign(campaign *RecruitmentCampaign) (*RecruitmentCampaign, error) {
	result := db.Create(&campaign)
	if result.Error != nil {
		return nil, result.Error
	}
	return campaign, nil
}

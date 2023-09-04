package services

import (
	"time"
)

type RecruitmentCampaign struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (r *RecruitmentCampaign) CreateRecruitmentCampaign(campaign RecruitmentCampaign) (*RecruitmentCampaign, error) {
	query := `INSERT INTO recruitment_campaigns (name, start_date, end_date)
	VALUES ($1, $2, $3)
	RETURNING id, name, start_date, end_date
	`

	err := db.QueryRow(query, campaign.Name, campaign.StartDate, campaign.EndDate).Scan(
		&campaign.Id, &campaign.Name, &campaign.StartDate, &campaign.EndDate,
	)
	if err != nil {
		return nil, err
	}

	return &campaign, nil
}

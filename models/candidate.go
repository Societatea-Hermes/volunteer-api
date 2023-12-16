package models

import (
	"context"
	"time"
)

type Candidate struct {
	gorm.Model
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name"`
	PersonalEmail         string    `json:"personal_email"`
	Phone                 string    `json:"phone"`
	Address               string    `json:"address"`
	BirthDate             time.Time `json:"birth_date"`
	Gender                string    `json:"gender"`
	StudiesType           string    `json:"students_type"`
	Specialization        string    `json:"specialization"`
	StudyGroup            string    `json:"study_group"`
	StudyLanguage         string    `json:"study_language"`
	FacebookProfile       string    `json:"facebook_profile"`
	InstagramProfile      string    `json:"instagram_profile"`
	RecruitmentStatus     string    `json:"recruitment_status"`
	RecruitmentCampaignId int64     `json:"recruitment_campaign_id"`
}

func (c *Candidate) GetAllCandidates(campaign_id int64) ([]*Candidate, error) {
	var volunteers []Volunteer
	result := db.Find(&volunteers)
	if result.Error != nil {
		return nil, result.Error
	}
	return volunteers, nil
}

func (v *Candidate) CreateCandidate(candidate* Candidate) (*Candidate, error) {
	result := db.Create(&candidate)
	if result.Error != nil {
		return nil, result.Error
	}
	return candidate, nil
}

func (c *Candidate) UpdateRecruitmentStatus(personal_email string, status string) (*Candidate, error) {
	var candidate Candidate
	result := db.Where("personal_email = ?", email).First(&candidate)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("candidate not found")
	} else if result.Error != nil {
		return nil, result.Error
	}

	candidate.Status = status
	result = db.Save(&candidate)
	if result.Error != nil {
		return nil, result.Error
	}

	return &candidate, nil
}

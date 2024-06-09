package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Candidate struct {
	gorm.Model
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name"`
	PersonalEmail         string    `json:"personal_email" validate:"required,email"`
	Phone                 string    `json:"phone" validate:"required,phone"`
	Address               string    `json:"address"`
	BirthDate             time.Time `json:"birth_date"`
	Gender                string    `json:"gender"`
	StudiesType           string    `json:"studies_type"`
	Specialization        string    `json:"specialization"`
	StudyGroup            string    `json:"study_group"`
	StudyLanguage         string    `json:"study_language"`
	FacebookProfile       string    `json:"facebook_profile"`
	InstagramProfile      string    `json:"instagram_profile"`
	RecruitmentStatus     string    `json:"recruitment_status" gorm:"default:'Pending'" `
	RecruitmentCampaignID uint      `json:"recruitment_campaign_id"`
}

func (c *Candidate) GetAllCandidates() ([]Candidate, error) {
	var candidates []Candidate
	result := db.Find(&candidates)
	if result.Error != nil {
		return nil, result.Error
	}
	return candidates, nil
}

func (v *Candidate) CreateCandidate(candidate *Candidate) (*Candidate, error) {
	result := db.Create(&candidate)
	if result.Error != nil {
		return nil, result.Error
	}
	return candidate, nil
}

func (c *Candidate) UpdateRecruitmentStatus(personal_email string, status string) (*Candidate, error) {
	var candidate Candidate
	result := db.Where("personal_email = ?", personal_email).First(&candidate)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("candidate not found")
	} else if result.Error != nil {
		return nil, result.Error
	}

	candidate.RecruitmentStatus = status
	result = db.Save(&candidate)
	if result.Error != nil {
		return nil, result.Error
	}

	return &candidate, nil
}

func (c *Candidate) GetAllCandidatesByCampaign(id int64) ([]Candidate, error) {
	var candidates []Candidate
	result := db.Where("recruitment_campaign_id = ?", id).Find(&candidates)
	if result.Error != nil {
		return nil, result.Error
	}
	return candidates, nil
}

func (c *Candidate) UpdateCandidate(body Candidate) (*Candidate, error) {
	var existingCandidate Candidate
	result := db.Where("id = ?", body.ID).First(&existingCandidate)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("volunteer not found")
	} else if result.Error != nil {
		return nil, result.Error
	}

	existingCandidate.FirstName = body.FirstName
	existingCandidate.LastName = body.LastName
	existingCandidate.PersonalEmail = body.PersonalEmail
	existingCandidate.Phone = body.Phone
	existingCandidate.Address = body.Address
	existingCandidate.BirthDate = body.BirthDate
	existingCandidate.Gender = body.Gender
	existingCandidate.StudiesType = body.StudiesType
	existingCandidate.Specialization = body.Specialization
	existingCandidate.StudyGroup = body.StudyGroup
	existingCandidate.StudyLanguage = body.StudyLanguage
	existingCandidate.FacebookProfile = body.FacebookProfile
	existingCandidate.InstagramProfile = body.InstagramProfile
	existingCandidate.RecruitmentStatus = body.RecruitmentStatus
	existingCandidate.RecruitmentCampaignID = body.RecruitmentCampaignID

	result = db.Save(&existingCandidate)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingCandidate, nil
}

func (c *Candidate) DeleteCandidate(email string) error {
	var candidate Candidate
	err := db.Where("personal_email = ?", email).First(&candidate).Error
	if err != nil {
		return err
	}

	err = db.Delete(&candidate).Error
	if err != nil {
		return err
	}

	return nil
}

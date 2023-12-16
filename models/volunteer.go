package models

import (
	"context"
	"log"
	"time"
)

type Volunteer struct {
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
	Email                 string    `json:"email"`
	Active                bool      `json:"active"`
	Department            string    `json:"department"`
	AuxMember             bool      `json:"aux_member"`
	Vegetarian            bool      `json:"vegetarian"`
	ShirtSize             string    `json:"shirt_size"`
	RecruitmentCampaignId int64     `json:"recruitment_campaign_id"`
}

func (v *Volunteer) GetAllVolunteers() ([]Volunteer, error) {
	var volunteers []Volunteer
	result := db.Find(&volunteers)
	if result.Error != nil {
		return nil, result.Error
	}
	return volunteers, nil
}

func (v *Volunteer) CreateVolunteer(volunteer* Volunteer) (*Volunteer, error) {
	result := db.Create(&volunteer)
	if result.Error != nil {
		return nil, result.Error
	}
	return volunteer, nil
}

func (v *Volunteer) GetVolunteerByEmail(email string) (*Volunteer, error) {
	var volunteer Volunteer
	result := db.First(&volunteer, email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &volunteer, nil
}

func (v *Volunteer) UpdateVolunteerActive(email string, active bool) (*Volunteer, error) {
	var volunteer Volunteer
	result := db.Where("email = ?", email).First(&volunteer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("volunteer not found")
	} else if result.Error != nil {
		return nil, result.Error
	}

	volunteer.Active = active
	result = db.Save(&volunteer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &volunteer, nil
}

func (v *Volunteer) UpdatePersonalInfo(email string, body Volunteer) (*Volunteer, error) {
	var existingVolunteer Volunteer
	result := db.Where("email = ?", email).First(&existingVolunteer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("volunteer not found")
	} else if result.Error != nil {
		return nil, result.Error
	}

	existingVolunteer.FirstName = body.FirstName
	existingVolunteer.LastName = body.LastName
	existingVolunteer.PersonalEmail = body.PersonalEmail
	existingVolunteer.Phone = body.Phone
	existingVolunteer.Address = body.Address
	existingVolunteer.BirthDate = body.BirthDate
	existingVolunteer.Gender = body.Gender
	existingVolunteer.StudiesType = body.StudiesType
	existingVolunteer.Specialization = body.Specialization
	existingVolunteer.StudyGroup = body.StudyGroup
	existingVolunteer.StudyLanguage = body.StudyLanguage
	existingVolunteer.FacebookProfile = body.FacebookProfile
	existingVolunteer.InstagramProfile = body.InstagramProfile
	existingVolunteer.Email = body.Email
	existingVolunteer.Active = body.Active
	existingVolunteer.Department = body.Department
	existingVolunteer.AuxMember = body.AuxMember
	existingVolunteer.Vegetarian = body.Vegetarian
	existingVolunteer.ShirtSize = body.ShirtSize
	existingVolunteer.RecruitmentCampaignId = body.RecruitmentCampaignId

	result = db.Save(&existingVolunteer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingVolunteer, nil
}

func (v *Volunteer) ChangeDepartment(email string, newDepartment string) (*Volunteer, error) {
	var volunteer Volunteer
	result := db.Where("email = ?", email).First(&volunteer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("volunteer not found")
	} else if result.Error != nil {
		return nil, result.Error
	}

	volunteer.Department = newDepartment
	result = db.Save(&volunteer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &volunteer, nil
}

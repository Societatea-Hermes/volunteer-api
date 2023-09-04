package services

import (
	"context"
	"time"
)

type Candidate struct {
	Id                    int64     `json:"id"`
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
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	RecruitmentCampaignId int64     `json:"recruitment_campaign_id"`
}

func (c *Candidate) GetAllCandidates(campaign_id int64) ([]*Candidate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, first_name, last_name, personal_email, phone, address, birth_date,
	gender, studies_type, specialization, study_group, study_language,
	facebook_profile, instagram_profile, recruitment_status, created_at,
	updated_at, recruitment_campaign_id  FROM candidates`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var candidates []*Candidate
	for rows.Next() {
		var candidate Candidate
		err := rows.Scan(
			&candidate.Id,
			&candidate.FirstName,
			&candidate.LastName,
			&candidate.PersonalEmail,
			&candidate.Phone,
			&candidate.Address,
			&candidate.BirthDate,
			&candidate.Gender,
			&candidate.StudiesType,
			&candidate.Specialization,
			&candidate.StudyGroup,
			&candidate.StudyLanguage,
			&candidate.FacebookProfile,
			&candidate.InstagramProfile,
			&candidate.RecruitmentStatus,
			&candidate.CreatedAt,
			&candidate.UpdatedAt,
			&candidate.RecruitmentCampaignId,
		)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, &candidate)
	}

	return candidates, nil
}

func (v *Candidate) CreateCandidate(candidate Candidate) (*Candidate, error) {
	query := `INSERT INTO candidates (
        first_name,
        last_name,
        personal_email,
        phone,
		address,
        birth_date,
        gender,
        studies_type,
        specialization,
		study_group,
        study_language,
        facebook_profile,
        instagram_profile,
		recruitment_status,
		created_at,
        updated_at,
		recruitment_campaign_id
		) values 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		returning id, created_at, updated_at, recruitment_status
		`
	var insertedId int64
	var createdAt time.Time
	var updatedAt time.Time
	var recruitmentStatus string

	err := db.QueryRow(
		query,
		candidate.FirstName,
		candidate.LastName,
		candidate.PersonalEmail,
		candidate.Phone,
		candidate.Address,
		candidate.BirthDate,
		candidate.Gender,
		candidate.StudiesType,
		candidate.Specialization,
		candidate.StudyGroup,
		candidate.StudyLanguage,
		candidate.FacebookProfile,
		candidate.InstagramProfile,
		"pending",
		time.Now(),
		time.Now(),
		candidate.RecruitmentCampaignId,
	).Scan(&insertedId, &createdAt, &updatedAt, &recruitmentStatus)
	if err != nil {
		return nil, err
	}

	candidate.Id = insertedId
	candidate.RecruitmentStatus = recruitmentStatus
	candidate.CreatedAt = createdAt
	candidate.UpdatedAt = updatedAt
	return &candidate, nil
}

func (c *Candidate) UpdateRecruitmentStatus(personal_email string, status string) (*Candidate, error) {
	query := `UPDATE candidates
	SET
		recruitment_status = $1
	WHERE personal_email = $2
	RETURNING id, first_name, last_name, personal_email, phone, address, birth_date, gender,
	studies_type, specialization, study_group, study_language, facebook_profile, instagram_profile,
	recruitment_status, created_at, updated_at, recruitment_campaign_id
	`

	var candidate Candidate
	err := db.QueryRow(query, status, personal_email).Scan(
		&candidate.Id,
		&candidate.FirstName,
		&candidate.LastName,
		&candidate.PersonalEmail,
		&candidate.Phone,
		&candidate.Address,
		&candidate.BirthDate,
		&candidate.Gender,
		&candidate.StudiesType,
		&candidate.Specialization,
		&candidate.StudyGroup,
		&candidate.StudyLanguage,
		&candidate.FacebookProfile,
		&candidate.InstagramProfile,
		&candidate.RecruitmentStatus,
		&candidate.CreatedAt,
		&candidate.UpdatedAt,
		&candidate.RecruitmentCampaignId,
	)

	if err != nil {
		return nil, err
	}
	return &candidate, nil
}

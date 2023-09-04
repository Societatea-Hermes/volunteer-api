package services

import (
	"context"
	"log"
	"time"
)

type Volunteer struct {
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
	Email                 string    `json:"email"`
	Active                bool      `json:"active"`
	Department            string    `json:"department"`
	AuxMember             bool      `json:"aux_member"`
	Vegetarian            bool      `json:"vegetarian"`
	ShirtSize             string    `json:"shirt_size"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	RecruitmentCampaignId int64     `json:"recruitment_campaign_id"`
}

func (v *Volunteer) GetAllVolunteers() ([]*Volunteer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, first_name, last_name, personal_email, phone, address, birth_date,
	gender, studies_type, specialization, study_group, study_language,
	facebook_profile, instagram_profile, email, active, department,
	aux_member, vegetarian, shirt_size, created_at, updated_at, recruitment_campaign_id  FROM volunteers`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var volunteers []*Volunteer
	for rows.Next() {
		var volunteer Volunteer
		err := rows.Scan(
			&volunteer.Id,
			&volunteer.FirstName,
			&volunteer.LastName,
			&volunteer.PersonalEmail,
			&volunteer.Phone,
			&volunteer.Address,
			&volunteer.BirthDate,
			&volunteer.Gender,
			&volunteer.StudiesType,
			&volunteer.Specialization,
			&volunteer.StudyGroup,
			&volunteer.StudyLanguage,
			&volunteer.FacebookProfile,
			&volunteer.InstagramProfile,
			&volunteer.Email,
			&volunteer.Active,
			&volunteer.Department,
			&volunteer.AuxMember,
			&volunteer.Vegetarian,
			&volunteer.ShirtSize,
			&volunteer.CreatedAt,
			&volunteer.UpdatedAt,
			&volunteer.RecruitmentCampaignId,
		)
		if err != nil {
			return nil, err
		}
		log.Printf("%v %s\n", volunteer.Id, volunteer.FirstName)
		volunteers = append(volunteers, &volunteer)
	}

	return volunteers, nil
}

func (v *Volunteer) CreateVolunteer(volunteer Volunteer) (*Volunteer, error) {
	query := `INSERT INTO volunteers (
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
        email,
		active,
        department,
        aux_member,
        vegetarian,
        shirt_size,
		created_at,
        updated_at,
		recruitment_campaign_id
		) values 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		returning id, created_at, updated_at
		`
	var insertedId int64
	var createdAt time.Time
	var updatedAt time.Time

	err := db.QueryRow(
		query,
		volunteer.FirstName,
		volunteer.LastName,
		volunteer.PersonalEmail,
		volunteer.Phone,
		volunteer.Address,
		volunteer.BirthDate,
		volunteer.Gender,
		volunteer.StudiesType,
		volunteer.Specialization,
		volunteer.StudyGroup,
		volunteer.StudyLanguage,
		volunteer.FacebookProfile,
		volunteer.InstagramProfile,
		volunteer.Email,
		volunteer.Active,
		volunteer.Department,
		volunteer.AuxMember,
		volunteer.Vegetarian,
		volunteer.ShirtSize,
		time.Now(),
		time.Now(),
		volunteer.RecruitmentCampaignId,
	).Scan(&insertedId, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	volunteer.Id = insertedId
	volunteer.CreatedAt = createdAt
	volunteer.UpdatedAt = updatedAt
	return &volunteer, nil
}

func (v *Volunteer) GetVolunteerByEmail(email string) (*Volunteer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, first_name, last_name, personal_email, phone, address, birth_date,
	gender, studies_type, specialization, study_group, study_language,
	facebook_profile, instagram_profile, email, active, department,
	aux_member, vegetarian, shirt_size, created_at, updated_at, recruitment_campaign_id FROM volunteers WHERE email = $1`
	var volunteer Volunteer

	row := db.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&volunteer.Id,
		&volunteer.FirstName,
		&volunteer.LastName,
		&volunteer.PersonalEmail,
		&volunteer.Phone,
		&volunteer.Address,
		&volunteer.BirthDate,
		&volunteer.Gender,
		&volunteer.StudiesType,
		&volunteer.Specialization,
		&volunteer.StudyGroup,
		&volunteer.StudyLanguage,
		&volunteer.FacebookProfile,
		&volunteer.InstagramProfile,
		&volunteer.Email,
		&volunteer.Active,
		&volunteer.Department,
		&volunteer.AuxMember,
		&volunteer.Vegetarian,
		&volunteer.ShirtSize,
		&volunteer.CreatedAt,
		&volunteer.UpdatedAt,
		&volunteer.RecruitmentCampaignId,
	)
	if err != nil {
		return nil, err
	}

	return &volunteer, nil
}

func (v *Volunteer) UpdateVolunteerActive(email string, active bool) (*Volunteer, error) {
	query := `UPDATE volunteers
	SET
		active = $1
	WHERE email = $2
	RETURNING id, first_name, last_name, personal_email, phone, address, birth_date, gender,
	studies_type, specialization, study_group, study_language, facebook_profile, instagram_profile,
	email, department, active, aux_member, vegetarian, shirt_size, created_at, updated_at, recruitment_campaign_id
	`

	var volunteer Volunteer
	err := db.QueryRow(query, active, email).Scan(
		&volunteer.Id,
		&volunteer.FirstName,
		&volunteer.LastName,
		&volunteer.PersonalEmail,
		&volunteer.Phone,
		&volunteer.Address,
		&volunteer.BirthDate,
		&volunteer.Gender,
		&volunteer.StudiesType,
		&volunteer.Specialization,
		&volunteer.StudyGroup,
		&volunteer.StudyLanguage,
		&volunteer.FacebookProfile,
		&volunteer.InstagramProfile,
		&volunteer.Email,
		&volunteer.Department,
		&volunteer.Active,
		&volunteer.AuxMember,
		&volunteer.Vegetarian,
		&volunteer.ShirtSize,
		&volunteer.CreatedAt,
		&volunteer.UpdatedAt,
		&volunteer.RecruitmentCampaignId,
	)

	if err != nil {
		return nil, err
	}
	return &volunteer, nil
}

func (v *Volunteer) UpdatePersonalInfo(email string, body Volunteer) (*Volunteer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE volunteers
	SET
		first_name = $1,
		last_name = $2,
		personal_email = $3,
		phone = $4,
		address = $5,
		gender = $6,
		studies_type = $7,
		specialization = $8,
		study_group = $9,
		study_language = $10,
		facebook_profile = $11,
		instagram_profile = $12,
		vegetarian = $13,
		shirt_size = $14,
		updated_at = $15
	WHERE email = $16
	`
	body.UpdatedAt = time.Now()

	_, err := db.ExecContext(ctx, query,
		body.FirstName,
		body.LastName,
		body.PersonalEmail,
		body.Phone,
		body.Address,
		body.Gender,
		body.StudiesType,
		body.Specialization,
		body.StudyGroup,
		body.StudyLanguage,
		body.FacebookProfile,
		body.InstagramProfile,
		body.Vegetarian,
		body.ShirtSize,
		body.UpdatedAt,
		email,
	)

	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (v *Volunteer) ChangeDepartment(email string, newDepartment string) (*Volunteer, error) {
	var volunteer Volunteer
	query := `UPDATE volunteers
	SET
		department = $1
	WHERE email = $2
	RETURNING id, first_name, last_name, personal_email, phone, address, birth_date, gender,
	studies_type, specialization, study_group, study_language, facebook_profile, instagram_profile,
	email, department, active, aux_member, vegetarian, shirt_size, created_at, updated_at, recruitment_campaign_id
	`
	err := db.QueryRow(query, newDepartment, email).Scan(
		&volunteer.Id,
		&volunteer.FirstName,
		&volunteer.LastName,
		&volunteer.PersonalEmail,
		&volunteer.Phone,
		&volunteer.Address,
		&volunteer.BirthDate,
		&volunteer.Gender,
		&volunteer.StudiesType,
		&volunteer.Specialization,
		&volunteer.StudyGroup,
		&volunteer.StudyLanguage,
		&volunteer.FacebookProfile,
		&volunteer.InstagramProfile,
		&volunteer.Email,
		&volunteer.Department,
		&volunteer.Active,
		&volunteer.AuxMember,
		&volunteer.Vegetarian,
		&volunteer.ShirtSize,
		&volunteer.CreatedAt,
		&volunteer.UpdatedAt,
		&volunteer.RecruitmentCampaignId,
	)

	if err != nil {
		return nil, err
	}
	return &volunteer, nil
}

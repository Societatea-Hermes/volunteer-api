package models

import (
	"database/sql"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}

type Models struct {
	Volunteer           Volunteer
	RecruitmentCampaign RecruitmentCampaign
	Candidate           Candidate
}

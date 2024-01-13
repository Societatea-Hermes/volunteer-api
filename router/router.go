package router

import (
	"fmt"
	"net/http"
	"volunteer-api/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const baseApiURL = "/api/v1"

func Routes() http.Handler {
	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Heartbeat("/ping"))

	// VOLUNTEER ROUTES
	volunteerBaseString := fmt.Sprintf("%s/volunteers", baseApiURL)
	router.Get(volunteerBaseString, controllers.GetAllVolunteers)
	router.Post(volunteerBaseString, controllers.CreateVolunteer)

	getByEmailURL := fmt.Sprintf("%s/volunteer/{email}", volunteerBaseString)
	router.Get(getByEmailURL, controllers.GetVolunteerByEmail)

	activateURL := fmt.Sprintf("%s/activate", getByEmailURL)
	router.Patch(activateURL, controllers.ActivateVolunteer)

	deactivateURL := fmt.Sprintf("%s/deactivate", getByEmailURL)
	router.Patch(deactivateURL, controllers.DeactivateVolunteer)

	updatePersonalInfoURL := fmt.Sprintf("%s/personal", getByEmailURL)
	router.Put(updatePersonalInfoURL, controllers.UpdateVolunteerPersonalInformation)

	changeDepartmentURL := fmt.Sprintf("%s/changeDepartment/{department}", getByEmailURL)
	router.Put(changeDepartmentURL, controllers.ChangeDepartmentVolunteer)

	// RECRUITMENT CAMPAIGNS ROUTES
	recruitmentCampaignsBasicURL := fmt.Sprintf("%s/recruitments", baseApiURL)
	router.Post(recruitmentCampaignsBasicURL, controllers.CreateRecruitmentCampaign)

	// RECRUITMENT CANDIDATES ROUTES
	candidateBasicURL := fmt.Sprintf("%s/candidates", recruitmentCampaignsBasicURL)
	router.Get(candidateBasicURL, controllers.GetAllCandidates)
	router.Put(candidateBasicURL, controllers.CreateCandidate)

	updateCandidateStatusURL := fmt.Sprintf("%s/candidate/{personal_email}/{status}", candidateBasicURL)
	router.Put(updateCandidateStatusURL, controllers.UpdateCandidateStatus)
	return router
}

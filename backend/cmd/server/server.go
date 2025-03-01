package cmdserver

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	httpadapter "github.com/rafaelpissolatto/formUnity/backend/internal/adapter/api/http"
	slackadapter "github.com/rafaelpissolatto/formUnity/backend/internal/adapter/notification/slack"
	postgresadapter "github.com/rafaelpissolatto/formUnity/backend/internal/adapter/repository/postgres"
	"github.com/rafaelpissolatto/formUnity/backend/internal/domain/volunteer"
	"github.com/slack-go/slack"

	_ "github.com/lib/pq"
)

func RunAPIServer(_ context.Context) {
	log.Println("running api server...")

	// gorm / db
	dsn := os.Getenv("DATABASE_DNS")
	dbClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	volunteerRepo := postgresadapter.NewPostgresVolunteerRepository(dbClient)
	volunteerRepo.VolunteerMigration()

	// slack notifier
	slackClient := slack.New(os.Getenv("SLACK_TOKEN"))
	notifier := slackadapter.NewSlackNotifierWithOptions(
		slackadapter.WithToken(slackClient),
	)

	// volunteer service
	volunteerService := volunteer.NewVolunteerService(
		volunteerRepo,
		notifier,
	)

	// http
	volunteerHandler := httpadapter.NewVolunteerHandler(volunteerService)
	r := mux.NewRouter()
	r.HandleFunc("/volunteers", volunteerHandler.AddVolunteer).Methods("POST")
	r.HandleFunc("/volunteers/{id}", volunteerHandler.GetVolunteer).Methods("GET")
	r.HandleFunc("/volunteers", volunteerHandler.GetAllVolunteers).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}

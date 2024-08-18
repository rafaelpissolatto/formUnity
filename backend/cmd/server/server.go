package cmdserver

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	httpadapter "github.com/rafaelpissolatto/formUnity/backend/internal/adapter/api/http"
	postgresadapter "github.com/rafaelpissolatto/formUnity/backend/internal/adapter/repository/postgres"
	slackadapter "github.com/rafaelpissolatto/formUnity/backend/internal/adapter/slack"
	"github.com/rafaelpissolatto/formUnity/backend/internal/domain/volunteer"
	"github.com/slack-go/slack"

	_ "github.com/lib/pq"
)

func RunAPIServer(_ context.Context) {
	log.Println("running api server...")
	// database pg
	clientDb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("error connecting to database")
		panic(err)
	}
	defer clientDb.Close()

	volunteerRepo := postgresadapter.NewPostgresVolunteerRepository(clientDb)

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
	http.HandleFunc("/volunteers", volunteerHandler.AddVolunteer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

package slack

import (
	"github.com/rafaelpissolatto/formUnity/backend/internal/domain/volunteer"
	"github.com/slack-go/slack"
)

type SlackNotifier struct {
	clientAPI *slack.Client
}

func NewSlackNotifierWithOptions(options ...func(*SlackNotifier)) *SlackNotifier {
	s := &SlackNotifier{}
	for _, option := range options {
		option(s)
	}
	return s
}

func WithToken(token *slack.Client) func(*SlackNotifier) {
	return func(s *SlackNotifier) {
		s.clientAPI = token
	}
}

func (s *SlackNotifier) Notify(v volunteer.Volunteer, message string) error {
	attachment := slack.Attachment{
		Pretext: "New volunteer registered",
		Text:    message,
	}
	_, _, err := s.clientAPI.PostMessage(v.Email, slack.MsgOptionAttachments(attachment))
	return err
}

package service

import (
	"fmt"
	"weather_forecast_sub/internal/config"
	"weather_forecast_sub/pkg/email"
	"weather_forecast_sub/pkg/logger"
)

type EmailService struct {
	sender      email.Sender
	emailConfig config.EmailConfig
	httpConfig  config.HTTPConfig
}

type confirmationEmailInput struct {
	ConfirmationLink string
}

func NewEmailsService(sender email.Sender, emailConfig config.EmailConfig, httpConfig config.HTTPConfig) *EmailService {
	return &EmailService{
		sender:      sender,
		emailConfig: emailConfig,
		httpConfig:  httpConfig,
	}
}

func (s *EmailService) SendConfirmationEmail(input ConfirmationEmailInput) error {
	subject := s.emailConfig.Subjects.Confirmation

	templateInput := confirmationEmailInput{
		ConfirmationLink: s.createConfirmationLink(input.Token),
	}
	sendInput := email.SendEmailInput{Subject: subject, To: input.Email}

	if err := sendInput.GenerateBodyFromHTML(s.emailConfig.Templates.Confirmation, templateInput); err != nil {
		logger.Errorf("failed to generate confirmation email body: %s", err.Error())
		return err
	}

	return s.sender.Send(sendInput)
}

func (s *EmailService) createConfirmationLink(token string) string {
	return fmt.Sprintf("http://%s/api/confirm/%s", s.httpConfig.Domain, token)
}

func (s *EmailService) SendWeatherForecastEmail(input WeatherForecastEmailInput) error {
	// TODO: implement this method
	return nil
}

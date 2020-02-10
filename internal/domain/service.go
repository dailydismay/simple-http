package domain

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Service interface {
	GetUserAgent(r *http.Request) (*UserAgent, error)
}

type service struct {
	logger       logrus.FieldLogger
}

func NewService(logger logrus.FieldLogger) Service {
	return &service{
		logger,
	}
}

func (s *service) GetUserAgent (r *http.Request) (*UserAgent, error) {
	userAgent := r.Header.Get("User-Agent")

	if len(userAgent) == 0 {
		err := errors.New("no User-Agent header got")
		s.logger.WithError(err).Error("Client sent no User-Agent header")
		return nil, err
	}

	return &UserAgent{userAgent}, nil
}

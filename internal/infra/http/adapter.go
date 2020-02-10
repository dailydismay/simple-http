package http

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"simple-http/internal/configs"
	"simple-http/internal/domain"
)

type Adapter interface {
	ListenAndServe() error
}

type adapter struct {
	logger  logrus.FieldLogger
	config  *configs.Config
	service domain.Service
	server  *http.Server
}

func NewAdapter(logger logrus.FieldLogger, config *configs.Config, service domain.Service) (Adapter, error) {
	a := &adapter{
		logger: logger,
		config: config,
		service: service,
	}

	r := http.NewServeMux()
	r.HandleFunc("/user-agent", a.getUserAgent)

	a.server = &http.Server{Addr: config.Port, Handler: r }

	return a, nil
}

func (a *adapter) ListenAndServe() error {
	a.logger.WithField("Port", a.config.Port).Info("Listening and serving HTTP requests.")
	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.logger.WithError(err).Error("Error listening and serving HTTP requests!")
		return err
	}
	return nil
}

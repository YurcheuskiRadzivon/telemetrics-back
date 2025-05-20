package auth

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type Auth struct {
	vldtr *validator.Validate
	lgr   *logger.Logger
}

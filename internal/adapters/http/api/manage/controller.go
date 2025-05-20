package manage

import (
	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type Manage struct {
	vldtr *validator.Validate
	lgr   *logger.Logger
}

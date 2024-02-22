package applicator

import (
	"mvc/internal/config"

	"go.uber.org/zap"
)

type Applicator struct {
	logger *zap.SugaredLogger
	config *config.Config
}
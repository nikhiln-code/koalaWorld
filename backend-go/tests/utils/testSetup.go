package utils

import (
	"github.com/nikhiln-code/koalaWorld/backend-go/internal/logger"
)

func SetupTestLogger() {
	logger.Init(false)
}

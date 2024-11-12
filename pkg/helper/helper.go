package helper

import (
	"github.com/go-playground/form/v4" // New import
	"github.com/yeungon/corpora/internal/config"
)

var appConfig *config.AppConfig

type Helper struct {
	FormDecoder *form.Decoder
}

func New(cf *config.AppConfig) *Helper {
	appConfig = cf
	return &Helper{
		FormDecoder: form.NewDecoder(),
	}
}

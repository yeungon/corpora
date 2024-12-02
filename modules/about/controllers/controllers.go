package about

import (
	"github.com/yeungon/corpora/internal/config"
	"github.com/yeungon/corpora/pkg/helper"
)

type Controller struct {
	//articleService ArticleService.ArticleServiceInterface
	config *config.AppConfig
	helper *helper.Helper
}

func New(cf *config.AppConfig) *Controller {
	return &Controller{
		//articleService: ArticleService.New(),
		config: cf,
		helper: helper.New(cf),
	}
}

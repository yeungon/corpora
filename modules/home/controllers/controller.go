package home

import "github.com/yeungon/corpora/internal/config"

type Controller struct {
	//articleService ArticleService.ArticleServiceInterface
	config *config.AppConfig
	//helper *helpers.Helper
}

func New(cf *config.AppConfig) *Controller {
	return &Controller{
		//articleService: ArticleService.New(),
		config: cf,
		//helper: helpers.New(cf),
	}
}

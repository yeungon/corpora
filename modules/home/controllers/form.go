package home

import "github.com/yeungon/corpora/pkg/validator"

type postCreateForm struct {
	Input_search        string `form:"keyword"`
	validator.Validator `form:"-"`
}

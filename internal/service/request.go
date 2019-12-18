package service

import (
	"github.com/astaxie/beego/validation"

	"Go-blog-server/pkg/logging"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}

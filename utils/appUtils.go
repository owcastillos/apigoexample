package utils

import (
	"log"

	"github.com/owcastillos/apigoexample/models"
)

func HandleError(err error) *models.Result {
	log.Println(err)
	return &models.Result{
		Status: err.Error(),
	}
}

func HandleString(err string) *models.Result {
	log.Println(err)
	return &models.Result{
		Status: err,
	}
}

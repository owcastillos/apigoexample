package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/owcastillos/apigoexample/models"
	dataPersistence "github.com/owcastillos/apigoexample/persistence"
	"github.com/owcastillos/apigoexample/utils"
)

func main() {
	fmt.Println("Initializing service")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", httpHandler)
	router.HandleFunc("/{idUser}", httpHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), router))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result interface{} = &models.Result{
		Status: "FAILED",
	}
	switch method := r.Method; method {
	case http.MethodGet:
		mapVars := mux.Vars(r)
		if mapVars["idUser"] == "" {
			result = dataPersistence.GetUsers()
		} else {
			result = dataPersistence.GetUserById(mapVars["idUser"])
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodPut:
		err, user := getUserFromBody(r.Body)
		if err != nil {
			result = utils.HandleError(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			result = dataPersistence.InsertUser(user)
			if strings.Index((result.(*models.Result)).Status, "ID:") == 0 {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	case http.MethodPost:
		err, user := getUserFromBody(r.Body)
		if err != nil {
			result = utils.HandleError(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			result = dataPersistence.UpdateUser(user)
			if (result.(*models.Result)).Status == "OK" {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	case http.MethodDelete:
		mapVars := mux.Vars(r)
		if mapVars["idUser"] == "" {
			result = dataPersistence.DeleteUsers()
		} else {
			result = dataPersistence.DeleteUserById(mapVars["idUser"])
		}
		if (result.(*models.Result)).Status == "OK" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	json.NewEncoder(w).Encode(result)
}

func getUserFromBody(body io.ReadCloser) (error, *models.User) {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	var user models.User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, &user
}

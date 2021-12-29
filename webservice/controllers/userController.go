package controllers

import (
	"encoding/json"
	"github.com/jaisonjose89m/go/webservice/models"
	"net/http"
	"regexp"
	"strconv"
)

type userController struct {
	userIDRegex *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("Method not supported"))
		}
	} else {
		id, err := uc.getID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to get user ID from path"))
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("Method not supported"))
		}
	}
}

func (uc userController) parseRequest(w http.ResponseWriter, r *http.Request) (models.User, error) {
	decoder := json.NewDecoder(r.Body)
	var u models.User
	err := decoder.Decode(&u)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return models.User{}, err
	}
	return u, nil
}

func writeError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

func newUserController() *userController {
	return &userController{
		userIDRegex: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (uc userController) getAll(w http.ResponseWriter) {
	EncodeAsJsonString(models.GetUsers(), w)
}

func (uc userController) get(id int, w http.ResponseWriter) {
	user, err := models.GetUserByID(id)
	if err != nil {
		writeError(w, err, http.StatusNotFound)
		return
	}
	EncodeAsJsonString(user, w)
}

func (uc userController) post(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(w, r)
	if err != nil {
		return
	}
	addedUser, err := models.AddUser(user)
	if err != nil {
		writeError(w, err, http.StatusUnprocessableEntity)
		return
	}
	EncodeAsJsonString(addedUser, w)
}

func (uc userController) put(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(w, r)
	if err != nil {
		return
	}
	user, err := models.UpdateUser(u)
	if err != nil {
		writeError(w, err, http.StatusNotFound)
		return
	}
	EncodeAsJsonString(user, w)
}

func (uc userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		writeError(w, err, http.StatusNotFound)
		return
	}
	_, writeErr := w.Write([]byte("Successfully deleted user"))
	if writeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (uc userController) getID(r *http.Request) (int, error) {
	matches := uc.userIDRegex.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		return 0, nil
	}
	return strconv.Atoi(matches[1])
}

package service

import (
	"cde/storage"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type CrateResponse struct {
	ID int `json:"ID"`
}

type Service struct {
	Storage *storage.Storage
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u storage.User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		usedID, err := s.Storage.Save(&u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		response := CrateResponse{ID: usedID}
		responseBody, _ := json.Marshal(response)
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u storage.MakeFriendsRequest
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ans, err := s.Storage.Add(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ans))
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Service) DelieteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u storage.DelieteUserRequest
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ans, err := s.Storage.DelieteU(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ans))
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Service) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parts := strings.Split(r.URL.Path, "/")
		userID := parts[len(parts)-1]

		friends, err := s.Storage.GetU(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		responseBody, err := json.Marshal(friends)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Service) NewAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		parts := strings.Split(r.URL.Path, "/")
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u storage.NewAgeRequest
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		userID := parts[len(parts)-1]
		userAge, err := s.Storage.NewAgeU(u, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(userAge))
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Service) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		user := s.Storage.AllUsers()
		responseBody, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

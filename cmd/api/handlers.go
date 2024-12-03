package main

import (
	"net/http"

	"github.com/myselfBZ/users/internal/store"
)

var successMsg = map[string]bool{
	"success": true,
}

const userKey = "user"

type apiResponse struct {
	Err    error `json:"error"`
	Data   any   `json:"data"`
	status int
}

type userCreateRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}

type apiHandler func(http.ResponseWriter, *http.Request) *apiResponse

// godoc
// @Summary		 wraps the apiHandler function
func makeHTTPHandler(f apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiResp := f(w, r)
		if apiResp.Err != nil {
			writeJSON(w, apiResp, apiResp.status)
			return
		}
		writeJSON(w, apiResp.Data, apiResp.status)
	}
}

func (a *API) registerUser(w http.ResponseWriter, r *http.Request) *apiResponse {
	resp := &apiResponse{}
	var userCrt userCreateRequest
	if err := readJSON(w, r, &userCreateRequest{}); err != nil {
		resp.Err = err
		resp.status = http.StatusBadRequest
		return resp
	}

	u := &store.User{
		Name:     userCrt.Name,
		LastName: userCrt.LastName,
		Password: userCrt.Password,
		Email:    userCrt.Email,
	}

	if err := a.store.Create(u); err != nil {
		resp.Err = err
		resp.status = http.StatusInternalServerError
		return resp
	}

	resp.Data = successMsg
	resp.status = http.StatusOK
	return resp
}

func (a *API) login(w http.ResponseWriter, r *http.Request) *apiResponse {
	return nil
}

func (a *API) deleteAccount(w http.ResponseWriter, r *http.Request) *apiResponse {
	return nil
}

func (a *API) updateAccount(w http.ResponseWriter, r *http.Request) *apiResponse {
	return nil
}

package users

import (
	"encoding/json"
	"net/http"

	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/helpers"
	"github.com/adiet95/Golang/GoRent/src/interfaces"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(reps interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: reps}
}

func (re *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	re.svc.GetAll().Send(w)
}

func (re *user_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Add(&datas).Send(w)
}

func (re *user_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("email")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		helpers.New(err.Error(), 400, true)
		return
	}
	re.svc.Update(&datas, val).Send(w)
}

func (re *user_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("email")
	re.svc.Delete(val).Send(w)
}

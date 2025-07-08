package api

import (
	"net/http"

	"github.com/joaopedroldavid-del/gobid/internal/json_utils"
	"github.com/joaopedroldavid-del/gobid/internal/usecase/user"
)

func (a *Api) HandleSignUpUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserRequest](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
	}
}
func (a *Api) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO: Implement HandleLoginUser")
}

func (a *Api) HandleLogOutUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO: Implement HandleLogOutUser")
}
package controller

import (
	nethttp "net/http"

	"path"
	"strconv"

	"github.com/kustavo/benchmark/go/application/dto"
	"github.com/kustavo/benchmark/go/application/interfaces"
	"github.com/kustavo/benchmark/go/domain/message"
	"github.com/kustavo/benchmark/go/domain/model"
	"github.com/kustavo/benchmark/go/infrastructure/http"
)

type userController struct {
	application *interfaces.Application
}

func NewUserController(application *interfaces.Application) *userController {
	return &userController{
		application: application,
	}
}

// swagger:operation POST /user/ user create
//
// Create new user
//
// This will create a new user
//
// ---
// consumes:
// - application/json
// parameters:
// - name: user
//   in: body
//   description: example parameters
//   schema:
//     "$ref": "#/definitions/user"
//   required: true
// responses:
//   default:
//     schema:
//       "$ref": "#/definitions/responseMessage"
func (h *userController) Create(w nethttp.ResponseWriter, r *nethttp.Request) {
	user := h.application.Commands.CreateUserRequest
	err := http.RequestModel(w, r, &user)
	if err != nil {
		http.ResponseBadRequest(w, err, nil)
		return
	}

	id, err := h.application.Commands.CreateUserCommand.Handle(r.Context(), &user)
	msg := []string{message.SuccessUserCreated}
	http.Response(w, err, msg, id)
}

// swagger:operation PUT /user/{id} user update
//
// Update a user
//
// This will update a user
//
// ---
// consumes:
// - application/json
// securityDefinitions:
//   isRegistered:
//     type: basic
// parameters:
// - name: id
//   in: path
//   description: user id
//   required: true
//   type: string
// - name: user
//   in: body
//   description: example parameters
//   schema:
//     "$ref": "#/definitions/user"
//   required: true
// responses:
//   default:
//     schema:
//       "$ref": "#/definitions/responseMessage"
func (h *userController) Update(w nethttp.ResponseWriter, r *nethttp.Request) {

	var user model.User
	err := http.RequestEntity(w, r, &user)
	if err != nil {
		http.ResponseBadRequest(w, err, nil)
		return
	}

	err = h.application.Commands.UpdateUserCommand.Handle(r.Context(), &user)
	http.Response(w, err, nil, nil)
}

// swagger:operation DELETE /user/{id} user delete
//
// Delete user
//
// This will delete a user
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: user id
//   required: true
//   type: string
// responses:
//   default:
//     schema:
//       "$ref": "#/definitions/responseMessage"
func (h *userController) Delete(w nethttp.ResponseWriter, r *nethttp.Request) {

	idStr := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(string(idStr), 10, 64)
	if err != nil {
		http.ResponseBadRequest(w, err, nil)
		return
	}

	err = h.application.Commands.DeleteUserCommand.Handle(r.Context(), id)
	http.Response(w, err, nil, nil)

}

// swagger:operation GET /user/get-by-email user getByEmail
//
// Get user by email
//
// This will get a user by email
//
// ---
// produces:
// - application/json
// parameters:
// - name: email
//   in: query
//   description: user email
//   required: true
//   type: string
// responses:
//   default:
//     schema:
//       "$ref": "#/definitions/responseMessage"
func (h *userController) GetByEmail(w nethttp.ResponseWriter, r *nethttp.Request) {

	email := r.URL.Query().Get("email")
	data, err := h.application.Queries.GetUserByEmailQuery.Handle(r.Context(), email)
	http.Response(w, err, nil, &data)
}

// swagger:operation GET /user/{id} user getById
//
// Get user by id
//
// This will get a user by id
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: user id
//   required: true
//   type: string
// responses:
//   default:
//     schema:
//       "$ref": "#/definitions/responseMessage"
func (h *userController) GetById(w nethttp.ResponseWriter, r *nethttp.Request) {

	idStr := path.Base(r.URL.Path)
	id, err := strconv.ParseUint(string(idStr), 10, 64)
	if err != nil {
		http.ResponseBadRequest(w, err, nil)
		return
	}

	data, err := h.application.Queries.GetUserByIdQuery.Handle(r.Context(), id)
	http.Response(w, err, nil, &data)
}

// swagger:operation POST /login user login
//
// Authenticate
//
// This will create the auth and refresh tokens
//
// ---
// security: []
// consumes:
// - application/json
// parameters:
// - name: user
//   in: body
//   description: example parameters
//   schema:
//     "$ref": "#/definitions/user"
//   required: true
// responses:
//   default:
//     schema:
//       "$ref": "#/definitions/responseMessage"
func (h *userController) Login(w nethttp.ResponseWriter, r *nethttp.Request) {

	var user model.User
	err := http.RequestModel(w, r, &user)
	if err != nil {
		http.ResponseBadRequest(w, err, nil)
		return
	}

	u, err := h.application.Queries.GetUserCredentialsQuery.Handle(r.Context(), user.Username, user.Password)
	if err != nil {
		http.ResponseBadRequest(w, err, nil)
		return
	}

	tokens, err := h.application.Authentication.CreateAuth(r.Context(), u.ID)
	http.Response(w, err, nil, tokens)
}

// swagger:operation POST /logout user logout
//
// Logout
//
// This will delete the auth and refresh tokens
//
// ---
// consumes:
// - application/json
// responses:
//   default:
//     schema:
//       "$ref": "#/definitions/responseMessage"
func (h *userController) Logout(w nethttp.ResponseWriter, r *nethttp.Request) {

	tokenString := http.RequestToken(r)
	err := h.application.Authentication.DeleteAuth(r.Context(), tokenString)
	http.Response(w, err, nil, nil)
}

// swagger:operation POST /refresh-auth user refresh-auth
//
// Refresh auth
//
// This will creates the new auth and refresh tokens and delete the old ones
//
// ---
// consumes:
// - application/json
// parameters:
// - name: token
//   in: body
//   description: token
//   schema:
//     "$ref": "#/definitions/tokenDTO"
//   required: true
// responses:
//   default:
//     schema:
//       "$ref": "#/definitions/responseMessage"
func (h *userController) RefreshAuth(w nethttp.ResponseWriter, r *nethttp.Request) {

	accessTokenString := http.RequestToken(r)

	var refreshToken dto.TokenDTO
	err := http.RequestModel(w, r, &refreshToken)
	if err != nil {
		http.ResponseBadRequest(w, err, nil)
		return
	}

	tokens, err := h.application.Authentication.RefreshAuth(r.Context(), accessTokenString, refreshToken.Token)
	http.Response(w, err, nil, tokens)
}

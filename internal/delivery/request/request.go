package request

import (
	"errors"
	"fmt"
	modelsData "forum/internal/models/data"
	modelsRequest "forum/internal/models/request"
	modelsUser "forum/internal/models/user"
	"forum/pkg"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func (h *RequestHandler) CreateRequest(w http.ResponseWriter, r *http.Request) {
	var tmplData modelsData.Data
	if r.URL.Path != "/request/create" {
		pkg.Error(w, http.StatusNotFound, nil)
		return
	}
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		fmt.Println("ERROR : unauthorized user")
		// Обработка случая, когда пользователь не аутентифицирован
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}

	user, ok := userValue.(modelsUser.User)
	if !ok {
		// Обработка случая, когда значение в контексте не является типом User
		fmt.Println("ERROR : invalid user data in context")
		pkg.Error(w, http.StatusInternalServerError, nil)
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/templates/create_request.html")
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, tmplData)
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}

		request := modelsRequest.Request{
			UserID:    user.ID,
			CreatedAt: time.Now(),
			Status:    "created",
		}

		if _, err := h.request.CreateRequest(request); err != nil {
			if errors.Is(err, errors.New("ErrEmptyBody")) {
				pkg.Error(w, http.StatusBadRequest, err)
				return
			}
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
	}
}

func (h *RequestHandler) UpdateRequestStatus(w http.ResponseWriter, r *http.Request) {
	// var tmplData modelsData.Data
	if r.URL.Path != "/admin/requests/change" {
		pkg.Error(w, http.StatusNotFound, nil)
		return
	}
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		fmt.Println("ERROR : unauthorized user")
		// Обработка случая, когда пользователь не аутентифицирован
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}
	if r.Method != http.MethodPost {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		return
	}
	user, ok := userValue.(modelsUser.User)
	if !ok {
		// Обработка случая, когда значение в контексте не является типом User
		fmt.Println("ERROR : invalid user data in context")
		pkg.Error(w, http.StatusInternalServerError, nil)
		return
	}
	if user.Role != "admin" {
		pkg.Error(w, http.StatusBadRequest, nil)
		return
	}
	requestID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect request id"))
		return
	}
	requestUserID, err := strconv.Atoi(r.FormValue("userId"))
	if err != nil {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect request user id"))
		return
	}
	status := r.FormValue("status")
	fmt.Println(status)
	if status != "accept" && status != "decline" {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect request status"))
		return
	}

	if err := h.request.UpdateRequestStatus(requestID, requestUserID, status); err != nil {
		pkg.Error(w, http.StatusInternalServerError, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *RequestHandler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users/type/change" {
		pkg.Error(w, http.StatusNotFound, nil)
		return
	}

	if r.Method != http.MethodPost {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		return
	}
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		fmt.Println("ERROR : unauthorized user")
		// Обработка случая, когда пользователь не аутентифицирован
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}
	user, ok := userValue.(modelsUser.User)
	if !ok {
		// Обработка случая, когда значение в контексте не является типом User
		fmt.Println("ERROR : invalid user data in context")
		pkg.Error(w, http.StatusInternalServerError, nil)
		return
	}

	if user.Role != "admin" {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect user Type"))
		return
	}
	userId, err := strconv.Atoi(r.FormValue("user-id"))
	if err != nil {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect userId Id "))
		return
	}
	userNewType := r.FormValue("type")

	if userNewType != "user" && userNewType != "moderator" {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect user new Type"))
		return
	}

	if err := h.request.UpdateUserRole(userId, userNewType); err != nil {
		pkg.Error(w, http.StatusBadRequest, err)
		return
	}

	http.Redirect(w, r, "/users", http.StatusFound)
}

func (h *RequestHandler) GetAllRequests(w http.ResponseWriter, r *http.Request) {
	var tmplData modelsData.Data
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		fmt.Println("ERROR : unauthorized user")
		// Обработка случая, когда пользователь не аутентифицирован
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}

	user, ok := userValue.(modelsUser.User)
	if !ok {
		// Обработка случая, когда значение в контексте не является типом User
		fmt.Println("ERROR : invalid user data in context")
		pkg.Error(w, http.StatusInternalServerError, nil)
		return
	}
	if user.Role == "admin" {
		if r.Method == http.MethodGet {
			requests, err := h.request.GetAllRequests()
			if err != nil {
				pkg.Error(w, http.StatusInternalServerError, err)
				return
			}
			tmplData.Requests = requests
			tmplData.User = user
			tmpl, err := template.ParseFiles("./ui/templates/allRequests.html")
			if err != nil {
				pkg.Error(w, http.StatusInternalServerError, err)
				return
			}
			tmpl.Execute(w, tmplData)
		}
	} else {
		pkg.Error(w, http.StatusBadRequest, errors.New("Only for admins"))
		return
	}
}

func (h *RequestHandler) MyRequest(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("hello")
	if r.URL.Path != "/my-request" {
		pkg.Error(w, http.StatusNotFound, nil)
		return
	}
	var tmplData modelsData.Data
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		fmt.Println("ERROR : unauthorized user")
		// Обработка случая, когда пользователь не аутентифицирован
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}

	user, ok := userValue.(modelsUser.User)
	if !ok {
		// Обработка случая, когда значение в контексте не является типом User
		fmt.Println("ERROR : invalid user data in context")
		pkg.Error(w, http.StatusInternalServerError, nil)
		return
	}
	if r.Method == http.MethodGet {
		request, err := h.request.GetRequestByUserID(user.ID)
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmplData.Request = request
		tmplData.User = user
		tmpl, err := template.ParseFiles("./ui/templates/my-request.html")
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, tmplData)
	}
}

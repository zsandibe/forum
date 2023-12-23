package report

import (
	"errors"
	"fmt"
	modelsData "forum/internal/models/data"
	modelsReport "forum/internal/models/report"
	modelsUser "forum/internal/models/user"
	"forum/pkg"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/moderator/report/create" {
		pkg.Error(w, http.StatusNotFound, nil)
		return
	}
	// var tmplData modelsData.Data
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

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}

		if user.Role != "moderator" {
			pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect user Type"))
			return
		}
		postId, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			pkg.Error(w, http.StatusBadRequest, nil)
			return
		}
		var report modelsReport.Report

		report.CreatedBy = user.ID
		report.CreatedAt = time.Now()
		report.PostID = postId
		report.ModeratorMsg = r.FormValue("report-text")
		report.Status = "created"
		_, err = h.report.CreateReport(report)
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, errors.New(err.Error()))
			return
		}
		http.Redirect(w, r, "/my-reports", http.StatusSeeOther)
	} else {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
	}
}

func (h *ReportHandler) GetAllReports(w http.ResponseWriter, r *http.Request) {
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
			reports, err := h.report.GetAllReports()
			if err != nil {
				pkg.Error(w, http.StatusInternalServerError, err)
				return
			}
			tmplData.Reports = reports
			tmplData.User = user
			tmpl, err := template.ParseFiles("./ui/templates/all-reports.html")
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

func (h *ReportHandler) GetReportByUserID(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/my-reports" {
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
		reports, err := h.report.GetReportByUserID(user.ID)
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmplData.Reports = reports
		tmplData.User = user
		tmpl, err := template.ParseFiles("./ui/templates/myreport.html")
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, tmplData)
	}
}

func (h *ReportHandler) UpdateReportStatus(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/reports/change" {
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

	if r.Method != http.MethodPost {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		return
	}

	if user.Role != "admin" {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect user type"))
		return
	}
	reportId, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect reportId Id "))
	}

	report, err := h.report.GetReportByID(reportId)
	if err != nil {
		pkg.Error(w, http.StatusInternalServerError, errors.New(err.Error()))
	}
	report.AdminMsg = r.FormValue("admin-text")

	status := r.FormValue("status")
	if status != "accept" && status != "decline" {
		pkg.Error(w, http.StatusBadRequest, errors.New("Incorrect modRequest status"))
	}
	report.Status = r.FormValue("status")

	err = h.report.UpdateReportStatus(report)
	if err != nil {
		pkg.Error(w, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}

	http.Redirect(w, r, "/admin/reports", http.StatusFound)
}

package report

import (
	"database/sql"
	reportModels "forum/internal/models/report"
)

type Report interface {
	CreateReport(report reportModels.Report) (int64, error)
	GetReportByUserID(userID int) ([]reportModels.Report, error)
	GetAllReports() ([]reportModels.Report, error)
	UpdateReportStatus(report reportModels.Report) error
	GetReportByID(reportID int) (reportModels.Report, error)
}

type ReportSql struct {
	db *sql.DB
}

func NewReportSql(db *sql.DB) *ReportSql {
	return &ReportSql{
		db: db,
	}
}

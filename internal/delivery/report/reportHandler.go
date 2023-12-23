package report

import (
	reportModels "forum/internal/models/report"
)

type reportProvider interface {
	CreateReport(report reportModels.Report) (int64, error)
	GetReportByUserID(userID int) ([]reportModels.Report, error)
	GetAllReports() ([]reportModels.Report, error)
	UpdateReportStatus(report reportModels.Report) error
	GetReportByID(reportID int) (reportModels.Report, error)
}

type ReportHandler struct {
	report reportProvider
}

func NewReportHandler(report reportProvider) *ReportHandler {
	return &ReportHandler{
		report: report,
	}
}

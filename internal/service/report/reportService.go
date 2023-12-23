package report

import (
	reportModels "forum/internal/models/report"
	repositoryReport "forum/internal/repository/report"
)

type Report interface {
	CreateReport(report reportModels.Report) (int64, error)
	GetReportByUserID(userID int) ([]reportModels.Report, error)
	GetAllReports() ([]reportModels.Report, error)
	UpdateReportStatus(report reportModels.Report) error
	GetReportByID(reportID int) (reportModels.Report, error)
}

type ReportService struct {
	repo repositoryReport.Report
}

func NewReportService(repo repositoryReport.Report) *ReportService {
	return &ReportService{
		repo: repo,
	}
}

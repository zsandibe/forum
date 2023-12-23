package report

import (
	"database/sql"
	"errors"
	reportModels "forum/internal/models/report"
)

func (s *ReportService) CreateReport(report reportModels.Report) (int64, error) {
	return s.repo.CreateReport(report)
}

func (s *ReportService) GetReportByUserID(userID int) ([]reportModels.Report, error) {
	report, err := s.repo.GetReportByUserID(userID)
	if err == sql.ErrNoRows {
		return []reportModels.Report{}, errors.New("report not found")
	} else if err != nil {
		return report, err
	}
	return report, err
}

func (s *ReportService) GetAllReports() ([]reportModels.Report, error) {
	return s.repo.GetAllReports()
}

func (s *ReportService) UpdateReportStatus(report reportModels.Report) error {
	return s.repo.UpdateReportStatus(report)
}

func (s *ReportService) GetReportByID(reportID int) (reportModels.Report, error) {
	return s.repo.GetReportByID(reportID)
}

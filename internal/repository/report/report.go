package report

import (
	"database/sql"
	"fmt"
	reportModels "forum/internal/models/report"
)

func (r *ReportSql) CreateReport(report reportModels.Report) (int64, error) {
	query := `
		INSERT INTO	reports (UserID,PostID,Created_at,Moderator_msg,Status) VALUES (?,?,?,?,?)
	`
	res, err := r.db.Exec(query, &report.CreatedBy, &report.PostID, &report.CreatedAt, &report.ModeratorMsg, &report.Status)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return id, nil
}

func (r *ReportSql) GetReportByUserID(reportID int) ([]reportModels.Report, error) {
	var reports []reportModels.Report
	query := `
		SELECT	reports.ID,reports.UserID,reports.PostID,reports.Created_at,reports.Moderator_msg,reports.Status,reports.Admin_msg FROM reports
		INNER JOIN users on  users.ID = reports.UserID
		WHERE reports.UserID = $1
	`
	rows, err := r.db.Query(query, reportID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var report reportModels.Report
		var adminMsg sql.NullString
		if err := rows.Scan(&report.ID, &report.CreatedBy, &report.PostID, &report.CreatedAt, &report.ModeratorMsg, &report.Status, &adminMsg); err != nil {
			return reports, err
		}
		report.AdminMsg = adminMsg.String
		reports = append(reports, report)
	}
	return reports, nil
}

func (r *ReportSql) GetAllReports() ([]reportModels.Report, error) {
	var reports []reportModels.Report

	query := `
		SELECT reports.ID,reports.UserID,reports.PostID,reports.Created_at,reports.Moderator_msg, reports.Status,reports.Admin_msg FROM reports
		ORDER By reports.ID DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var adminMsg sql.NullString
		var report reportModels.Report
		if err := rows.Scan(&report.ID, &report.CreatedBy, &report.PostID, &report.CreatedAt, &report.ModeratorMsg, &report.Status, &adminMsg); err != nil {
			return reports, err
		}
		report.AdminMsg = adminMsg.String
		reports = append(reports, report)
	}
	return reports, nil
}

func (r *ReportSql) UpdateReportStatus(report reportModels.Report) error {
	query := `
		UPDATE reports SET Moderator_msg =?,Admin_msg = ?, Status = ? WHERE ID = $1
	`

	if _, err := r.db.Exec(query, report.ModeratorMsg, report.AdminMsg, report.Status, report.ID); err != nil {
		return fmt.Errorf("can`t update report: %w", err)
	}
	return nil
}

func (r *ReportSql) GetReportByID(reportID int) (reportModels.Report, error) {
	query := `
		SELECT * FROM reports WHERE ID = $1
	`
	var report reportModels.Report
	var adminMsg sql.NullString
	var updatedAt sql.NullTime
	row := r.db.QueryRow(query, reportID)
	err := row.Scan(&report.ID, &report.CreatedBy, &report.PostID, &report.CreatedAt, &updatedAt, &report.ModeratorMsg, &adminMsg, &report.Status)
	if err != nil {
		fmt.Println("TUT")
		return report, err
	}
	report.UpdatedAt = updatedAt.Time
	report.AdminMsg = adminMsg.String
	return report, nil
}

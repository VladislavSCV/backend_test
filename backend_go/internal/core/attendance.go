package core

import (
	"database/sql"
	"fmt"
	"github.com/VladislavSCV/internal/models"
)

func GetAttendanceByStudentID(db *sql.DB, studentID int) ([]models.Attendance, error) {
	var attendances []models.Attendance

	rows, err := db.Query("SELECT id, student_id, subject_id, date, status, created_at, updated_at FROM attendance WHERE student_id = $1", studentID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch attendance: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		if err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.SubjectID, &attendance.Date, &attendance.Status, &attendance.CreatedAt, &attendance.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan attendance: %v", err)
		}
		attendances = append(attendances, attendance)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over attendance: %v", err)
	}

	return attendances, nil
}

func GetAttendanceByGroupID(db *sql.DB, groupID int) ([]models.Attendance, error) {
	var attendances []models.Attendance

	rows, err := db.Query(`
        SELECT a.id, a.student_id, a.subject_id, a.date, a.status, a.created_at, a.updated_at
        FROM attendance a
        JOIN users u ON a.student_id = u.id
        WHERE u.group_id = $1
    `, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch attendance: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		if err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.SubjectID, &attendance.Date, &attendance.Status, &attendance.CreatedAt, &attendance.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan attendance: %v", err)
		}
		attendances = append(attendances, attendance)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over attendance: %v", err)
	}

	return attendances, nil
}

func CreateAttendance(db *sql.DB, attendance models.Attendance) (int, error) {
	var attendanceID int

	err := db.QueryRow(`
        INSERT INTO attendance (student_id, subject_id, date, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        RETURNING id
    `, attendance.StudentID, attendance.SubjectID, attendance.Date, attendance.Status).Scan(&attendanceID)
	if err != nil {
		return 0, fmt.Errorf("failed to create attendance: %v", err)
	}

	return attendanceID, nil
}

func UpdateAttendance(db *sql.DB, attendance models.Attendance) error {
	_, err := db.Exec(`
        UPDATE attendance
        SET student_id = $1, subject_id = $2, date = $3, status = $4, updated_at = NOW()
        WHERE id = $5
    `, attendance.StudentID, attendance.SubjectID, attendance.Date, attendance.Status, attendance.ID)
	if err != nil {
		return fmt.Errorf("failed to update attendance: %v", err)
	}

	return nil
}

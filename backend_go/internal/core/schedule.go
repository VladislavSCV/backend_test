package core

import (
	"database/sql"
	"fmt"
	"github.com/VladislavSCV/internal/models"
)

func GetAllSchedules(db *sql.DB) ([]models.Schedule, error) {
	var schedules []models.Schedule

	rows, err := db.Query("SELECT id, group_id, subject_id, teacher_id, day_of_week, start_time, end_time, location, created_at, updated_at FROM schedules")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schedules: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var schedule models.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.SubjectID, &schedule.TeacherID, &schedule.DayOfWeek, &schedule.StartTime, &schedule.EndTime, &schedule.Location, &schedule.CreatedAt, &schedule.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan schedule: %v", err)
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over schedules: %v", err)
	}

	return schedules, nil
}
func GetScheduleByID(db *sql.DB, id int) ([]models.Schedule, error) {
	var schedules []models.Schedule

	rows, err := db.Query(`
        SELECT id, group_id, subject_id, teacher_id, day_of_week, start_time, end_time, location, created_at, updated_at
        FROM schedules
        WHERE group_id = $1 OR teacher_id = $1
    `, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schedule: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var schedule models.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.SubjectID, &schedule.TeacherID, &schedule.DayOfWeek, &schedule.StartTime, &schedule.EndTime, &schedule.Location, &schedule.CreatedAt, &schedule.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan schedule: %v", err)
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over schedules: %v", err)
	}

	return schedules, nil
}

func CreateSchedule(db *sql.DB, schedule models.Schedule) (int, error) {
	var scheduleID int

	err := db.QueryRow(`
        INSERT INTO schedules (group_id, subject_id, teacher_id, day_of_week, start_time, end_time, location, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
        RETURNING id
    `, schedule.GroupID, schedule.SubjectID, schedule.TeacherID, schedule.DayOfWeek, schedule.StartTime, schedule.EndTime, schedule.Location).Scan(&scheduleID)
	if err != nil {
		return 0, fmt.Errorf("failed to create schedule: %v", err)
	}

	return scheduleID, nil
}

func UpdateSchedule(db *sql.DB, schedule models.Schedule) error {
	_, err := db.Exec(`
        UPDATE schedules
        SET group_id = $1, subject_id = $2, teacher_id = $3, day_of_week = $4, start_time = $5, end_time = $6, location = $7, updated_at = NOW()
        WHERE id = $8
    `, schedule.GroupID, schedule.SubjectID, schedule.TeacherID, schedule.DayOfWeek, schedule.StartTime, schedule.EndTime, schedule.Location, schedule.ID)
	if err != nil {
		return fmt.Errorf("failed to update schedule: %v", err)
	}

	return nil
}

func DeleteSchedule(db *sql.DB, scheduleID int) error {
	_, err := db.Exec("DELETE FROM schedules WHERE id = $1", scheduleID)
	if err != nil {
		return fmt.Errorf("failed to delete schedule: %v", err)
	}

	return nil
}

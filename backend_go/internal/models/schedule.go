package models

import "time"

type Schedule struct {
    ID        int       `json:"id"`
    GroupID   int       `json:"group_id"`
    SubjectID int       `json:"subject_id"`
    TeacherID int       `json:"teacher_id"`
    DayOfWeek int       `json:"day_of_week"` // 1-7 (понедельник-воскресенье)
    StartTime string    `json:"start_time"`  // Формат "HH:MM"
    EndTime   string    `json:"end_time"`    // Формат "HH:MM"
    Location  string    `json:"location"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
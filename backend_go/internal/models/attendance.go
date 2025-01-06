package models

import "time"

type Attendance struct {
    ID        int       `json:"id"`
    StudentID int       `json:"student_id"`
    SubjectID int       `json:"subject_id"`
    Date      time.Time `json:"date"`
    Status    string    `json:"status"` // present, absent, excused
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
package models

import "time"

type User struct {
    ID          int       `json:"id"`
    FirstName   string    `json:"first_name"`
    MiddleName  string    `json:"middle_name"`
    LastName    string    `json:"last_name"`
    RoleID      int       `json:"role_id"`
    GroupID     *int      `json:"group_id,omitempty"`
    Login       string    `json:"login"`
    Password    string    `json:"-"`
    Salt        string    `json:"-"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
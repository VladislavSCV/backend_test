package core

import (
	"database/sql"
	"fmt"
	"github.com/VladislavSCV/internal/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	var users []models.User

	rows, err := db.Query("SELECT id, first_name, middle_name, last_name, role_id, group_id, login, created_at, updated_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.RoleID, &user.GroupID, &user.Login, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %v", err)
	}

	return users, nil
}

func GetStudents(db *sql.DB) ([]models.User, error) {
	var students []models.User

	rows, err := db.Query(`
        SELECT id, first_name, middle_name, last_name, role_id, group_id, login, created_at, updated_at
        FROM users
        WHERE role_id = (SELECT id FROM roles WHERE value = 'student')
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch students: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var student models.User
		if err := rows.Scan(&student.ID, &student.FirstName, &student.MiddleName, &student.LastName, &student.RoleID, &student.GroupID, &student.Login, &student.CreatedAt, &student.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan student: %v", err)
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over students: %v", err)
	}

	return students, nil
}

func GetTeachers(db *sql.DB) ([]models.User, error) {
	var teachers []models.User

	rows, err := db.Query(`
        SELECT id, first_name, middle_name, last_name, role_id, group_id, login, created_at, updated_at
        FROM users
        WHERE role_id = (SELECT id FROM roles WHERE value = 'teacher')
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch teachers: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var teacher models.User
		if err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.MiddleName, &teacher.LastName, &teacher.RoleID, &teacher.GroupID, &teacher.Login, &teacher.CreatedAt, &teacher.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan teacher: %v", err)
		}
		teachers = append(teachers, teacher)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over teachers: %v", err)
	}

	return teachers, nil
}

func GetUserByID(db *sql.DB, userID int) (*models.User, error) {
	var user models.User

	err := db.QueryRow(`
        SELECT id, first_name, middle_name, last_name, role_id, group_id, login, created_at, updated_at
        FROM users
        WHERE id = $1
    `, userID).Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.RoleID, &user.GroupID, &user.Login, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	return &user, nil
}

func UpdateUser(db *sql.DB, user models.User) error {
	_, err := db.Exec(`
        UPDATE users
        SET first_name = $1, middle_name = $2, last_name = $3, role_id = $4, group_id = $5, login = $6, updated_at = NOW()
        WHERE id = $7
    `, user.FirstName, user.MiddleName, user.LastName, user.RoleID, user.GroupID, user.Login, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func DeleteUser(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}

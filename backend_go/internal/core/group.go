package core

import (
	"database/sql"
	"fmt"
	"github.com/VladislavSCV/internal/models"
)

func GetAllGroups(db *sql.DB) ([]models.Group, error) {
	var groups []models.Group

	rows, err := db.Query("SELECT id, name, created_at, updated_at FROM groups")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.ID, &group.Name, &group.CreatedAt, &group.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan group: %v", err)
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over groups: %v", err)
	}

	return groups, nil
}

func GetGroupByID(db *sql.DB, groupID int) (*models.Group, error) {
	var group models.Group

	err := db.QueryRow(`
        SELECT id, name, created_at, updated_at
        FROM groups
        WHERE id = $1
    `, groupID).Scan(&group.ID, &group.Name, &group.CreatedAt, &group.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("group not found")
		}
		return nil, fmt.Errorf("failed to fetch group: %v", err)
	}

	return &group, nil
}

func CreateGroup(db *sql.DB, group models.Group) (int, error) {
	var groupID int

	err := db.QueryRow(`
        INSERT INTO groups (name, created_at, updated_at)
        VALUES ($1, NOW(), NOW())
        RETURNING id
    `, group.Name).Scan(&groupID)
	if err != nil {
		return 0, fmt.Errorf("failed to create group: %v", err)
	}

	return groupID, nil
}

func UpdateGroup(db *sql.DB, group models.Group) error {
	_, err := db.Exec(`
        UPDATE groups
        SET name = $1, updated_at = NOW()
        WHERE id = $2
    `, group.Name, group.ID)
	if err != nil {
		return fmt.Errorf("failed to update group: %v", err)
	}

	return nil
}

func DeleteGroup(db *sql.DB, groupID int) error {
	_, err := db.Exec("DELETE FROM groups WHERE id = $1", groupID)
	if err != nil {
		return fmt.Errorf("failed to delete group: %v", err)
	}

	return nil
}

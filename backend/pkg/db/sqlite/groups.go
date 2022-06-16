package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type GroupRepository struct {
	DB *sql.DB
}

func (repo *GroupRepository) GetAllAndRelations(userID string) ([]models.Group, error) {
	var groups []models.Group
	rows, err := repo.DB.Query("SELECT group_id, name, (SELECT COUNT(*) FROM group_users WHERE group_users.group_id = groups.group_id AND group_users.user_id = " + userID + ") as member, administrator = " + userID + " as admin FROM groups;")
	if err != nil {
		return groups, err
	}
	for rows.Next() {
		var group models.Group
		var member int
		var admin int
		rows.Scan(&group.ID, &group.Name, &member, &admin)
		if member != 0 {
			group.Member = true
		}
		if admin != 0 {
			group.Administrator = true
		}
		groups = append(groups, group)
	}
	return groups, nil
}

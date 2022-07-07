package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type NotifRepository struct {
	DB *sql.DB
}

// NOT TESTED
func (repo *NotifRepository) Save(notification models.Notification) error {
	stmt, err := repo.DB.Prepare("INSERT INTO notifications (notif_id, user_id,type,content) values (?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(notification.ID, notification.TargetID, notification.Type, notification.Content); err != nil {
		return err
	}
	return nil
}

func (repo *NotifRepository) Delete(notificationId string) error {
	_, err := repo.DB.Exec("DELETE FROM notifications WHERE notif_id=?", notificationId)
	if err != nil {
		return err
	}
	return nil
}

// NOT TESTED
func (repo *NotifRepository) GetGroupRequestsIds(groupId string) ([]string, error) {
	var requests = []string{}
	rows, err := repo.DB.Query("SELECT content FROM notifications WHERE user_id = ? AND type = 'GROUP_REQUEST';", groupId)
	if err != nil {
		return requests, err
	}
	for rows.Next() {
		var id string
		rows.Scan(&id)
		requests = append(requests, id)
	}
	return requests, nil
}

func (repo *NotifRepository) GetUserFromRequest(notificationId string) (string, error) {
	row := repo.DB.QueryRow("SELECT content FROM notifications WHERE notif_id = ? ", notificationId)
	var userId string
	if err := row.Scan(&userId); err != nil {
		return userId, err
	}
	return userId, nil
}

// NOT TESTED
func (repo *NotifRepository) CheckIfExists(notif models.Notification) (bool, error) {
	row := repo.DB.QueryRow("SELECT COUNT() FROM notifications WHERE user_id = ? AND content = ? AND type =? ", notif.TargetID, notif.Content, notif.Type)
	var resp int
	if err := row.Scan(&resp); err != nil {
		return false, err
	}
	if resp == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

//NOT TESTED
func (repo *NotifRepository) GetGroupId(notificationId string) (string, error) {
	row := repo.DB.QueryRow("SELECT content FROM notifications WHERE notif_id = ? ", notificationId)
	var groupId string
	if err := row.Scan(&groupId); err != nil {
		return groupId, err
	}
	return groupId, nil
}

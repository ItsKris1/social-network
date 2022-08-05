package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type MsgRepository struct {
	DB *sql.DB
}

func (repo *MsgRepository) Save(msg models.ChatMessage) error {
	stmt, err := repo.DB.Prepare("INSERT INTO messages (message_id, sender_id, receiver_id, type, content) values (?,?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(msg.ID, msg.SenderId, msg.ReceiverId, msg.Type, msg.Content); err != nil {
		return err
	}
	return nil
}

func (repo *MsgRepository) SaveGroupMsg(msg models.ChatMessage) error {
	stmt, err := repo.DB.Prepare("INSERT INTO group_messages (message_id, receiver_id) values (?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(msg.ID, msg.ReceiverId); err != nil {
		return err
	}
	return nil
}

// needs RECEIVER and SENDER as input
func (repo *MsgRepository) GetAll(msgIn models.ChatMessage) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	rows, err := repo.DB.Query("SELECT message_id,sender_id, receiver_id, type, content FROM messages WHERE (receiver_id = ? AND sender_id = ? )OR (receiver_id = ? AND sender_id = ? ) ORDER BY created_at ASC;", msgIn.ReceiverId, msgIn.SenderId, msgIn.SenderId, msgIn.ReceiverId)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		var msg models.ChatMessage
		rows.Scan(&msg.ID, &msg.SenderId, &msg.ReceiverId, &msg.Type, &msg.Content)
		messages = append(messages, msg)
	}
	return messages, nil
}

func (repo *MsgRepository) GetAllGroup(userId, groupId string) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	rows, err := repo.DB.Query("SELECT message_id,sender_id, receiver_id, type, content FROM messages WHERE (sender_id = ? AND receiver_id = ? ) OR (receiver_id = ? AND ((SELECT COUNT() FROM groups WHERE group_id = ? AND administrator = ?) = 1 OR (SELECT COUNT() FROM group_users WHERE group_id =? AND user_id =?) = 1) ) ORDER BY created_at ASC;", userId, groupId, groupId, groupId, userId, groupId, userId)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		var msg models.ChatMessage
		rows.Scan(&msg.ID, &msg.SenderId, &msg.ReceiverId, &msg.Type, &msg.Content)
		messages = append(messages, msg)
	}
	return messages, nil
}

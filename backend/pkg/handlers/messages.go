package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
	ws "social-network/pkg/wsServer"
)

// get all previous messages for chat
// waits for POST request with RECEIVER as target and TYPE
// respondes with all messages through simple http response
func (handler *Handler) Messages(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	/* ------------------- // get incoming data in msg format ------------------- */
	var msgIn models.ChatMessage
	err := json.NewDecoder(r.Body).Decode(&msgIn)
	if err != nil {
		utils.RespondWithError(w, "Error on reading the incomming message", 200)
		return
	}
	msgIn.SenderId = r.Context().Value(utils.UserKey).(string)

	/* ----------------------- get massages form database ----------------------- */
	var messages []models.ChatMessage
	if msgIn.Type == "PERSON" {
		messages, err = handler.repos.MsgRepo.GetAll(msgIn)
		if err != nil {
			utils.RespondWithError(w, "Error on getting the messages", 200)
			return
		}
	} else if msgIn.Type == "GROUP" {
		messages, err = handler.repos.MsgRepo.GetAllGroup(msgIn.SenderId, msgIn.ReceiverId)
		if err != nil {
			utils.RespondWithError(w, "Error on getting the messages", 200)
			return
		}
	}
	/* --------------------------- attach sender data --------------------------- */
	for i := 0; i < len(messages); i++ {
		messages[i].Sender, _ = handler.repos.UserRepo.GetDataMin(messages[i].SenderId)
	}

	utils.RespondWithMessages(w, messages, 200)
}

//new chat message wits for POST requet with SENDER, RECEIVER AND TYPE
// function saves new message and responds
// to sender with regular http response
// to recievers through websocket connection
func (handler *Handler) NewMessage(wsServer *ws.Server, w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	/* --------------------------- read incoming data --------------------------- */
	var msg models.ChatMessage
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		utils.RespondWithError(w, "Error on reading the incomming message", 200)
		return
	}

	/* -------------------- attach sender id and generate new id ------------------- */
	msg.SenderId = r.Context().Value(utils.UserKey).(string)
	msg.ID = utils.UniqueId()
	/* ---------------------------- save in database ---------------------------- */
	err = handler.repos.MsgRepo.Save(msg)
	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, "Error on saving message", 200)
		return
	}
	/* --------------------------- attach sender  info -------------------------- */
	msg.Sender, _ = handler.repos.UserRepo.GetDataMin(msg.SenderId)
	// send message respond with new message to sender
	utils.RespondWithMessages(w, []models.ChatMessage{msg}, 200)

	/* ------------------ respond through websocket to receiver ----------------- */
	if msg.Type == "PERSON" {
		for client := range wsServer.Clients {
			if client.ID == msg.ReceiverId {
				client.SendChatMessage(msg)
			}
		}
	} else if msg.Type == "GROUP" { //incase of group find and respond to all recievers
		// find all group members + admin except the sender
		allMembers, err := handler.repos.GroupRepo.GetMembers(msg.ReceiverId)
		if err != nil {
			utils.RespondWithError(w, "Error on geting group members", 200)
			return
		}
		for i := 0; i < len(allMembers); i++ {
			if allMembers[i].ID != msg.SenderId {
				/* -------- save also in group_messages table -------- */
				err = handler.repos.MsgRepo.SaveGroupMsg(models.ChatMessage{ID: msg.ID, ReceiverId: allMembers[i].ID})
			}
			for client := range wsServer.Clients {
				if client.ID == msg.SenderId {
					continue
				}
				if client.ID == allMembers[i].ID {
					msg.GroupUserReceiverId = allMembers[i].ID
					client.SendChatMessage(msg)
				}
			}

		}
	}
}

package main

import (
    "fmt"

    "github.com/Arman92/go-tdlib"
)

// Send Online status request to keep your telegram always online
func (f9 *Client) AlwaysOnline() {
    f9.client.SetOption("online", tdlib.NewOptionValueBoolean(true))
}

// Telegram Client will send an offline request to the server when you closed the telegram client.
// Then the server will acknowledge you that your status as offline.
// This feature is to track the server response.
// If the response is your offline status acknowledge, send online request to the server.
func (f9 *Client) StatusHook() {
    fmt.Println("[F9Hook] Offline Status hooked.")
    eventFilter := func(msg *tdlib.TdMessage) bool {
        updateMsg := (*msg).(*tdlib.UpdateUserStatus)
        if updateMsg.UserID == f9.UID && updateMsg.Status.GetUserStatusEnum() == "userStatusOffline" {
            return true
        }
        return false
    }

    receiver := f9.client.AddEventReceiver(&tdlib.UpdateUserStatus{}, eventFilter, 100)
    for range receiver.Chan {
        f9.AlwaysOnline()
    }
}

// Edit the message with "✔️✔️" when the recipient has read the message.
// The edit time refers to when the recipient read the message.
func (f9 *Client) MessageHook() {
    fmt.Println("[F9HOOK] Outgoing Message hooked.")
    eventFilter := func(msg *tdlib.TdMessage) bool {
        updateMsg := (*msg).(*tdlib.UpdateChatReadOutbox)
        chatID := updateMsg.ChatID
        // Ensure a user is not a group, channel.
        if chatID > 0 {
            return true
        }
        return false
    }

    receiver := f9.client.AddEventReceiver(&tdlib.UpdateChatReadOutbox{}, eventFilter, 100)
    for newMsg := range receiver.Chan {
        updateMsg := (newMsg).(*tdlib.UpdateChatReadOutbox)
        chatID := updateMsg.ChatID
        u, err := f9.client.GetUser(int32(chatID))
        if err != nil {
            fmt.Println(err.Error())
            continue
        }
        // Ensure a user is not a bot.
        if u.Type.GetUserTypeEnum() != "userTypeRegular" {
            continue
        }
        msgID := updateMsg.LastReadOutboxMessageID
        m, err := f9.client.GetMessage(chatID, msgID)
        if err != nil {
            fmt.Println(err.Error())
            continue
        }
        mType := m.Content.GetMessageContentEnum()
        if mType == "messageText" {
            msgText := m.Content.(*tdlib.MessageText).Text.Text
            msgEnt := m.Content.(*tdlib.MessageText).Text.Entities
            inputMsg := tdlib.NewInputMessageText(tdlib.NewFormattedText(msgText+"\n✔️✔️", msgEnt), false, true)
            f9.client.EditMessageText(chatID, msgID, nil, inputMsg)
        }
        if mType == "messagePhoto" {
            msgText := m.Content.(*tdlib.MessagePhoto).Caption.Text
            msgEnt := m.Content.(*tdlib.MessagePhoto).Caption.Entities
            inputMsg := tdlib.NewFormattedText(msgText+"\n✔️✔️", msgEnt)
            f9.client.EditMessageCaption(chatID, msgID, nil, inputMsg)
        }
    }
}

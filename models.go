package goTelegramBot


type Update struct {
	UpdateId int 					`json:"update_id"`
	Message Message					`json:"message"`
}

type Message struct {
	Chat Chat						`json:"chat"`
	Text string						`json:"text"`
	UserName User					`json:"from"`
}

type Chat struct {
	ChatId int						`json:"id"`
}

type RestResponse struct {
	Result []Update					`json:"result"`
}

type BotMessage struct {
	ChatId int						`json:"chat_id"`
	Text string						`json:"text"`
}

type User struct {
	UserId int 						`json:"id"`
	IsBot bool						`json:"is_bot"`
	FistName string					`json:"first_name"`
	SecondName string				`json:"last_name"`
	UserName string					`json:"username"`
	LanguageCode string				`json:"language_code"`
	CanJoinGroups bool				`json:"can_join_groups"`
	CanReadAllGroupMessage bool		`json:"can_read_all_group_messages"`
	SupportsInlineQueries bool		`json:"supports_inline_queries"`
}

package goTelegramBot


type Update struct {
	UpdateId int 					`json:"update_id"`
	Message Message					`json:"message"`
	EditMessage Message				`json:"edited_message"`
}

type BotCommand struct {
	Command string					`json:"command"`					// Текст команды, 1-32 символа. Может содержать только строчные английские буквы, цифры и подчеркивания.
	Description string				`json:"description"`				// Описание команды, 3-256 символов
}

type Message struct {
	Chat Chat							`json:"chat"`
	Text string							`json:"text"`
	UserInfo User						`json:"from"`
	CaptionEntities []MessageEntity		`json:"caption_entities"`
}

type MessageEntity struct {
	Type string 					`json:"type"`
	Offset int 						`json:"offset"`
	Length int						`json:"length"`
	Url string						`json:"url"`
	UserInfo User 					`json:"user"`
	Language string					`json:"language"`
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

type ReplyKeyboardMarkup struct {
	Keyboard        	[][]KeyboardButton		`json:"keyboard"`
	ResizeKeyboard  	bool               		`json:"resize_keyboard"`   // optional
	OneTimeKeyboard 	bool               		`json:"one_time_keyboard"` // optional
	Selective       	bool               		`json:"selective"`         // optional
}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

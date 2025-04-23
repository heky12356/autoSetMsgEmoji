package model

type Message struct {
	Action string        `json:"action"`
	Params MessageParams `json:"params"`
	Echo   string        `json:"echo"`
}

type MessageParams struct {
	Group_id int64  `json:"group_id"`
	Message  string `json:"message"`
}

type eMessageParams struct {
	Message_id int32 `json:"message_id"`
}

type gMessageParams struct {
	Group_id int64  `json:"group_id"`
	Content  string `json:"content"`
}

type EssenceMessage struct {
	Action string         `json:"action"`
	Params eMessageParams `json:"params"`
}

type Group_notice struct {
	Action string         `json:"action"`
	Params gMessageParams `json:"params"`
	Echo   string         `json:"echo"`
}

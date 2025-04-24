package model

type Message struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
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

type SetMsgEmoji struct {
	Message_id int32 `json:"message_id"`
	Emoji_id   int32 `json:"emoji_id"`
	Set        bool  `json:"set"`
}

// 上报消息

type Sender struct {
	User_id  int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Role     string `json:"role"`
}

type ReMessage struct {
	Type string `json:"type"`
	Data struct {
		Text string `json:"text"`
	} `json:"data"`
}

type Response struct {
	Self_id        int64  `json:"self_id"`
	User_id        int64  `json:"user_id"`
	Time           int64  `json:"time"`
	Message_id     int64  `json:"message_id"`
	Message_seq    int64  `json:"message_seq"`
	Real_id        int64  `json:"real_id"`
	Real_seq       string `json:"real_seq"`
	Message_type   string `json:"message_type"`
	Sender         Sender `json:"sender"`
	Raw_message    string `json:"raw_message"`
	Font           int64  `json:"font"`
	Subtype        string `json:"sub_type"`
	Message        string `json:"message"`
	Message_format string `json:"message_format"`
	Post_type      string `json:"post_type"`
	Group_id       int64  `json:"group_id"`
}

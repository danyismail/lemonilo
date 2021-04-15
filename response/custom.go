package response

type Response struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type Users struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

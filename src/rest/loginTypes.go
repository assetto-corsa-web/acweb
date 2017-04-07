package rest

type loginRes struct {
	Id int64 `json:"user_id"`
}

type loginReq struct {
	Login string `json:"login"`
	Pwd   string `json:"pwd"`
}

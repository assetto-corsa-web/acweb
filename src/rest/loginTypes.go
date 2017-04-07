package rest

type loginRes struct {
	Id int64 `json:"user_id"`
}

type loginReq struct {
	Login string `json:"login"`
	Pwd   string `json:"pwd"`
}

type removeUserReq struct {
	Id int64 `json:"id"`
}

type addEditUserReq struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Pwd1  string `json:"pwd1"`
	Pwd2  string `json:"pwd2"`
}

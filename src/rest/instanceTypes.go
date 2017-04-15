package rest

type startInstanceReq struct {
	Name          string `json:"name"`
	Configuration int64  `json:"config"`
}

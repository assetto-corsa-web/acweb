package rest

type saveSettingsReq struct {
	Folder string `json:"folder"`
	Cmd    string `json:"cmd"`
}

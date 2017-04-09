package rest

type saveSettingsReq struct {
	Folder     string `json:"folder"`
	Executable string `json:"executable"`
	Args       string `json:"args"`
}

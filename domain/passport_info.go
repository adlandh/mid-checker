package domain

type PassportStatus struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Color        string `json:"color"`
	Subscription bool   `json:"subscription"`
}

type InternalStatus struct {
	Name    string `json:"name"`
	Percent int    `json:"percent"`
}

type PassportInfo struct {
	Uid            string         `json:"uid"`
	ReceptionDate  string         `json:"receptionDate"`
	PassportStatus PassportStatus `json:"passportStatus"`
	InternalStatus InternalStatus `json:"internalStatus"`
}

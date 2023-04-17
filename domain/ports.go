package domain

type PassportInfoFetcher interface {
	GetPassportStatus() (*PassportInfo, error)
}

type PassportChangedEventSender interface {
	SendChangedStatus(*PassportInfo) error
}

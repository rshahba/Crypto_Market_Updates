package client

type ResponseFormat interface {
	GetUrlStr() (string, error)
}

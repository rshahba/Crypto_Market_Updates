package client

type ResponseFormat interface {
	GetUrlStr(url string) (string, error)
}
package redisutil

type Redis interface {
	Read(string) (string, error)
	Write(string, string) error
}

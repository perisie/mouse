package mouse

type Mouse interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
}

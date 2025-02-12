package storage

type Storage interface {
	Save(short, original string) error
	Get(short string) (string, bool, error)
}

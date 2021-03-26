package factory

type (
	Storage interface {
		Store([]byte) error
	}
)

// S3実装
type s3Storage struct{}

func (s s3Storage) Store(data []byte) error {
	// TODO Implement this func
	return nil
}

// Disk実装
type diskStorage struct{}

func (s diskStorage) Store(data []byte) error {
	// TODO Implement this func
	return nil
}

// Memory実装
type memoryStorage struct{}

func (s memoryStorage) Store(data []byte) error {
	// TODO Implement this func
	return nil
}

type NopeStorage struct{}

func (s NopeStorage) Store(data []byte) error {
	// TODO Implement this func
	return nil
}

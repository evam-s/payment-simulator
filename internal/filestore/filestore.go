package filestore

import (
	"encoding/json"
	"os"
	"sync"
)

type Payment struct {
	Id     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type FileStore struct {
	mu   sync.Mutex
	path string
}

func NewFileStore(path string) *FileStore {
	return &FileStore{path: path}
}

func (fs *FileStore) Save(payments []Payment) error {

	if previousPayments, err1 := fs.Load(); err1 != nil {
		return err1
	} else {
		payments = append(payments, previousPayments...)
	}
	fs.mu.Lock()
	defer fs.mu.Unlock()
	data, err := json.MarshalIndent(payments, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.path, data, 0644)
}

func (fs *FileStore) Load() ([]Payment, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	if _, err := os.Stat(fs.path); os.IsNotExist(err) {
		return []Payment{}, nil
	}

	data, err := os.ReadFile(fs.path)
	if err != nil {
		return nil, err
	}

	var payments []Payment
	if err := json.Unmarshal(data, &payments); err != nil {
		return nil, err
	}
	return payments, nil
}

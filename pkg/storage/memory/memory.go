package memory

import (
	"sync"
	"tx_parser/pkg/parser"
)

func init() {
	sd := &memoryStorage{sync.RWMutex{}, make(map[string][]parser.Transaction)}
	parser.SetStorage(sd)
}

type memoryStorage struct {
	lock sync.RWMutex
	data map[string][]parser.Transaction
}

func (s *memoryStorage) Insert(address string, transaction parser.Transaction) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data[address] = append(s.data[address], transaction)
	return nil
}

func (s *memoryStorage) Get(address string) ([]parser.Transaction, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.data[address], nil
}

package parser

import "errors"

type Storage interface {
	Insert(address string, transaction Transaction) error
	Get(address string) ([]Transaction, error)
}

var storage Storage = nil

func SetStorage(sd Storage) {
	if storage != nil {
		panic("storage is already initialized")
	}
	storage = sd
}

func checkInitialized() error {
	if storage == nil {
		return errors.New("storage must be initialized")
	}
	return nil
}

func insert(address string, transaction Transaction) error {
	if err := checkInitialized(); err != nil {
		return err
	}
	return storage.Insert(address, transaction)
}

func get(address string) ([]Transaction, error) {
	if err := checkInitialized(); err != nil {
		return []Transaction{}, err
	}
	return storage.Get(address)
}

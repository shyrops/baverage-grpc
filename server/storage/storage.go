package storage

import (
	"context"
	"log"

	pb "beverage-grpc/pkg"
)

type BeverageStorage interface {
	Add(ctx context.Context, beverage *pb.Beverage) (*pb.Beverage, error)
	List(ctx context.Context) (*pb.BeverageList, error)
}

func NewStorage(storageType string) BeverageStorage {
	switch storageType {
	case memoryStorageTypeName:
		return NewMemoryStorage()
	case fileStorageTypeName:
		return NewFileStorage()
	case sqlStorageTypeName:
		return NewSqlStorage()
	default:
		log.Fatalf("incorrect storage type provided: %s", storageType)
	}

	return nil
}

package storage

import (
	"context"
	"errors"
	"io/fs"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"

	pb "beverage-grpc/pkg"
)

const fileStorageTypeName = "file"

type file struct{}

func NewFileStorage() *file {
	return &file{}
}

func (f *file) Add(ctx context.Context, beverage *pb.Beverage) (*pb.Beverage, error) {
	if beverage == nil {
		log.Fatalf("file add: nil bevrage provided")
	}

	var beveragesList *pb.BeverageList = &pb.BeverageList{}

	readedBytes, err := os.ReadFile("files/beverages.json")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			log.Print("File beverages.json was not found. Creating a new file")
			beveragesList.Beverages = append(beveragesList.Beverages, beverage)
			jsonBytes, err := protojson.Marshal(beveragesList)
			if err != nil {
				log.Fatalf("JSON marshaling failed: %v", err)
			}
			if err = os.WriteFile("files/beverages.json", jsonBytes, 0644); err != nil {
				log.Fatalf("error writing to file: %v", err)
			}
		} else {
			log.Fatalf("error reading file: %v", err)
		}
	}

	if err = protojson.Unmarshal(readedBytes, beveragesList); err != nil {
		log.Fatalf("failed to parse beverages list: %v", err)
	}
	beveragesList.Beverages = append(beveragesList.Beverages, beverage)
	jsonBytes, err := protojson.Marshal(beveragesList)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %v", err)
	}
	if err = os.WriteFile("files/beverages.json", jsonBytes, 0644); err != nil {
		log.Fatalf("error writing to file: %v", err)
	}

	return beverage, nil
}

func (f *file) List(ctx context.Context) (*pb.BeverageList, error) {
	jsonBytes, err := os.ReadFile("server/beverages.json")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var beveragesList *pb.BeverageList = &pb.BeverageList{}
	if err = protojson.Unmarshal(jsonBytes, beveragesList); err != nil {
		log.Fatalf("failed to parse beverages list: %v", err)
	}

	return beveragesList, nil
}

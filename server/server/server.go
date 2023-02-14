package server

import (
	"context"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "beverage-grpc/pkg"
	"beverage-grpc/server/storage"
)

const (
	port = ":50051"
)

func NewServer() *beverageManagementServer {
	return &beverageManagementServer{
		storage: storage.NewStorage("file"),
	}
}

type beverageManagementServer struct {
	pb.UnimplementedBeveragesManagementServer
	storage storage.Storage
}

func (b *beverageManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to init listener: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBeveragesManagementServer(s, b)
	log.Printf("server listening on: %s", lis.Addr())

	return s.Serve(lis)
}

func (b *beverageManagementServer) CreateBeverage(ctx context.Context, req *pb.CreateBeverageRequest) (bev *pb.Beverage, err error) {
	attr := req.GetAttr()
	log.Printf(
		"Start creating new beverage with name: %s, type: %s, volume: %d",
		attr.GetName(),
		attr.GetType(),
		attr.GetVolume(),
	)

	price, err := calculatePrice(req.GetAttr())
	if err != nil {
		return bev, err
	}

	bev = &pb.Beverage{
		Id:    rand.Uint64(),
		Attr:  attr,
		Price: price,
	}

	return b.storage.StoreBeverage(ctx, bev)
}

func (b *beverageManagementServer) GetBeverages(
	ctx context.Context,
	req *pb.GetBeveragesParams,
) (*pb.BeverageList, error) {
	return b.storage.GetBeverages(ctx)
}

func calculatePrice(attr *pb.BeverageMainAttributes) (int32, error) {
	t := attr.GetType()
	if t == pb.BeverageType_BEVERAGE_TYPE_UNSPECIFIED {
		return 0, status.Error(codes.InvalidArgument, "beverage type unspecified")
	}

	oneMlPrice := map[pb.BeverageType]int32{
		pb.BeverageType_BEVERAGE_TYPE_BEER:  30,
		pb.BeverageType_BEVERAGE_TYPE_VINE:  45,
		pb.BeverageType_BEVERAGE_TYPE_VODKA: 55,
		pb.BeverageType_BEVERAGE_TYPE_GIN:   75,
		pb.BeverageType_BEVERAGE_TYPE_RUM:   70,
	}

	if price, ok := oneMlPrice[t]; !ok {
		return price, status.Error(codes.InvalidArgument, "unsupported beverage type")
	}

	return oneMlPrice[t] * attr.GetVolume(), nil
}

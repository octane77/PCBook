package service

import (
	"PCBook/pb/PCBook/proto"
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// LaptopServer Laptop Server is the server that provides laptop services
type LaptopServer struct {
	Store LaptopStorage
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer(store LaptopStorage) *LaptopServer {
	return &LaptopServer{store}
}

// CreateLaptop is a unary RPC to create a new Laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, request *proto.CreateLaptopRequest) (*proto.CreateLaptopResponse, error) {
	laptop := request.GetLaptop()
	log.Printf("Received CreateLaptop request with ID: %s\n", laptop.Id)

	if len(laptop.Id) > 0 {
		//Check if it's valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Laptop ID %s: %v", laptop.Id, err)
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Unable to generate laptop ID: %v", err)
			}
			laptop.Id = id.String()
		}
	}
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "Error saving laptop: %v", err)
	}
	log.Printf("Saved laptop with id: %s", laptop.Id)
	response := &proto.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return response, nil
}

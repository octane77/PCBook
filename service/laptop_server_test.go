package service_test

import (
	"PCBook/pb/PCBook/proto"
	"PCBook/sample"
	"PCBook/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestServerCreateLaptop(test *testing.T) {
	test.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.Nil(test, err)

	testCases := []struct {
		name   string
		laptop *proto.Laptop
		store  service.LaptopStorage
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failed_invalid_id",
			laptop: laptopInvalidID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failed_duplicate_id",
			laptop: laptopDuplicateID,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		test.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			request := &proto.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := service.NewLaptopServer()
		})
	}
}

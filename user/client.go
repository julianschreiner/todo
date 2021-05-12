package user

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "todo"
)

const NormalUserId = 3
const rpcDial = "user-uas:5040"

type UserClient interface {
	GetActiveUserById(id uint64) (*User, error)
}

type userClient struct {
	client pb.UserClient
}

func NewUserClient() UserClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(rpcDial, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	log.Printf("grpc onnection %f", conn)
	//defer conn.Close()

	client := pb.NewUserClient(conn)

	return &userClient{
		client: client,
	}
}

func (u *userClient) GetActiveUserById(id uint64) (*User, error) {
	response, err := u.client.GetUserInformation(context.Background(), &pb.GetUserInformationRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("Error when calling GetUserInformation: %s", err)
		return nil, err
	}

	log.Printf("[GET per ID] Response from server: %s", response)
	if response.Size() > 0 {
		if response.User.Size() > 0 {
			user := User{
				Id:       uint(response.User.Id),
				Email:    response.User.Email,
				Role:     uint(response.User.Role),
				Dob:      response.User.Dob,
				Active:   response.User.Active,
				Forename: response.User.Forename,
				Surname:  response.User.Surname,
			}
			return &user, nil
		}
	}

	return nil, nil
}

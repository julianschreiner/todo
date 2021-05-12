package auth

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "todo"
)

const NormalAuthId = 3
const rpcDial = "auth-uas:5040"

type AuthClient interface {
	Validate(refresh string) error
}

type authClient struct {
	client pb.AuthenticationClient
}

func NewauthClient() AuthClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(rpcDial, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	log.Printf("grpc connection %f", conn)
	//defer conn.Close()

	client := pb.NewAuthenticationClient(conn)

	return &authClient{
		client: client,
	}
}

func (u *authClient) Validate(refresh string) error {
	response, err := u.client.Refresh(context.Background(), &pb.RefreshRequest{
		Refresh: refresh,
	})
	if err != nil {
		log.Printf("Error when calling GetAuthInformation: %s", err)
		return err
	}

	log.Printf("[GET REFRESH] Response from server: %s", response)
	if response.Size() > 0 {
		println(response.Refresh)
	}

	return nil
}

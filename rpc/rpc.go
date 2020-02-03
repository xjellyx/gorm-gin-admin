package userRpc

import (
	"context"
	pb "github.com/olefen/model.grpc"
	"github.com/olefen/userDetail/model"
	"google.golang.org/grpc"
)

// ServeRpc
type ServeRpc struct {
	pb.UnimplementedUserBaseServer
	cc     *grpc.ClientConn
	client pb.UserBaseClient
}

func (s *ServeRpc) GetUserToken(ctx context.Context, arg *pb.ArgLogin) (ret *pb.TokenGetRes, err error) {
	var (
		token string
	)
	if token, err = model.GetUserToken(arg); err != nil {
		return
	}

	//
	ret = new(pb.TokenGetRes)
	ret.Token = token
	return
}

func NewRpc(addr string) (ret *ServeRpc, err error) {
	var (
		cc *grpc.ClientConn
	)
	cc, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	ret = &ServeRpc{cc: cc, client: pb.NewUserBaseClient(cc)}
	return
}

func (s *ServeRpc) Close() {
	s.cc.Close()
}

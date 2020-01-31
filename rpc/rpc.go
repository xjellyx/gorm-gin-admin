package userRpc

import (
	pb "github.com/srlemon/model.grpc"
	"google.golang.org/grpc"
)

// ServeRpc
type ServeRpc struct {
	cc     *grpc.ClientConn
	client pb.UserBaseClient
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

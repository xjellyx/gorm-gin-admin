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

func transportUserDetailToPb(u *model.UserDetail, pb *pb.UserDetailResp) {
	pb.Uid = u.Uid
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

func (s *ServeRpc) GetUserDetail(ctx context.Context, arg *pb.UserDetailReq) (ret *pb.UserDetailResp, err error) {
	var (
		data *model.UserDetail
	)
	if len(arg.GetUid()) > 0 {
		if data, err = model.PubUserGet(arg.GetUid()); err != nil {
			return
		}
	} else if len(arg.GetPhone()) > 0 {
		if data, err = model.PubUserGetByPhone(arg.GetLocNum(), arg.GetPhone()); err != nil {
			return
		}
	} else if len(arg.GetEmail()) > 0 {
		if data, err = model.PubUserGetByEmail(arg.GetEmail()); err != nil {
			return
		}
	} else if len(arg.GetUsername()) > 0 {
		if data, err = model.PubUserGetByUsername(arg.GetUsername()); err != nil {
			return
		}
	}

	ret = new(pb.UserDetailResp)
	transportUserDetailToPb(data, ret)
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

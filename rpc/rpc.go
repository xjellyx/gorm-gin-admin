package userRpc

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/olongfen/contrib"
	pb "github.com/olongfen/model.grpc"
	"github.com/olongfen/userDetail/model"
	"google.golang.org/grpc"
)

// ServeRpc
type ServeRpc struct {
	pb.UnimplementedUserBaseServer
	cc     *grpc.ClientConn
	client pb.UserBaseClient
}

func transportUserDetailToPb(u *model.UserDetail, val *pb.UserDetailResp) {
	val.Uid = u.Uid
	val.Status = int32(u.Status)
	val.Verified = u.Verified
	val.IsChangeUsername = u.IsChangeUsername
	val.Nickname = u.Nickname
	val.Username = u.Username
	val.RealName = u.RealName
	val.Sign = u.Sign
	val.HeadIcon = u.HeadIcon
	val.LocNum = u.LocNum
	val.Phone = u.Phone
	val.IsAdmin = u.IsAdmin
	val.Email = u.Email
	val.Level = u.Level
	val.TimeData = new(pb.TimeData)
	val.TimeData.CreatedAt, _ = ptypes.TimestampProto(u.CreatedAt)
	val.TimeData.UpdatedAt, _ = ptypes.TimestampProto(u.UpdatedAt)
	for _, v := range u.BankCards {
		_b := new(pb.BankCard)
		transportBankCard(v, _b)
		val.Cards = append(val.Cards, _b)
	}

	for _, v := range u.Addr {
		_b := new(pb.AddressDetail)
		transportAddr(v, _b)

		val.Addr = append(val.Addr, _b)
	}

}

func transportBankCard(u *model.BankCard, v *pb.BankCard) {
	v.Number = u.Number
	v.Name = u.Name
	v.BankStart = u.BankStart
	v.Bank = u.Bank
}

func transportAddr(u *model.AddressDetail, v *pb.AddressDetail) {
	v.Address = u.Address
	v.Province = u.Province
	v.Country = u.Country
	v.City = u.City
	v.District = u.District
}

func (s *ServeRpc) GetUserToken(ctx context.Context, arg *pb.ArgLogin) (ret *pb.TokenGetRes, err error) {
	var (
		token string
		uid   string
	)
	if token, uid, err = model.GetUserToken(arg); err != nil {
		return
	}

	//
	ret = new(pb.TokenGetRes)
	ret.Token = token
	ret.Uid = uid
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
	} else {
		err = contrib.ErrParamInvalid
		return
	}

	ret = new(pb.UserDetailResp)
	transportUserDetailToPb(data, ret)
	return
}

func (s *ServeRpc) CheckToken(ctx context.Context, arg *pb.CheckTokenReq) (ret *pb.CheckTokenRes, err error) {
	if err = model.TokenCheck(arg.Token, arg.IsAdmin); err != nil {
		return
	}

	ret = new(pb.CheckTokenRes)
	return
}

func (s *ServeRpc) UpdateUserDetail(ctx context.Context, req *pb.UpdateUserDetailReq) (ret *pb.UserDetailResp, err error) {
	return nil, nil
}
func (s *ServeRpc) GetUserIDCard(ctx context.Context, req *pb.GetUserIDCardReq) (ret *pb.UserIDCardResp, err error) {
	return
}
func (s *ServeRpc) AddUserBankCard(ctx context.Context, req *pb.AddBankCardReq) (ret *pb.BankCardResp, err error) {
	return
}
func (s *ServeRpc) GetUserBankCardList(ctx context.Context, req *pb.GetUserBankCardReq) (ret *pb.GetUserBankCardResp, err error) {
	return
}
func (s *ServeRpc) DeleteUserBankCard(ctx context.Context, req *pb.DelUserBankCardReq) (ret *pb.PubNoneResp, err error) {
	return
}
func (s *ServeRpc) AddUserAddress(ctx context.Context, req *pb.AddUserAddrReq) (ret *pb.UserAddrResp, err error) {
	return
}
func (s *ServeRpc) GetUserAddressList(ctx context.Context, req *pb.GetUserAddressReq) (ret *pb.GetUserAddressResp, err error) {
	return
}
func (s *ServeRpc) UpdateUserAddress(ctx context.Context, req *pb.UpdateUserAddrReq) (ret *pb.UserAddrResp, err error) {
	return
}
func (s *ServeRpc) DeleteUserAddress(ctx context.Context, req *pb.DelUserAddrReq) (ret *pb.PubNoneResp, err error) {
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

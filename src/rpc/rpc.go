package rpc

//package userRpc
//
//import (
//	"context"
//	"github.com/golang/protobuf/ptypes"
//	"github.com/olongfen/contrib"
//	"github.com/olongfen/gorm-gin-admin/src/models"
//	"github.com/olongfen/gorm-gin-admin/src/utils"
//	pb "github.com/olongfen/models.grpc"
//	"google.golang.org/grpc"
//)
//
//// ServeRpc
//type ServeRpc struct {
//	pb.UnimplementedUserBaseServer
//	cc     *grpc.ClientConn
//	client pb.UserBaseClient
//}
//
//func transportUserDetailToPb(u *models.UserDetail, val *pb.UserDetailResp) {
//	val.Uid = u.Uid
//	val.Status = int32(u.Status)
//	val.Verified = u.Verified
//	val.IsChangeUsername = u.IsChangeUsername
//	val.Nickname = u.Nickname
//	val.Username = u.Username
//	val.RealName = u.RealName
//	val.Sign = u.Sign
//	val.HeadIcon = u.HeadIcon
//	val.LocNum = u.LocNum
//	val.Phone = u.Phone
//	val.IsAdmin = u.IsAdmin
//	val.Email = u.Email
//	val.TimeData = new(pb.TimeData)
//	val.TimeData.CreatedAt, _ = ptypes.TimestampProto(u.CreatedAt)
//	val.TimeData.UpdatedAt, _ = ptypes.TimestampProto(u.UpdatedAt)
//	for _, v := range u.BankCards {
//		_b := new(pb.BankCard)
//		transportBankCard(v, _b)
//		val.Cards = append(val.Cards, _b)
//	}
//
//	for _, v := range u.Addr {
//		_b := new(pb.AddressDetail)
//		transportAddr(v, _b)
//
//		val.Addr = append(val.Addr, _b)
//	}
//
//}
//
//func transportBankCard(u *models.BankCard, v *pb.BankCard) {
//	v.Number = u.Number
//	v.Name = u.Name
//	v.BankStart = u.BankStart
//	v.Bank = u.Bank
//}
//
//func transportAddr(u *models.AddressDetail, v *pb.AddressDetail) {
//	v.Address = u.Address
//	v.Province = u.Province
//	v.Country = u.Country
//	v.City = u.City
//	v.District = u.District
//}
//
//func transportIDCard(u *models.IDCard, v *pb.UserIDCardResp) {
//	v.Name = u.Name
//	v.Sex = int32(u.Sex)
//	v.Nation = u.Nation
//	v.ValidPeriod = u.ValidPeriod
//	v.Birthday = u.Birthday
//	v.IssueOrg = u.IssueOrg
//	v.IdCard = u.IDCard
//	v.IdCardAddr = u.IDCardAddr
//}
//func (s *ServeRpc) GetUserToken(ctx context.Context, arg *pb.ArgLogin) (ret *pb.TokenGetRes, err error) {
//	var (
//		token string
//		uid   string
//	)
//
//	if token, uid, err = models.GetUserToken(arg); err != nil {
//		return
//	}
//
//	//
//	ret = new(pb.TokenGetRes)
//	ret.Token = token
//	ret.Uid = uid
//	return
//}
//
//func (s *ServeRpc) GetUserDetail(ctx context.Context, arg *pb.UserDetailReq) (ret *pb.UserDetailResp, err error) {
//	var (
//		data *models.UserDetail
//	)
//	if len(arg.GetUid()) > 0 {
//		if data, err = models.PubUserGet(arg.GetUid()); err != nil {
//			return
//		}
//	} else if len(arg.GetPhone()) > 0 {
//		if data, err = models.PubUserGetByPhone(arg.GetLocNum(), arg.GetPhone()); err != nil {
//			return
//		}
//	} else if len(arg.GetEmail()) > 0 {
//		if data, err = models.PubUserGetByEmail(arg.GetEmail()); err != nil {
//			return
//		}
//	} else if len(arg.GetUsername()) > 0 {
//		if data, err = models.PubUserGetByUsername(arg.GetUsername()); err != nil {
//			return
//		}
//	} else {
//		err = contrib.ErrParamInvalid
//		return
//	}
//
//	ret = new(pb.UserDetailResp)
//	transportUserDetailToPb(data, ret)
//	return
//}
//
//func (s *ServeRpc) CheckToken(ctx context.Context, arg *pb.CheckTokenReq) (ret *pb.CheckTokenRes, err error) {
//	if err = models.TokenCheck(arg.Token, arg.IsAdmin); err != nil {
//		return
//	}
//
//	ret = new(pb.CheckTokenRes)
//	return
//}
//
//func (s *ServeRpc) UpdateUserDetail(ctx context.Context, req *pb.UpdateUserDetailReq) (ret *pb.UserDetailResp, err error) {
//	var (
//		data *models.UserDetail
//		form = &utils.FormEditUser{
//			Nickname: nil,
//			Username: nil,
//			LocNum:   nil,
//			Phone:    nil,
//			HeadIcon: nil,
//			Sign:     nil,
//		}
//	)
//	if len(req.GetUsername()) > 0 {
//		n := req.GetUsername()
//		form.Username = &n
//	}
//	if len(req.GetNickname()) > 0 {
//		d := req.GetNickname()
//		form.Nickname = &d
//	}
//	if len(req.GetPhone()) > 0 {
//		d := req.GetPhone()
//		form.Phone = &d
//	}
//	if len(req.GetLocNum()) > 0 {
//		d := req.GetLocNum()
//		form.LocNum = &d
//	}
//	if len(req.GetHeadIcon()) > 0 {
//		d := req.GetHeadIcon()
//		form.HeadIcon = &d
//	}
//	if len(req.GetSign()) > 0 {
//		d := req.GetSign()
//		form.Sign = &d
//	}
//	if data, err = models.PubUserUpdate(req.Uid, form); err != nil {
//		return
//	}
//
//	//
//	ret = new(pb.UserDetailResp)
//	transportUserDetailToPb(data, ret)
//	return nil, nil
//}
//
//func (s *ServeRpc) GetUserIDCard(ctx context.Context, req *pb.GetUserIDCardReq) (ret *pb.UserIDCardResp, err error) {
//	var (
//		data *models.IDCard
//	)
//	if data, err = models.PubGetIDCard(req.Uid); err != nil {
//		return
//	}
//	ret = new(pb.UserIDCardResp)
//	transportIDCard(data, ret)
//	return
//}
//
//func (s *ServeRpc) AddUserBankCard(ctx context.Context, req *pb.AddBankCardReq) (ret *pb.BankCard, err error) {
//	var (
//		data *models.BankCard
//	)
//	if data, err = models.PubBankCardAdd(req.Uid, &utils.FormBankCard{
//		Number:    req.Number,
//		Name:      req.Name,
//		Bank:      req.Bank,
//		BankStart: req.BankStart,
//	}); err != nil {
//		return
//	}
//
//	ret = new(pb.BankCard)
//	transportBankCard(data, ret)
//	return
//}
//
//func (s *ServeRpc) GetUserBankCardList(ctx context.Context, req *pb.GetUserBankCardReq) (ret *pb.GetUserBankCardResp, err error) {
//	var (
//		data []*models.BankCard
//	)
//	if data, err = models.PubBankCardGetList(req.Uid); err != nil {
//		return
//	}
//
//	ret = new(pb.GetUserBankCardResp)
//	for _, v := range data {
//		d := new(pb.BankCard)
//		transportBankCard(v, d)
//		ret.Data = append(ret.Data, d)
//	}
//	return
//}
//
//func (s *ServeRpc) DeleteUserBankCard(ctx context.Context, req *pb.DelUserBankCardReq) (ret *pb.PubNoneResp, err error) {
//
//	if err = models.PubBankCardDel(req.Uid, req.Number); err != nil {
//		return
//	}
//
//	ret = new(pb.PubNoneResp)
//	return
//}
//
//func (s *ServeRpc) AddUserAddress(ctx context.Context, req *pb.AddUserAddrReq) (ret *pb.AddressDetail, err error) {
//	var (
//		data *models.AddressDetail
//		form = &utils.FormAddr{
//			Country:  req.Country,
//			Province: req.Province,
//			City:     req.City,
//			District: nil,
//			Address:  nil,
//		}
//	)
//	if len(req.GetDistrict()) > 0 {
//		d := req.GetDistrict()
//		form.District = &d
//	}
//	if len(req.GetAddress()) > 0 {
//		d := req.GetAddress()
//		form.Address = &d
//	}
//	if data, err = models.PubAddressAdd(req.Uid, form); err != nil {
//		return
//	}
//
//	ret = new(pb.AddressDetail)
//	transportAddr(data, ret)
//	return
//}
//
//func (s *ServeRpc) GetUserAddressList(ctx context.Context, req *pb.GetUserAddressReq) (ret *pb.GetUserAddressResp, err error) {
//	var (
//		data []*models.AddressDetail
//	)
//	if data, err = models.PubAddressGetList(req.Uid); err != nil {
//		return
//	}
//
//	ret = new(pb.GetUserAddressResp)
//	for _, v := range data {
//		d := new(pb.AddressDetail)
//		transportAddr(v, d)
//		ret.Data = append(ret.Data, d)
//	}
//	return
//}
//
//func (s *ServeRpc) UpdateUserAddress(ctx context.Context, req *pb.UpdateUserAddrReq) (ret *pb.AddressDetail, err error) {
//	var (
//		data *models.AddressDetail
//		form = &utils.FormAddr{
//			Country:  req.Country,
//			Province: req.Province,
//			City:     req.City,
//		}
//	)
//
//	if len(req.GetDistrict()) > 0 {
//		d := req.GetDistrict()
//		form.District = &d
//	}
//	if len(req.GetAddress()) > 0 {
//		d := req.GetAddress()
//		form.Address = &d
//	}
//	if data, err = models.PubAddressUpdate(req.Uid, form); err != nil {
//		return
//	}
//
//	ret = new(pb.AddressDetail)
//	transportAddr(data, ret)
//	return
//}
//func (s *ServeRpc) DeleteUserAddress(ctx context.Context, req *pb.DelUserAddrReq) (ret *pb.PubNoneResp, err error) {
//	if err = models.PubAddressDelete(req.Uid, req.Id); err != nil {
//		return
//	}
//	ret = new(pb.PubNoneResp)
//	return
//}
//
//func (s *ServeRpc) UserOffline(ctx context.Context, req *pb.UserOfflineReq) (ret *pb.PubNoneResp, err error) {
//	if err = models.UserOfflineDo(req.Uid); err != nil {
//		return
//	}
//	ret = new(pb.PubNoneResp)
//	return
//}
//
//func (s *ServeRpc) AddUserDetail(ctx context.Context, req *pb.ArgRegistry) (ret *pb.UserDetailResp, err error) {
//	var (
//		data *models.UserDetail
//	)
//	models.logModel.Infoln("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
//	if data, err = models.PubUserAdd(&utils.FormRegister{
//		Phone:    req.Phone,
//		Password: req.Password,
//		Code:     req.Code,
//	}); err != nil {
//		models.logModel.Infoln("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq", err)
//		return
//	}
//
//	ret = new(pb.UserDetailResp)
//	transportUserDetailToPb(data, ret)
//	return
//}
//
//func NewRpc(addr string) (ret *ServeRpc, err error) {
//	var (
//		cc *grpc.ClientConn
//	)
//	cc, err = grpc.Dial(addr, grpc.WithInsecure())
//	if err != nil {
//		return
//	}
//	ret = &ServeRpc{cc: cc, client: pb.NewUserBaseClient(cc)}
//	return
//}
//
//func (s *ServeRpc) Close() {
//	s.cc.Close()
//}

package handlers

import (
	"../db"
	pb "../protos"
	"context"
	"errors"
	"fmt"
	"log"
)
type UserLogin struct {
	user_id 		string
	user_password	string
}

// 用户登录
func (*UserServer) UserLogin(ctx context.Context, req *pb.LoginRequest) (*pb.Null, error){
	fmt.Println("user_id", req.UserId, "user_password:",req.UserPassword)
	ok, err:= db.SearchUserForLogin(req.UserId, req.UserPassword)
	if err != nil {
		log.Println(err.Error())
		return  &pb.Null{},err
	}
	if ok != true {
		log.Println("账号或密码不正确！")
		return &pb.Null{},errors.New("账号或密码不正确！")
	}
	return &pb.Null{},nil
}

// 用户注册
func (*UserServer) UserRegister(ctx context.Context, req *pb.RegisterRequest) (*pb.UserId, error){
	log.Println("user_id", req.UserName, "user_password:",req.UserPassword)
	user_id, err := db.AddUer(req.UserName, req.UserPassword)
	if err != nil {
		log.Println(err.Error())
		return &pb.UserId{},err
	}
	var ret pb.UserId
	ret.UserId = user_id
	return &ret,nil
}

// 用户信息查询
func (*UserServer) GetUserInfo(ctx context.Context, req *pb.UserId) (*pb.UserInfoReply, error) {
	log.Println("user_id", req.UserId)
	var ret pb.UserInfoReply
	user, err := db.SearchUser(req.UserId)
	if err != nil {
		log.Println(err.Error())
		return &ret,err
	}
	ret.UserId = user.User_id
	ret.UserName = user.User_name
	ret.UserPassword = user.Password
	ret.UserScore = int32(user.Score)
	return &ret,nil
}

// 用户信息修改
func (*UserServer)  UserUpdate(ctx context.Context, req *pb.UserUpdateRequest) (*pb.Null, error) {
	log.Println("user_id", req.UserId, "user_password:",req.UserPassword)
	err := db.UpdateUser(req.UserId, req.UserName, req.UserPassword)
	if err != nil {
		log.Println(err.Error())
		return &pb.Null{}, err
	}
	return &pb.Null{}, nil
}

// 用户垃圾数据统计
func (*UserServer) WasteCount(ctx context.Context, req *pb.UserId) (*pb.WasteCountReply, error) {
	log.Println("userId:",req.UserId)
	var ret  pb.WasteCountReply
	sum, types, err := db.GetWasteCount(req.UserId)
	if err != nil {
		log.Println(err.Error())
		return &pb.WasteCountReply{}, err
	}
	ret.Sum = int32(sum)
	for _, val := range types {
		ret.Type = append(ret.Type, int32(val))
	}
	return &ret, nil
}

// 用户最近一周垃圾数据统计
func (*UserServer) WeekWasteCount(ctx context.Context, req *pb.UserId) (*pb.WeekWasteCountReply, error) {
	log.Println("userId:",req.UserId)
	var ret  pb.WeekWasteCountReply
	sum, types, err := db.GetWasteCount(req.UserId)
	if err != nil {
		log.Println(err.Error())
		return &pb.WeekWasteCountReply{}, err
	}
	ret.Sum = int32(sum)
	for _, val := range types {
		ret.Type = append(ret.Type, int32(val))
	}
	return &ret, nil
}

// 实时获取用户垃圾桶状态
func (*UserServer) GetBinStatus(ctx context.Context, req *pb.UserId) (*pb.BinStatusReply, error) {
	log.Println("userId:",req.UserId)
	var ret pb.BinStatusReply
	binStatus, err := db.GetBinStatuses(req.UserId)
	if err != nil {
		return &ret, err
	}
	ret.BinStatus = binStatus
	return &ret, nil
}

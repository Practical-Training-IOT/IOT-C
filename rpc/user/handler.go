package main

import (
	"context"
	"errors"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/config"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/database"
	user "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/user"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// Register implements the UserImpl interface.
func (s *UserImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterRes, err error) {
	// TODO: Your code here...
	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	users := database.UserC{
		Username:         req.UserName,
		Password:         string(password),
		RegistrationDate: time.Now(),
	}
	err = config.DB.Create(&users).Error
	return &user.RegisterRes{
		Id: int32(users.ID),
	}, nil
}

// Login implements the UserImpl interface.
func (s *UserImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginRes, err error) {
	// TODO: Your code here...
	var users database.UserC
	err = config.DB.Where("username = ?", req.UserName).First(&users).Error
	if err != nil {
		return nil, errors.New("登录失败")
	}
	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("登录失败")
	}
	return &user.LoginRes{
		Id: int32(users.ID),
	}, nil
}

// UserInfo implements the UserImpl interface.
func (s *UserImpl) UserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoRes, err error) {
	// TODO: Your code here...
	var users database.UserC
	err = config.DB.Where("id = ?", req.Id).First(&users).Error
	if err != nil {
		return nil, errors.New("请刷新重试")
	}
	var phone string
	var email string
	if users.PhoneNumber == "" {
		phone = "暂未设置"
	} else {
		phone = users.PhoneNumber
	}
	if users.Email == "" {
		email = "暂未设置"
	} else {
		email = users.Email
	}
	resp = &user.UserInfoRes{
		Username:   users.Username,
		Department: users.Department,
		Email:      email,
		Phone:      phone,
		Avatar:     users.ProfilePictureURL,
	}
	return resp, nil
}

// UserInfoUpload implements the UserImpl interface.
func (s *UserImpl) UserInfoUpload(ctx context.Context, req *user.UserInfoUploadReq) (resp *user.UserInfoUploadRes, err error) {
	// TODO: Your code here...
	users := database.UserC{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: strconv.FormatInt(req.Phone, 10),
		Department:  req.Department,
	}
	err = config.DB.Model(&database.UserC{}).Where("id=?", req.Id).Updates(&users).Error
	if err != nil {
		return nil, errors.New("修改失败")
	}
	return &user.UserInfoUploadRes{}, nil
}

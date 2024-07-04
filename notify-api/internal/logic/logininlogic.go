package logic

import (
	"context"
	"errors"
	"notify/common"
	notify_model "notify/notify-model"

	"notify/notify-api/internal/svc"
	"notify/notify-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginInLogic {
	return &LoginInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginInLogic) LoginIn(req *types.LoginInRequest) (resp *types.LoginInResponse, err error) {

	// 查找是否存在这个用户
	user := notify_model.User{
		UserID: req.UserID,
	}
	err = l.svcCtx.DB.Find(&user).Error
	if err != nil {
		return nil, errors.New("没有此用户")
	}
	// 查看密码是否一致
	if user.Password != req.Password {
		return nil, errors.New("用户名或密码错误")
	}
	// 获取token
	token, err := common.GetToken(req.UserID)
	if err != nil {
		return nil, errors.New("登录失败")
	}

	return &types.LoginInResponse{
		token,
	}, nil
}

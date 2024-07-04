package logic

import (
	"context"
	"errors"
	"notify/common"
	"notify/notify-api/internal/svc"
	"notify/notify-api/internal/types"
	notify_model "notify/notify-model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetNoticeLogic {
	return &SetNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetNoticeLogic) SetNotice(req *types.SetNoticeRequest) (resp *types.SetNoticeResponse, err error) {
	// 验证
	userid, err := common.Auth(req.Token)
	if err != nil {
		return nil, err
	}

	// 消息的userid与jwt的userid不一致
	if userid != req.Notice.UserID {
		return nil, errors.New("只能操作自己的notice")
	}

	//拿到消息
	notice := notify_model.Notice{
		CreatorID:    req.Notice.UserID,
		Content:      req.Notice.Content,
		ReminderTime: time.Unix(req.Notice.NotifyTime, 0),
	}

	err = l.svcCtx.DB.Create(&notice).Error
	if err != nil {
		return nil, errors.New("插入数据失败")
	}
	return
}

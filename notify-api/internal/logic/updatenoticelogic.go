package logic

import (
	"context"
	"errors"
	"notify/common"
	notify_model "notify/notify-model"
	"time"

	"notify/notify-api/internal/svc"
	"notify/notify-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogic {
	return &UpdateNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNoticeLogic) UpdateNotice(req *types.UpdateNoticeRequest) (resp *types.UpdateNoticeResponse, err error) {
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
		ID:           req.Id,
		CreatorID:    req.Notice.UserID,
		Content:      req.Notice.Content,
		ReminderTime: time.Unix(req.Notice.NotifyTime, 0),
	}

	// 更新数据
	err = l.svcCtx.DB.Model(&notice).Updates(notice).Error
	if err != nil {
		return nil, errors.New("更新数据失败")
	}

	return nil, nil
}

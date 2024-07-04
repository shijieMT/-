package logic

import (
	"context"
	"errors"
	"notify/common"
	"notify/notify-api/internal/svc"
	"notify/notify-api/internal/types"
	notify_model "notify/notify-model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoticeLogic {
	return &DeleteNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNoticeLogic) DeleteNotice(req *types.DeleteNoticeRequest) (resp *types.DeleteNoticeResponse, err error) {
	// 验证
	userid, err := common.Auth(req.Token)
	if err != nil {
		return nil, err
	}

	// 删除
	err = l.svcCtx.DB.Where("creator_id = ?", userid).Where("id = ?", req.Id).Delete(&notify_model.Notice{}).Error
	if err != nil {
		return nil, errors.New("删除数据错误")
	}

	return nil, nil
}

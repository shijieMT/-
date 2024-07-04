package logic

import (
	"context"
	"errors"
	"fmt"
	"notify/common"
	"notify/notify-api/internal/svc"
	"notify/notify-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeListLogic {
	return &GetNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoticeListLogic) GetNoticeList(req *types.GetNoticeListRequest) (resp *types.GetNoticeListResponse, err error) {
	// 验证
	userid, err := common.Auth(req.Token)
	if err != nil {
		return nil, err
	}

	notices := make([]types.Notice, 0)

	err = l.svcCtx.DB.Where("creator_id = ?", userid).Find(&notices).Error
	if err != nil {
		return nil, errors.New("查询数据失败")
	}
	// todo
	for _, notice := range notices {
		fmt.Println(notice)
	}

	return &types.GetNoticeListResponse{
		NoticeList: notices,
	}, nil
}

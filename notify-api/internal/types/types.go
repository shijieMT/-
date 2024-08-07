// Code generated by goctl. DO NOT EDIT.
package types

type DeleteNoticeRequest struct {
	Token string `json:"token"`
	Id    int    `json:"id"`
}

type DeleteNoticeResponse struct {
}

type GetNoticeListRequest struct {
	Token string `json:"token"`
}

type GetNoticeListResponse struct {
	NoticeList []Notice `json:"noticeList"`
}

type LoginInRequest struct {
	UserID   int    `json:"userId"`
	Password string `json:"password"`
}

type LoginInResponse struct {
	Token string `json:"token"`
}

type Notice struct {
	UserID     int    `json:"userId"`
	Content    string `json:"content"`
	NotifyTime int64  `json:"notifyTime"`
}

type SetNoticeRequest struct {
	Token  string `json:"token"`
	Notice Notice `json:"notice"`
}

type SetNoticeResponse struct {
}

type UpdateNoticeRequest struct {
	Token  string `json:"token"`
	Id     int    `json:"id"`
	Notice Notice `json:"notice"`
}

type UpdateNoticeResponse struct {
}

type WSNoticeRequest struct {
	Token string `json:"token"`
}

type WSNoticeResponse struct {
}

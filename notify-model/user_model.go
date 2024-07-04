package notify_model

// User 用户表
type User struct {
	UserID   int    `gorm:"primaryKey" json:"user_id"` // 用户ID
	Password string `gorm:"size:64" json:"password"`   // 用户密码
}

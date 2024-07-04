package notify_model

import "time"

// Notice 通知表
type Notice struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"` // 主键ID，自动递增
	CreatorID    int       `json:"creator_id"`                         // 通知创建者ID
	Content      string    `gorm:"size:255" json:"content"`            // 提醒内容
	ReminderTime time.Time `json:"reminder_time"`                      // 提醒时间
}

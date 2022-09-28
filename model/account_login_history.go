package model

import "github.com/cnpythongo/goal/pkg/basic"

type LoginHistory struct {
	basic.BaseModel
	UserID int `json:"user_id" gorm:"index:loginhistory_user_id;column:user_id;type:int(11);not null"` // 用户ID
}

func (h *LoginHistory) TableName() string {
	return "account_login_history"
}

func NewLoginHistory() *LoginHistory {
	return &LoginHistory{}
}

func NewLoginHistoryList() []*LoginHistory {
	return make([]*LoginHistory, 0)
}

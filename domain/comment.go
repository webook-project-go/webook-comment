package domain

import "time"

type Comment struct {
	ID  int64 `json:"id"`
	UID int64 `json:"uid"`
	PID int64 `json:"pid"`
	RID int64 `json:"rid"`

	Biz   string `json:"biz"`
	BizID int64  `json:"biz_id"`

	Content string    `json:"content"`
	Ctime   time.Time `json:"ctime"`
}

package dao

import "database/sql"

type Comment struct {
	ID   int64 `gorm:"primaryKey;autoIncrement"`
	UID  int64
	RID  sql.NullInt64 `gorm:"index"`
	Root *Comment      `gorm:"ForeignKey:RID;AssociationForeignKey:ID;constraint:OnDelete:CASCADE"`
	PID  sql.NullInt64 `gorm:"index"`

	Biz   string `gorm:"type:varchar(30);index:biz_bizid,priority:1"`
	BizID int64  `gorm:"index:biz_bizid,priority:2"`

	Content string `gorm:"type:text"`
	Ctime   int64
}

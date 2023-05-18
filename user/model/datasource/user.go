package datasource

import (
	"gorm.io/gorm"
	"time"
)

// User 表
type User struct {
	gorm.Model
	//使用雪花id
	UserId int64 `gorm:"primary_key;not null" json:"userId"`
	//推特id
	TwitterId string `gorm:"type:varchar(256);not null;index:idx_twitter_id" json:"twitterId"`
	//关联 Authentication 表
	Authentications []Authentication `gorm:"foreignKey:UserId"` //外键连接authentication表
	//关联 Profile 表
	Profile Profile `gorm:"foreignKey:UserId"`
	//关联 Black 表
	Blacks []Black `gorm:"foreignKey:UserId"`
	//关联 Invitation 表
	Invitations []Invitation `gorm:"foreignKey:InviterId"` // 外键连接 Invitation 表，表示一个用户发送的所有邀请
}

// Invitation 表
type Invitation struct {
	gorm.Model
	//使用雪花id
	InvitationId int64 `gorm:"primary_key;not null" json:"invitationId"`
	InviterId    int64 `gorm:"not null;index:idx_inviter_id" json:"inviterId"` // 邀请人Id
	InviteeId    int64 `gorm:"not null;index:idx_invitee_id" json:"inviteeId"` // 被邀请人Id
}

type InvitationTree struct {
	gorm.Model
	InvitationTreeId int64 `gorm:"primary_key;not null" json:"invitationTreeId"`
	UserID           int64 `gorm:"not null;index:idx_user_id" json:"userId"`     // 用户 ID
	ParentID         int64 `gorm:"not null;index:idx_parent_id" json:"parentId"` // 父节点用户 ID
	Depth            int   `gorm:"not null;index:idx_depth" json:"depth"`        // 深度信息
}

// Profile 表
type Profile struct {
	gorm.Model
	//使用雪花id 方便迁移
	ProfileId int64 `gorm:"primary_key;not null" json:"ProfileId"`
	//用户id
	UserId int64 `gorm:"type:bigint;not null;index:idx_user_id" json:"userId"`
	//用户名
	UserName string `gorm:"type:varchar(256);not null;index:idx_user_name" json:"userName"`
	//用户头像
	Avatar string `gorm:"type:varchar(256);not null" json:"avatar"`
}

// Authentication 表
type Authentication struct {
	gorm.Model
	//使用雪花id
	AuthenticationId int64 `gorm:"primary_key;not null" json:"authenticationId"`
	//用户id
	UserId int64 `gorm:"type:bigint;not null;index:idx_user_id" json:"userId"`
	//用户认证类型
	AuthType string `gorm:"type:varchar(256);not null;index:idx_auth_type" json:"authType"` //密码认证、短信认证、支付宝认证、微信认证、twitter认证
	//用户认证状态
	AuthStatus string `gorm:"type:varchar(256);not null;index:idx_auth_status" json:"authStatus"` //已经认证，认证失败，待认证
	//长token
	LongToken string `gorm:"type:varchar(256);not null;index:idx_long_token" json:"longToken"`
	//短token
	ShortToken string `gorm:"type:varchar(256);not null;index:idx_short_token" json:"shortToken"`
}

// Black 表
type Black struct {
	gorm.Model
	//使用雪花id
	BlackId int64 `gorm:"primary_key;not null" json:"blackId"`
	//用户id
	UserId int64 `gorm:"type:bigint;not null;index:idx_user_id" json:"userId"`
	//用户
	User User `gorm:"foreignKey:UserId"`
	//黑名单用户id
	BlackUserId int64 `gorm:"type:bigint;not null;index:idx_black_user_id" json:"blackUserId"`
	//被黑名单的用户
	BlackUser User `gorm:"foreignKey:BlackUserId"`
	//开始时间
	StartTime time.Time `gorm:"type:datetime;not null;index:idx_start_time" json:"startTime"`
	//结束时间
	EndTime time.Time `gorm:"type:datetime;not null;index:idx_end_time" json:"endTime"`
	//黑名单状态 0:正常 1:拉黑
	BlackStatus string `gorm:"type:varchar(256);not null;index:idx_black_status" json:"blackStatus"`
}

// Notice 表
type Notice struct {
	gorm.Model
	//使用雪花id
	NoticeId int64 `gorm:"primary_key;not null" json:"noticeId"`
	//用户id
	UserId int64 `gorm:"type:bigint;not null;index:idx_user_id" json:"userId"`
	//用户
	User User `gorm:"foreignKey:UserId"`
	//通知标题
	NoticeTitle string `gorm:"type:varchar(256);index:idx_notice_title" json:"noticeTitle"`
	//通知内容
	NoticeContent string `gorm:"type:varchar(256);not null;index:idx_notice_content" json:"noticeContent"`
	//通知状态 0:正常 1:删除
	NoticeStatus string `gorm:"type:varchar(256);not null;index:idx_notice_status" json:"noticeStatus"`
	//通知类型 0:系统通知 1:用户通知
	NoticeType string `gorm:"type:varchar(256);not null;index:idx_notice_type" json:"noticeType"`
	//通知时间
	NoticeTime time.Time `gorm:"type:datetime;not null;index:idx_notice_time" json:"noticeTime"`
	//是否已读
	IsRead string `gorm:"type:varchar(256);not null;index:idx_is_read" json:"isRead"`
}

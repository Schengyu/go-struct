package models

import "time"

// User 用户账户表（严格对应 MySQL 表结构）
type User struct {
	ID                      int        `gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键;unsigned;not null" json:"id"`
	CreateTime              time.Time  `gorm:"column:create_time;type:datetime;comment:创建时间;default:NULL" json:"create_time"`
	UpdateTime              time.Time  `gorm:"column:update_time;type:datetime;comment:更新时间;default:NULL" json:"update_time"`
	IsDel                   *int       `gorm:"column:is_del;type:smallint;comment:是否删除（0:否,1:是）;default:NULL" json:"is_del"`
	Account                 string     `gorm:"column:account;type:varchar(64);charset=utf8mb3;collation=utf8mb3_general_ci;comment:账号;not null" json:"account"`
	Password                string     `gorm:"column:password;type:varchar(256);charset=utf8mb3;collation=utf8mb3_general_ci;comment:密码" json:"password"`
	TowPassword             string     `gorm:"column:tow_password;type:varchar(256);charset=utf8mb3;collation=utf8mb3_general_ci;comment:二次密码" json:"tow_password"`
	Permission              *int       `gorm:"column:permission;type:smallint;comment:权限等级;default:NULL" json:"permission"`
	Remark                  string     `gorm:"column:remark;type:varchar(100);charset=utf8mb3;collation=utf8mb3_general_ci;comment:备注" json:"remark"`
	RegisterID              *int       `gorm:"column:register_id;type:int;comment:注册来源ID;default:NULL" json:"register_id"`
	TeamUID                 *int       `gorm:"column:team_uid;type:int;comment:团队UID;default:NULL" json:"team_uid"`
	MasterUID               *int       `gorm:"column:master_uid;type:int;comment:主账户UID;default:NULL" json:"master_uid"`
	LeaseHost               string     `gorm:"column:lease_host;type:varchar(100);charset=utf8mb3;collation=utf8mb3_general_ci;comment:租赁主机" json:"lease_host"`
	LeaseDB                 string     `gorm:"column:lease_db;type:varchar(20);charset=utf8mb3;collation=utf8mb3_general_ci;comment:租赁数据库" json:"lease_db"`
	SteamWalletAppHost      string     `gorm:"column:steam_wallet_app_host;type:varchar(30);charset=utf8mb3;collation=utf8mb3_general_ci;comment:Steam钱包应用Host" json:"steam_wallet_app_host"`
	SteamWalletAppShareUser string     `gorm:"column:steam_wallet_app_share_user;type:varchar(30);charset=utf8mb3;collation=utf8mb3_general_ci;comment:Steam钱包共享用户" json:"steam_wallet_app_share_user"`
	SteamWalletAppSymbol    string     `gorm:"column:steam_wallet_app_symbol;type:varchar(30);charset=utf8mb3;collation=utf8mb3_general_ci;comment:Steam钱包货币符号" json:"steam_wallet_app_symbol"`
	Lock                    *int       `gorm:"column:lock;type:smallint;comment:锁定状态（0:正常,1:锁定）;default:NULL" json:"lock"`
	DelAccount              string     `gorm:"column:del_account;type:varchar(64);charset=utf8mb3;collation=utf8mb3_general_ci;comment:删除账号" json:"del_account"`
	Phone                   string     `gorm:"column:phone;type:varchar(16);charset=utf8mb3;collation=utf8mb3_general_ci;comment:手机号" json:"phone"`
	Integration             int        `gorm:"column:integration;type:int;comment:积分;default:0;not null" json:"integration"`
	NickName                string     `gorm:"column:nick_name;type:varchar(128);charset=utf8mb3;collation=utf8mb3_general_ci;comment:昵称" json:"nick_name"`
	Avatar                  string     `gorm:"column:avatar;type:varchar(256);charset=utf8mb3;collation=utf8mb3_general_ci;comment:头像" json:"avatar"`
	AvatarSign              string     `gorm:"column:avatar_sign;type:varchar(64);charset=utf8mb3;collation=utf8mb3_general_ci;comment:头像签名" json:"avatar_sign"`
	Sex                     int        `gorm:"column:sex;type:smallint;comment:性别（0:未知,1:男,2:女）;default:0;not null" json:"sex"`
	Summary                 string     `gorm:"column:summary;type:varchar(200);charset=utf8mb3;collation=utf8mb3_general_ci;comment:个人简介" json:"summary"`
	Birthday                *time.Time `gorm:"column:birthday;type:datetime;comment:生日;default:NULL" json:"birthday"`
	YymBalance              float64    `gorm:"column:yym_balance;type:decimal(12,5);comment:YYM余额;default:NULL" json:"yym_balance"`
	YymFenBalance           int64      `gorm:"column:yym_fen_balance;type:bigint;comment:YYM分余额" json:"yym_fen_balance"`
	RealName                string     `gorm:"column:real_name;type:varchar(20);charset=utf8mb3;collation=utf8mb3_general_ci;comment:真实姓名" json:"real_name"`
	RealID                  string     `gorm:"column:real_id;type:varchar(20);charset=utf8mb3;collation=utf8mb3_general_ci;comment:身份证号" json:"real_id"`
	RealStatus              int        `gorm:"column:real_status;type:smallint;comment:实名状态（0:未实名,1:审核中,2:已通过,3:已拒绝）;default:0;not null" json:"real_status"`
	RealTime                *time.Time `gorm:"column:real_time;type:datetime;comment:实名时间;default:NULL" json:"real_time"`
}

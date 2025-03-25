package models

import "gorm.io/gorm"

// Inventory 库存信息
type YYMCSGOBoxInventory struct {
	gorm.Model
	SkuCodeI       string       `gorm:"column:sku_code;type:varchar(255);comment:属性组合（如 1:10_3:6_5:21）" json:"sku_code"`
	SteamAccountID int          `gorm:"column:steam_account_id;type:int;comment:steam账号ID" json:"steam_account_id"`
	YYMAccount     string       `gorm:"column:yym_account;type:varchar(50);comment:元游猫账号，存储账号名称" json:"yym_account"`
	BoxNumber      int          `gorm:"column:box_number;type:int;comment:箱子数量" json:"box_number"`
	KeyNumber      int          `gorm:"column:key_number;type:int;comment:钥匙数量" json:"key_number"`
	ValidFrom      int64        `gorm:"column:valid_from;type:bigint;comment:从发货开始计算的有效期，记录该库存商品可使用或有效的起始时间点(存储时间戳)" json:"valid_from"`
	IsUse          int          `gorm:"column:is_use;type:int;comment:是否已使用" json:"is_use"`
	CreateP        string       `gorm:"column:create_p;type:varchar(50);comment:创建人 管理员名称" json:"create_p"`
	YYMGoodsSKU    YYMGoodsSKU  `gorm:"foreignKey:SkuCodeI;references:SkuCode"`
	SteamAccount   SteamAccount `gorm:"foreignKey:SteamAccountID"`
}

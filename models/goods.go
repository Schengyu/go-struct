package models

import "gorm.io/gorm"

// Good 商品信息
type YYMGoodMaster struct {
	gorm.Model
	CategoryID  int         `gorm:"column:category_id;type:int(11);comment:关联类别表的主键，标识商品所属的类目" json:"category_id"`
	ProductName string      `gorm:"column:product_name;type:varchar(100);comment:商品名称，例如 “超级 VIP 服务”、“CSGO 箱子”、“游戏礼包” 等" json:"product_name"`
	Description string      `gorm:"column:description;type:text;comment:商品描述，详细介绍商品的特点、内容" json:"description"`
	Price       float64     `gorm:"column:price;type:decimal(10,2);comment:商品基础价格，可作为默认价格（SKU 可能存在差异）" json:"price"`
	IsActive    int8        `gorm:"column:is_active;type:tinyint(1);comment:是否上架，1 表示上架状态，0 表示下架" json:"is_active"`
	CreateP     string      `gorm:"column:create_p;type:varchar(50);comment:创建人 / 管理员名称或账号" json:"create_p"`
	YYMCategory YYMCategory `gorm:"foreignKey:CategoryID"`
}

// Category 商品类目信息
type YYMCategory struct {
	gorm.Model
	CategoryName     string `gorm:"column:category_name;type:varchar(50);not null;comment:商品类目的名称，长度不超过 50 字符" json:"category_name"`
	ParentCategoryID int    `gorm:"column:parent_category_id;type:int;default:null;comment:父类目的 ID，若为顶级类目则为 NULL" json:"parent_category_id"`
	Description      string `gorm:"column:description;type:text;comment:商品类目的详细描述，可说明该类目下商品特点" json:"description"`
	CreateP          string `gorm:"column:create_p;type:varchar(50);comment:创建人 管理员名称" json:"create_p"`
}

// Attribute 属性表信息
type YYMAttribute struct {
	gorm.Model
	CategoryID    int         `gorm:"column:category_id;type:int(11);comment:关联类目表的主键 Id，建立与类目表的多对一关系，表明该属性属于哪个类目" json:"category_id"`
	AttributeType int8        `gorm:"column:attribute_type;type:tinyint(2);comment:属性字段类型  0对应属性  1对应属性值" json:"attribute_type"`
	Attribute     string      `gorm:"column:attribute;type:varchar(50);comment:属性/属性值  当attribute_type为0时该字段为属性 为1是该字段为属性值" json:"attribute"`
	ParentID      int64       `gorm:"column:parent_id;type:bigint;comment:关联父级属性，对应父级属性相应的属性值" json:"parent_id"`
	SortOrder     int         `gorm:"column:sort_order;type:int(2);comment:属性的排序顺序，用于控制在前端展示或其他场景下属性的显示顺序，值越小越靠前" json:"sort_order"`
	CreateP       string      `gorm:"column:create_p;type:varchar(50);comment:创建人存储管理员名称/账号" json:"create_p"`
	AttributeCode string      `gorm:"column:attribute_code;type:varchar(50);comment:属性值具体对应的数值" json:"attribute_code"`
	YYMCategory   YYMCategory `gorm:"foreignKey:CategoryID"`
}

// SKU 商品的 SKU 信息
type YYMGoodsSKU struct {
	gorm.Model
	GoodsId       int           `gorm:"column:goods_id;type:int(11);comment:关联 good_id，标识该 SKU 属于哪个商品" json:"goods_id"`
	SkuUUID       string        `gorm:"column:sku_uuid;type:varchar(50);comment:SKU 编码，可自动生成" json:"sku_uuid"`
	SkuCode       string        `gorm:"column:sku_code;type:varchar(255);unique;comment:属性组合（如 1:10_3:6_5:21）" json:"sku_code"`
	IsActive      int8          `gorm:"column:is_active;type:tinyint(1);comment:是否上架，1=上架，0=下架" json:"is_active"`
	CreateP       string        `gorm:"column:create_p;type:varchar(50);comment:创建人存储管理员名称/账号" json:"create_p"`
	YYMGoodMaster YYMGoodMaster `gorm:"foreignKey:GoodsId"`
}

// SellerGoodsSKU 表示卖家商品 SKU 表的结构体   也就是说卖家自己发布的所有商品都会存储在这张表当中
type YYMSellerGoodsSKU struct {
	gorm.Model
	UserID         int            `gorm:"column:user_id;type:int;comment:用户ID" json:"user_id"`
	SellerGoodsID  int            `gorm:"column:seller_goods_id;type:int;comment:卖家商品ID" json:"seller_goods_id"`
	SkuCodeS       string         `gorm:"column:sku_code;type:varchar(255);comment:属性组合（如 1:10_3:6_5:21）" json:"sku_code"`
	Price          float64        `gorm:"column:price;type:decimal(10,2);comment:价格" json:"price"`
	MasterGoodsID  int            `gorm:"column:master_goods_id;type:int;comment:商品主表ID" json:"master_goods_id"`
	YYMGoodsSKU    YYMGoodsSKU    `gorm:"foreignKey:SkuCodeS;references:SkuCode"`
	YYMGoodMaster  YYMGoodMaster  `gorm:"foreignKey:MasterGoodsID"`
	User           User           `gorm:"foreignKey:UserID"`
	YYMSellerGoods YYMSellerGoods `gorm:"foreignKey:SellerGoodsID"`
}

// 卖家商品表
type YYMSellerGoods struct {
	gorm.Model
	UserID        int           `gorm:"column:user_id;type:int;comment:用户ID" json:"user_id"`
	GoodName      string        `gorm:"column:goods_name;type:varchar(100);comment:卖家发布商品的名称" json:"goods_name"`
	MasterGoodsID int           `gorm:"column:master_goods_id;type:int;comment:商品主表ID" json:"master_goods_id"`
	IsListing     int           `gorm:"column:is_listing;type:int;comment:是否上架" json:"is_listing"`
	GoodMaster    YYMGoodMaster `gorm:"foreignKey:MasterGoodsID"`
	User          User          `gorm:"foreignKey:UserID"`
}

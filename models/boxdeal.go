package models

import "gorm.io/gorm"

// CSGOBox  CSGO 箱子商品信息
type YYMCSGOBox struct {
	gorm.Model
	ProductName   string      `gorm:"column:product_name;type:varchar(50);comment:CSGO 箱子的名称，例如 “光谱 2 号武器箱”“狂牙大行动武器箱” 等" json:"product_name"`
	ProductType   int         `gorm:"column:product_type;type:int(3);comment:箱子的类型，比如普通箱子、纪念包、特殊活动箱子等" json:"product_type"`
	CateID        int         `gorm:"column:cate_id;type:int(3);comment:对应类别表当中的主键id" json:"cate_id"`
	Description   string      `gorm:"column:description;type:text;comment:对箱子的详细描述，包括可能包含的武器皮肤系列、稀有度特点等" json:"description"`
	ReleaseDate   int64       `gorm:"column:release_date;type:bigint;comment:箱子的发布日期，记录该箱子在游戏中正式推出的时间" json:"release_date"`
	Price         float64     `gorm:"column:price;type:decimal(10,2);comment:箱子的售卖价格，精确到小数点后两位" json:"price"`
	IsActive      int8        `gorm:"column:is_active;type:tinyint(1);comment:商品的上架状态，1 表示上架可售，0 表示下架" json:"is_active"`
	ImageURL      string      `gorm:"column:image_url;type:varchar(255);comment:箱子图片的 URL 地址，用于在展示页面显示箱子的外观" json:"image_url"`
	RarityLevel   int         `gorm:"column:rarity_level;type:int(2);comment:箱子的稀有度等级，如普通、罕见、非凡等" json:"rarity_level"`
	StockQuantity int         `gorm:"column:stock_quantity;type:int(11);comment:该箱子的库存数量，可用于整体库存管理，与卖家账号中实际数量区分开，比如可设置为总库存数量等" json:"stock_quantity"`
	CreateP       string      `gorm:"column:create_p;type:varchar(50);comment:创建人 管理员名称" json:"create_p"`
	YYMCategory   YYMCategory `gorm:"foreignKey:CateID"`
}

// SKU 结构体定义，用于表示商品库存保有单位（SKU）的相关信息   无用表
type YYMCSGOBoxSKU struct {
	gorm.Model
	GoodID      int         `gorm:"column:good_id;type:int(11);comment:商品的主键id，建立关联关系" json:"good_id"`
	ProductName string      `gorm:"column:product_name;type:varchar(50);comment:商品的名称，如各种CSGO箱子名称" json:"product_name"`
	ProductType int         `gorm:"column:product_type;type:int(2);comment:商品的类型，以整数标识不同CSGO箱子类型" json:"product_type"`
	PackageType int         `gorm:"column:package_type;type:int(3);comment:商品的套餐类型，用整数表示套装数量" json:"package_type"`
	Description string      `gorm:"column:description;type:text;comment:商品的详细描述，如箱子内包含的武器皮肤系列" json:"description"`
	ReleaseDate int64       `gorm:"column:release_date;type:bigint;comment:箱子的发布日期，记录该箱子在游戏中正式推出的时间" json:"release_date"`
	IsActive    int8        `gorm:"column:is_active;type:tinyint(1);comment:商品的上架状态，1表示上架，0表示下架" json:"is_active"`
	CreateP     string      `gorm:"column:create_p;type:varchar(50);comment:创建人 管理员名称" json:"create_p"`
	YYMCSGOBox  YYMCSGOBox  `gorm:"foreignKey:GoodID"`
	YYMCategory YYMCategory `gorm:"foreignKey:ProductType"`
}

// UnboxingRecord 开箱记录信息
type YYMUnboxingRecord struct {
	gorm.Model
	BuyerID              int                  `gorm:"column:buyer_id;type:int;comment:买家的唯一标识ID，关联买家用户表，方便查询买家相关信息" json:"buyer_id"`
	SellerID             int                  `gorm:"column:seller_id;type:int;comment:卖家的唯一标识ID，关联卖家用户表，便于获取卖家资料" json:"seller_id"`
	InventoryID          int                  `gorm:"column:inventory_id;type:int;comment:关联库存表的ID，表明买家开启武器箱所使用的账号在库存中的记录" json:"inventory_id"`
	OpeningTime          int64                `gorm:"column:opening_time;type:bigint;comment:开箱操作发生的时间" json:"opening_time"`
	ItemID               int                  `gorm:"column:item_id;type:int;comment:开出物品的唯一标识ID，用于在系统中唯一确定一个物品，方便后续对物品进行单独查询、管理和统计等操作" json:"item_id"`
	ItemName             string               `gorm:"column:item_name;type:varchar(20);comment:开出物品的名称" json:"item_name"`
	ItemType             int                  `gorm:"column:item_type;type:int(2);comment:物品类型，如武器皮肤、刀具、手套等，方便分类统计" json:"item_type"`
	ItemQuality          int                  `gorm:"column:item_quality;type:int(2);comment:物品品质，常见、罕见、稀有、神话、传说等，用于衡量物品稀有度" json:"item_quality"`
	ItemValue            float64              `gorm:"column:item_value;type:decimal(10, 2);comment:物品的价值，根据市场价格或其他评估标准确定" json:"item_value"`
	CaseName             string               `gorm:"column:case_name;type:varchar(20);comment:开启的武器箱名称，记录买家打开的是哪种箱子" json:"case_name"`
	ServerRegion         string               `gorm:"column:server_region;type:varchar(50);comment:开箱操作所在的服务器区域，若涉及不同地区服务器" json:"server_region"`
	TransactionID        int                  `gorm:"column:transaction_id;type:int;comment:关联的订单交易ID，可追溯此次开箱行为对应的购买交易" json:"transaction_id"`
	YYMCSGOBoxInventory  YYMCSGOBoxInventory  `gorm:"foreignKey:InventoryID"`
	YYMTransactionRecord YYMTransactionRecord `gorm:"foreignKey:TransactionID"`
	UserBuyer            User                 `gorm:"foreignKey:BuyerID"`
	UserSeller           User                 `gorm:"foreignKey:SellerID"`
}

// ChildOrderDetail 子订单明细信息
type YYMChildOrderDetail struct {
	gorm.Model
	POrderSn            string              `gorm:"column:p_order_sn;type:varchar(32);comment:关联父订单表的订单编号，明确该子订单属于哪个父订单" json:"p_order_sn"`
	InventoryID         int                 `gorm:"column:inventory_id;type:int;comment:关联库存表的ID，用于标识子订单中的账号具体对应库存中的哪一条记录，建立与库存的联系" json:"inventory_id"`
	AccountType         int                 `gorm:"column:account_type;type:int;comment:账号类型，取值可以为Steam账号、元游猫账号或其他类型，方便区分不同平台的账号" json:"account_type"`
	AccountName         string              `gorm:"column:account_name;type:varchar(50);comment:账号名称，记录具体的账号标识" json:"account_name"`
	AccountStatus       int                 `gorm:"column:account_status;type:int;comment:账号状态，取值为活跃、不活跃、冻结，用于了解账号当前的可用情况" json:"account_status"`
	TransferTime        int64               `gorm:"column:transfer_time;type:bigint;comment:账号转移时间，记录该账号完成交易转移的时间点" json:"transfer_time"`
	Remark              string              `gorm:"column:remark;type:text;comment:备注信息，可记录关于该账号交易的特殊说明、问题或其他相关事项" json:"remark"`
	YYMCSGOBoxInventory YYMCSGOBoxInventory `gorm:"foreignKey:InventoryID"`
	YYMBoxOrder         YYMBoxOrder         `gorm:"foreignKey:POrderSn;references:OrderSn"`
}

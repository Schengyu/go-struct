package models

import "gorm.io/gorm"

// Order 订单信息
type YYMBoxOrder struct {
	gorm.Model
	OrderSn        string        `gorm:"column:order_sn;type:varchar(32);not null;unique;comment:订单号，系统自动生成的唯一编号，用于快速识别和查询订单" json:"order_sn"`
	BuyerID        int           `gorm:"column:buyer_id;type:int;comment:买家的唯一标识ID，关联买家用户表，方便查询买家相关信息" json:"buyer_id"`
	SellerID       int           `gorm:"column:seller_id;type:int;comment:卖家的唯一标识ID，关联卖家用户表，便于获取卖家资料" json:"seller_id"`
	BoxSKU         int           `gorm:"column:box_sku;type:int;comment:关联 CSGOBoxSKU 表的 SKU，用于建立与具体商品 SKU 的联系" json:"box_sku"`
	OrderAmount    float64       `gorm:"column:order_amount;type:decimal(10, 2);comment:订单总金额，包含商品价格、手续费等所有费用" json:"order_amount"`
	OrderTime      int64         `gorm:"column:order_time;type:bigint;comment:下单时间，记录用户提交订单的具体时刻" json:"order_time"`
	PaymentStatus  int           `gorm:"column:payment_status;type:int;comment:支付状态，取值为未支付、已支付、已退款，用于跟踪订单的支付情况" json:"payment_status"`
	PaymentTime    int64         `gorm:"column:payment_time;type:bigint;comment:支付时间，记录订单完成支付的时间点" json:"payment_time"`
	DeliveryStatus int           `gorm:"column:delivery_status;type:int;comment:发货状态，取值为未发货、已发货、已交付，方便买家和卖家了解订单物流进度" json:"delivery_status"`
	YYMCSGOBoxSKU  YYMCSGOBoxSKU `gorm:"foreignKey:BoxSKU"` //后面再说
	UserBuyer      User          `gorm:"foreignKey:BuyerID"`
	UserSeller     User          `gorm:"foreignKey:SellerID"`
}

// TransactionRecord 交易记录信息
type YYMTransactionRecord struct {
	gorm.Model
	TransactionSn       string      `gorm:"column:transaction_sn;type:varchar(50);not null;unique;comment:交易编号，系统自动生成的唯一编号，方便快速识别和查询交易" json:"transaction_sn"`
	OrderSns            string      `gorm:"column:order_sn;type:varchar(32);comment:关联的订单编号，与订单表中的订单编号对应，明确该交易所属订单" json:"order_sn"`
	BuyerID             int         `gorm:"column:buyer_id;type:int;comment:买家的唯一标识ID，关联买家用户表，用于查询买家相关信息" json:"buyer_id"`
	SellerID            int         `gorm:"column:seller_id;type:int;comment:卖家的唯一标识ID，关联卖家用户表，用于获取卖家资料" json:"seller_id"`
	TransactionAmount   float64     `gorm:"column:transaction_amount;type:decimal(10, 2);comment:交易金额，包含商品价格、手续费等所有费用" json:"transaction_amount"`
	TransactionTime     int64       `gorm:"column:transaction_time;type:bigint;comment:交易发生的时间，记录用户完成支付或其他交易操作的时刻" json:"transaction_time"`
	PaymentMethod       int         `gorm:"column:payment_method;type:int(2);comment:支付方式，取值包括信用卡、借记卡、PayPal、银行转账、其他等，用于记录交易采用的支付途径" json:"payment_method"`
	TransactionStatus   int         `gorm:"column:transaction_status;type:int(2);comment:交易状态，取值为待处理、已完成、失败、已取消，用于跟踪交易的当前状态" json:"transaction_status"`
	PaymentGateway      string      `gorm:"column:payment_gateway;type:varchar(100);comment:支付网关名称，如支付宝、微信支付等第三方支付平台名称，若使用第三方支付" json:"payment_gateway"`
	TransactionCurrency int         `gorm:"column:transaction_currency;type:int;comment:交易货币类型，如CNY（人民币）、USD（美元）等" json:"transaction_currency"`
	RefundAmount        float64     `gorm:"column:refund_amount;type:decimal(10, 2);comment:退款金额，若存在退款情况，记录退款的具体金额" json:"refund_amount"`
	YYMBoxOrder         YYMBoxOrder `gorm:"foreignKey:OrderSns;references:OrderSn"`
	UserBuyer           User        `gorm:"foreignKey:BuyerID"`
	UserSeller          User        `gorm:"foreignKey:SellerID"`
}

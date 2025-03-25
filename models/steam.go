package models

import "time"

// SteamAccount Steam 账号信息表
type SteamAccount struct {
	ID                  uint       `gorm:"column:id;type:int;primaryKey;autoIncrement;unsigned;not null" json:"id"`
	CreateTime          *time.Time `gorm:"column:create_time;type:datetime;default:NULL" json:"create_time"`
	UpdateTime          *time.Time `gorm:"column:update_time;type:datetime;default:NULL" json:"update_time"`
	IsDel               *int       `gorm:"column:is_del;type:smallint;default:NULL" json:"is_del"`
	UserID              *int       `gorm:"column:user_id;type:int;index:ix_steam_account_user_id" json:"user_id"`
	GroupID             int        `gorm:"column:group_id;type:int;index" json:"group_id"`
	SteamType           string     `gorm:"column:steam_type;type:varchar(32);charset=utf8mb3;collation=utf8mb3_general_ci" json:"steam_type"`
	SteamAccount        string     `gorm:"column:steam_account;type:varchar(32);charset=utf8mb3;collation=utf8mb3_general_ci" json:"steam_account"`
	SteamPassword       string     `gorm:"column:steam_password;type:varchar(300);charset=utf8mb3;collation=utf8mb3_general_ci" json:"steam_password"`
	SteamWallet         string     `gorm:"column:steam_wallet;type:varchar(64);charset=utf8mb3;collation=utf8mb3_general_ci" json:"steam_wallet"`
	CostRate            *float64   `gorm:"column:cost_rate;type:float" json:"cost_rate"`
	BalanceCNY          *float64   `gorm:"column:balance_cny;type:float" json:"balance_cny"`
	Country             string     `gorm:"column:country;type:varchar(100);charset=utf8mb3;collation=utf8mb3_general_ci" json:"country"`
	Currency            *int       `gorm:"column:currency;type:int" json:"currency"`
	AddressCity         string     `gorm:"column:address_city;type:varchar(64);charset=utf8mb3;collation=utf8mb3_general_ci" json:"address_city"`
	AddressState        string     `gorm:"column:address_state;type:varchar(64);charset=utf8mb3;collation=utf8mb3_general_ci" json:"address_state"`
	SteamApikey         string     `gorm:"column:steam_apikey;type:varchar(128);charset=utf8mb3;collation=utf8mb3_general_ci" json:"steam_apikey"`
	ProxyURL            string     `gorm:"column:proxy_url;type:varchar(128);charset=utf8mb3;collation=utf8mb3_general_ci" json:"proxy_url"`
	Steamid             string     `gorm:"column:steamid;type:varchar(32);charset=utf8mb3;collation=utf8mb3_general_ci" json:"steamid"`
	InventoryURL        string     `gorm:"column:inventory_url;type:varchar(200);charset=utf8mb3;collation=utf8mb3_general_ci" json:"inventory_url"`
	TradeURL            string     `gorm:"column:trade_url;type:varchar(200);charset=utf8mb3;collation=utf8mb3_general_ci" json:"trade_url"`
	BanMarket           *int       `gorm:"column:ban_market;type:smallint" json:"ban_market"`
	Guard               string     `gorm:"column:guard;type:varchar(2000);charset=utf8mb3;collation=utf8mb3_general_ci" json:"guard"`
	Cookie              string     `gorm:"column:cookie;type:text;charset=utf8mb3;collation=utf8mb3_general_ci" json:"cookie"`
	OpenDeliver         *int       `gorm:"column:open_deliver;type:smallint" json:"open_deliver"`
	AutoDeliverTime     *time.Time `gorm:"column:auto_deliver_time;type:datetime;default:NULL" json:"auto_deliver_time"`
	SyncInventoryTime   *time.Time `gorm:"column:sync_inventory_time;type:datetime;default:NULL" json:"sync_inventory_time"`
	MarketHistoryCursor string     `gorm:"column:market_history_cursor;type:varchar(64);charset=utf8mb3;collation=utf8mb3_general_ci" json:"market_history_cursor"`
	MarketBuyTotalSpend *float64   `gorm:"column:market_buy_total_spend;type:float" json:"market_buy_total_spend"`
	Exc                 string     `gorm:"column:exc;type:text;charset=utf8mb3;collation=utf8mb3_general_ci" json:"exc"`
	RefreshToken        string     `gorm:"column:refresh_token;type:varchar(2000);charset=utf8mb3;collation=utf8mb3_general_ci" json:"refresh_token"`
	AccessToken         string     `gorm:"column:access_token;type:varchar(2000);charset=utf8mb3;collation=utf8mb3_general_ci" json:"access_token"`
	Remark              string     `gorm:"column:remark;type:varchar(200);charset=utf8mb3;collation=utf8mb3_general_ci" json:"remark"`
	Mail                string     `gorm:"column:mail;type:varchar(64);charset=utf8mb3;collation=utf8mb3_general_ci" json:"mail"`
	MailPassword        string     `gorm:"column:mail_password;type:varchar(32);charset=utf8mb3;collation=utf8mb3_general_ci" json:"mail_password"`
	Phone               string     `gorm:"column:phone;type:varchar(20);charset=utf8mb3;collation=utf8mb3_general_ci" json:"phone"`
	BoundPhone          string     `gorm:"column:bound_phone;type:varchar(20);charset=utf8mb3;collation=utf8mb3_general_ci" json:"bound_phone"`
	BoundMail           string     `gorm:"column:bound_mail;type:varchar(30);charset=utf8mb3;collation=utf8mb3_general_ci" json:"bound_mail"`
	Restrictions        string     `gorm:"column:restrictions;type:varchar(1000);charset=utf8mb3;collation=utf8mb3_general_ci" json:"restrictions"`
	RestrictionExpire   string     `gorm:"column:restriction_expire;type:varchar(30);charset=utf8mb3;collation=utf8mb3_general_ci" json:"restriction_expire"`
	Online              int        `gorm:"column:online;type:smallint;default:0" json:"online"`
	SteamPoints         *int       `gorm:"column:steam_points;type:int" json:"steam_points"`
	Avatar              string     `gorm:"column:avatar;type:varchar(300);charset=utf8mb3;collation=utf8mb3_general_ci" json:"avatar"`
	MemberSince         string     `gorm:"column:member_since;type:varchar(300);charset=utf8mb3;collation=utf8mb3_general_ci" json:"member_since"`
	Privacy             string     `gorm:"column:privacy;type:varchar(20);charset=utf8mb3;collation=utf8mb3_general_ci" json:"privacy"`
	CostMoney           *float64   `gorm:"column:cost_money;type:float" json:"cost_money"`
	CostDiscount        *float64   `gorm:"column:cost_discount;type:float" json:"cost_discount"`
	IsTop               int        `gorm:"column:is_top;type:int;index:idx_steam_account_is_top_user_id" json:"is_top"`
	TopTime             *time.Time `gorm:"column:top_time;type:datetime;default:NULL" json:"top_time"`
	UpdateValueTime     *time.Time `gorm:"column:update_value_time;type:datetime;default:NULL" json:"update_value_time"`
	CountryCode         string     `gorm:"column:country_code;type:varchar(10);charset=utf8mb3;collation=utf8mb3_general_ci" json:"country_code"`
	LoginTime           *time.Time `gorm:"column:login_time;type:datetime;default:NULL" json:"login_time"`
	LoginFailMsg        string     `gorm:"column:login_fail_msg;type:varchar(512);charset=utf8mb3;collation=utf8mb3_general_ci" json:"login_fail_msg"`
	CommunityBanned     *int       `gorm:"column:community_banned;type:smallint" json:"community_banned"`
	GiftValue           *float64   `gorm:"column:gift_value;type:float" json:"gift_value"`
	GameSpendValue      *float64   `gorm:"column:game_spend_value;type:float" json:"game_spend_value"`
	GameGiftGaveValue   *float64   `gorm:"column:game_gift_gave_value;type:float" json:"game_gift_gave_value"`
	FriendCount         *int       `gorm:"column:friend_count;type:int" json:"friend_count"`
	CountrySet          *int       `gorm:"column:country_set;type:smallint" json:"country_set"`
	ClientRefreshToken  string     `gorm:"column:client_refresh_token;type:varchar(2000);charset=utf8mb3;collation=utf8mb3_general_ci" json:"client_refresh_token"`
	VacBanned           *int       `gorm:"column:vac_banned;type:smallint" json:"vac_banned"`
	NumberOfVacBanned   *int       `gorm:"column:number_of_vac_banned;type:smallint" json:"number_of_vac_banned"`
	MarketEnabled       *int       `gorm:"column:market_enabled;type:smallint" json:"market_enabled"`
	LoginEresult        *int       `gorm:"column:login_eresult;type:int" json:"login_eresult"`
}

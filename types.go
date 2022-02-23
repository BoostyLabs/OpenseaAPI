package OpenseaAPI

type Asset struct {
	ID                   int64         `json:"id"`
	NumSales             int64         `json:"num_sales"`
	BackgroundColor      string        `json:"background_color,omitempty"`
	ImageURL             string        `json:"image_url,omitempty"`
	ImagePreviewURL      string        `json:"image_preview_url,omitempty"`
	ImageThumbnailURL    string        `json:"image_thumbnail_url,omitempty"`
	ImageOriginalURL     string        `json:"image_original_url,omitempty"`
	AnimationURL         string        `json:"animation_url,omitempty"`
	AnimationOriginalURL string        `json:"animation_original_url,omitempty"`
	Name                 string        `json:"name,omitempty"`
	Description          string        `json:"description,omitempty"`
	ExternalLink         string        `json:"external_link,omitempty"`
	AssetContract        AssetContract `json:"asset_contract"`
	Permalink            string        `json:"permalink,omitempty"`
	Collection           Collection    `json:"collection"`
	Decimals             int64         `json:"decimals"`
	TokenMetadata        string        `json:"token_metadata,omitempty"`
	Owner                FromAccount   `json:"owner"`
	TokenID              string        `json:"token_id,omitempty"`
}

type AssetContract struct {
	Address                     string `json:"address,omitempty"`
	AssetContractType           string `json:"asset_contract_type,omitempty"`
	CreatedDate                 string `json:"created_date,omitempty"`
	Name                        string `json:"name,omitempty"`
	NFTVersion                  string `json:"nft_version",omitempty`
	OpenseaVersion              string `json:"opensea_version,omitempty"`
	Owner                       int64  `json:"owner"`
	SchemaName                  string `json:"schema_name,omitempty"`
	Symbol                      string `json:"symbol,omitempty"`
	TotalSupply                 string `json:"total_supply,omitempty"`
	Description                 string `json:"description,omitempty"`
	ExternalLink                string `json:"external_link,omitempty"`
	ImageURL                    string `json:"image_url,omitempty"`
	DefaultToFiat               bool   `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int64  `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int64  `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int64  `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int64  `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int64  `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int64  `json:"seller_fee_basis_points"`
	PayoutAddress               string `json:"payout_address,omitempty"`
}

type DisplayData struct {
	CardDisplayStyle string `json:"card_display_style,omitempty"`
}

type Collection struct {
	BannerImageURL              string      `json:"banner_image_url,omitempty"`
	ChatURL                     string      `json:"chat_url,omitempty"`
	CreatedDate                 string      `json:"created_date,omitempty"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	Description                 string      `json:"description,omitempty"`
	DevBuyerFeeBasisPoints      string      `json:"dev_buyer_fee_basis_points,omitempty"`
	DevSellerFeeBasisPoints     string      `json:"dev_seller_fee_basis_points,omitempty"`
	DiscordURL                  string      `json:"discord_url,omitempty"`
	DisplayData                 DisplayData `json:"display_data"`
	ExternalURL                 string      `json:"external_url,omitempty"`
	Featured                    bool        `json:"featured"`
	FeaturedImageURL            string      `json:"featured_image_url,omitempty"`
	Hidden                      bool        `json:"hidden"`
	SafelistRequestStatus       string      `json:"safelist_request_status,omitempty"`
	ImageURL                    string      `json:"image_url,omitempty"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
	LargeImageURL               string      `json:"large_image_url,omitempty"`
	MediumUsername              string      `json:"medium_username,omitempty"`
	Name                        string      `json:"name,omitempty"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points,omitempty"`
	OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points,omitempty"`
	PayoutAddress               string      `json:"payout_address,omitempty"`
	RequireEmail                bool        `json:"require_email"`
	ShortDescription            string      `json:"short_description,omitempty"`
	Slug                        string      `json:"slug,omitempty"`
	TelegramURL                 string      `json:"telegram_url,omitempty"`
	TwitterUsername             string      `json:"twitter_username,omitempty"`
	InstagramUsername           string      `json:"instagram_username,omitempty"`
	WikiURL                     string      `json:"wiki_url,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`
}

type FromAccount struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url,omitempty"`
	Address       string `json:"address,omitempty"`
	Config        string `json:"config,omitempty"`
}

type PaymentToken struct {
	ID       int64  `json:"id"`
	Symbol   string `json:"symbol,omitempty"`
	Address  string `json:"address,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	Name     string `json:"name,omitempty"`
	Decimals int64  `json:"decimals"`
	ETHPrice string `json:"eth_price,omitempty"`
	USDPrice string `json:"usd_price,omitempty"`
}

type AssetEvent struct {
	ApprovedAccount         string       `json:"approved_account,omitempty"`
	Asset                   Asset        `json:"asset"`
	AssetBundle             AssetBundle  `json:"asset_bundle,omitempty"`
	AuctionType             string       `json:"auction_type,omitempty"`
	BidAmount               string       `json:"bid_amount,omitempty"`
	CollectionSlug          string       `json:"collection_slug,omitempty"`
	ContractAddress         string       `json:"contract_address,omitempty"`
	CreatedDate             string       `json:"created_date,omitempty"`
	CustomEventName         string       `json:"custom_event_name,omitempty"`
	DevFeePaymentEvent      string       `json:"dev_fee_payment_event,omitempty"`
	DevSellerFeeBasisPoints int64        `json:"dev_seller_fee_basis_points"`
	Duration                string       `json:"duration,omitempty"`
	EndingPrice             string       `json:"ending_price,omitempty"`
	EventType               string       `json:"event_type,omitempty"`
	FromAccount             FromAccount  `json:"from_account"`
	ID                      int64        `json:"id"`
	IsPrivate               bool         `json:"is_private,omitempty"`
	OwnerAccount            string       `json:"owner_account,omitempty"`
	PaymentToken            PaymentToken `json:"payment_token"`
	Quantity                string       `json:"quantity,omitempty"`
	Seller                  User         `json:"seller"`
	StartingPrice           string       `json:"starting_price,omitempty"`
	ToAccount               string       `json:"to_account,omitempty"`
	TotalPrice              string       `json:"total_price,omitempty"`
	Transaction             Transaction  `json:"transaction,omitempty"`
	WinnerAccount           FromAccount  `json:"winner_account,omitempty"`
	ListingTime             string       `json:"listing_time,omitempty"`
}

type Transaction struct {
	BlockHash        string      `json:"block_hash"`
	BlockNumber      string      `json:"block_number"`
	FromAccount      FromAccount `json:"from_account"`
	Id               int64       `json:"id"`
	Timestamp        string      `json:"timestamp"`
	ToAccount        FromAccount `json:"to_account"`
	TransactionHash  string      `json:"transaction_hash"`
	TransactionIndex string      `json:"transaction_index"`
}

type Stats struct {
	OneDayVolume          float64 `json:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change"`
	OneDaySales           float64 `json:"one_day_sales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price"`
	SevenDayVolume        float64 `json:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change"`
	SevenDaySales         float64 `json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change"`
	ThirtyDaySales        float64 `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
	TotalVolume           float64 `json:"total_volume"`
	TotalSales            float64 `json:"total_sales"`
	TotalSupply           float64 `json:"total_supply"`
	Count                 float64 `json:"count"`
	NumOwners             int64   `json:"num_owners"`
	AveragePrice          float64 `json:"average_price"`
	NumReports            int64   `json:"num_reports"`
	MarketCap             float64 `json:"market_cap"`
	FloorPrice            float64 `json:"floor_price"`
}

type AssetBundle struct {
	Maker         FromAccount   `json:"maker"`
	Slug          string        `json:"slug"`
	Assets        []Asset       `json:"assets"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	ExternalLink  interface{}   `json:"external_link"`
	AssetContract AssetContract `json:"asset_contract"`
	Permalink     string        `json:"permalink"`
	SellOrders    interface{}   `json:"sell_orders"`
}

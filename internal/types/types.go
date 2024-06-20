// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// required : true
	// min : 0
	Page uint64 `json:"page" validate:"required,number,gt=0"`
	// Page size | 单页数据行数
	// required : true
	// max : 100000
	PageSize uint64 `json:"pageSize" validate:"required,number,lt=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic ID request (int32) | 基础ID参数请求 (int32)
// swagger:model IDInt32Req
type IDInt32Req struct {
	// ID
	// Required: true
	Id int32 `json:"id" validate:"number"`
}

// Basic IDs request (int32) | 基础ID数组参数请求 (int32)
// swagger:model IDsInt32Req
type IDsInt32Req struct {
	// IDs
	// Required: true
	Ids []int32 `json:"ids"`
}

// Basic ID request (int32) | 基础ID地址参数请求 (int32)
// swagger:model IDInt32PathReq
type IDInt32PathReq struct {
	// ID
	// Required: true
	Id int32 `path:"id"`
}

// Basic ID request (uint32) | 基础ID参数请求 (uint32)
// swagger:model IDUint32Req
type IDUint32Req struct {
	// ID
	// Required: true
	Id uint32 `json:"id" validate:"number"`
}

// Basic IDs request (uint32) | 基础ID数组参数请求 (uint32)
// swagger:model IDsUint32Req
type IDsUint32Req struct {
	// IDs
	// Required: true
	Ids []uint32 `json:"ids"`
}

// Basic ID request (uint32) | 基础ID地址参数请求 (uint32)
// swagger:model IDUint32PathReq
type IDUint32PathReq struct {
	// ID
	// Required: true
	Id uint32 `path:"id"`
}

// Basic ID request (int64) | 基础ID参数请求 (int64)
// swagger:model IDInt64Req
type IDInt64Req struct {
	// ID
	// Required: true
	Id int64 `json:"id" validate:"number"`
}

// Basic IDs request (int64) | 基础ID数组参数请求 (int64)
// swagger:model IDsInt64Req
type IDsInt64Req struct {
	// IDs
	// Required: true
	Ids []int64 `json:"ids"`
}

// Basic ID request (int64) | 基础ID地址参数请求 (int64)
// swagger:model IDInt64PathReq
type IDInt64PathReq struct {
	// ID
	// Required: true
	Id int64 `path:"id"`
}

// Basic UUID request in path | 基础UUID地址参数请求
// swagger:model UUIDPathReq
type UUIDPathReq struct {
	// ID
	// Required: true
	Id string `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// required : true
	// max length : 36
	// min length : 36
	Id string `json:"id" validate:"required,len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
}

// The base ID response data | 基础ID信息
// swagger:model BaseIDInfo
type BaseIDInfo struct {
	// ID
	Id *uint64 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The base ID response data (int64) | 基础ID信息 (int64)
// swagger:model BaseIDInt64Info
type BaseIDInt64Info struct {
	// ID
	Id *int64 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The base ID response data (int32) | 基础ID信息 (int32)
// swagger:model BaseIDInt32Info
type BaseIDInt32Info struct {
	// ID
	Id *int32 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The base ID response data (uint32) | 基础ID信息 (uint32)
// swagger:model BaseIDUint32Info
type BaseIDUint32Info struct {
	// ID
	Id *uint32 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础UUID信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id *string `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// swagger:model FileInfo
type FileInfo struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	FileType uint32 `json:"fileType"`
}

// GameRuleInfo
// swagger:model GameRuleInfo
type GameRuleInfo struct {
	// Title
	Title string `json:"title"`
	// Content
	Content string `json:"content"`
}

// RobotNumInfo
// swagger:model RobotNumInfo
type RobotNumInfo struct {
	// Min
	Min int64 `json:"min"`
	// Max
	Max int64 `json:"max"`
}

// RobotLampNumInfo
// swagger:model RobotLampNumInfo
type RobotLampNumInfo struct {
	// Min
	Min int64 `json:"min"`
	// Max
	Max int64 `json:"max"`
}

// SettingInfo
// swagger:model SettingInfo
type SettingInfo struct {
	// GameRuleInfo | 游戏规则
	GameRule GameRuleInfo `json:"gameRule"`
	// RobotNumInfo | 机器人数量
	RobotNum RobotNumInfo `json:"robotNum"`
	// RobotLampNumInfo | 机器人投入羊数量
	RobotLampNum RobotLampNumInfo `json:"robotLampNum"`
	// WithdrawCommission | 提现手续费
	WithdrawCommission float32 `json:"withdrawCommission"`
	// MinWithdrawNum | 最小提现数量
	MinWithdrawNum uint32 `json:"minWithdrawNum"`
	// IdleTime | 机器人空闲秒数
	IdleTime uint32 `json:"idleTime"`
	// GameCommission | 系统抽佣
	GameCommission float32 `json:"gameCommission"`
	// WithdrawThreshold | 提现免审额度
	WithdrawThreshold uint32 `json:"withdrawThreshold"`
}

// FindSettingReq
// swagger:model FindSettingReq
type FindSettingReq struct {
	Module string `json:"module"`
}

// FindSettingResp
// swagger:model FindSettingResp
type FindSettingResp struct {
	BaseMsgResp
	// Data
	Data SettingInfo `json:"data"`
}

// UpdateSettingReq
// swagger:model UpdateSettingReq
type UpdateSettingReq struct {
	SettingInfo
}

// FindNoticeReq
// swagger:model FindNoticeReq
type FindNoticeReq struct {
	Module string `json:"module"`
}

// FindNoticeResp
// swagger:model FindNoticeResp
type FindNoticeResp struct {
	BaseDataInfo
}

// UpdateNoticeReq
// swagger:model UpdateNoticeReq
type UpdateNoticeReq struct {
	Notice string `json:"notice"`
}

// PlayerInfo
// swagger:model PlayerInfo
type PlayerInfo struct {
	// Id
	Id uint64 `json:"id"`
	// CreatedAt
	CreatedAt int64 `json:"createdAt"`
	// UpdatedAt
	UpdatedAt int64 `json:"updatedAt"`
	// Status
	Status uint32 `json:"status"`
	// 代理等级
	Rank uint32 `json:"rank"`
	// 余币数量
	Amount float64 `json:"amount"`
	// 邀请好友数量
	InvitedNum uint32 `json:"invitedNum"`
	// 累计奖励
	TotalIncome float64 `json:"totalIncome"`
	// 近100局胜率
	Recent100WinPercent float32 `json:"recent100WinPercent"`
	// 邀请码
	InviteCode string `json:"inviteCode"`
	// 邀请人ID
	InviterId uint64 `json:"inviterId"`
	// 邀请人邀请码
	InvitedCode string `json:"invitedCode"`
	// 剩余羊数量
	Lamb float32 `json:"lamb"`
	// web3钱包地址
	DepositAddress string `json:"depositAddress"`
	// 平台奖励比例（%）
	SystemCommission float32 `json:"systemCommission"`
	// 邮箱
	Email string `json:"email"`
}

// CreatePlayerReq
// swagger:model CreatePlayerReq
type CreatePlayerReq struct {
	// Status
	Status uint32 `json:"status"`
	// Rank
	Rank uint32 `json:"rank"`
	// Amount
	Amount float64 `json:"amount"`
	// 剩余羊数量
	Lamb float32 `json:"lamb"`
	// InvitedNum
	InvitedNum uint32 `json:"invitedNum"`
	// InvitedCode
	InvitedCode string `json:"invitedCode"`
	// 平台奖励比例（%）
	SystemCommission float32 `json:"systemCommission"`
	// web3钱包地址
	DepositAddress string `json:"depositAddress"`
	// 邮箱
	Email string `json:"email"`
}

// UpdatePlayerReq
// swagger:model UpdatePlayerReq
type UpdatePlayerReq struct {
	// Id
	Id uint64 `json:"id"`
	// Status
	Status uint32 `json:"status,optional"`
	// Rank
	Rank uint32 `json:"rank,optional"`
	// Amount
	Amount float64 `json:"amount,optional"`
	// 剩余羊数量
	Lamb float32 `json:"lamb"`
	// InvitedNum
	InvitedNum uint32 `json:"invitedNum,optional"`
	// InvitedCode
	InvitedCode string `json:"invitedCode,optional"`
	// 平台奖励比例（%）
	SystemCommission float32 `json:"systemCommission,optional"`
	// web3钱包地址
	DepositAddress string `json:"depositAddress,optional"`
	// 邮箱
	Email string `json:"email,optional"`
}

// DeletePlayerReq
// swagger:model DeletePlayerReq
type DeletePlayerReq struct {
	// Id
	Id uint64 `json:"id"`
}

// FindPlayerReq
// swagger:model FindPlayerReq
type FindPlayerReq struct {
	// Id
	Id uint64 `json:"id"`
}

// swagger:model FindPlayerResp
type FindPlayerResp struct {
	BaseMsgResp
	// Data
	Data PlayerInfo `json:"data"`
}

// ListPlayerReq
// swagger:model ListPlayerReq
type ListPlayerReq struct {
	// Page
	Page uint64 `json:"page"`
	// PageSize
	PageSize uint64 `json:"pageSize"`
	// Status
	Status *uint32 `json:"status,optional"`
	// Rank
	Rank *uint32 `json:"rank,optional"`
	// InviteCode
	InviteCode *string `json:"inviteCode,optional"`
	// InvitedCode
	InvitedCode *string `json:"invitedCode,optional"`
	// 邮箱
	Email *string `json:"email,optional"`
}

// ListBannerResp
// swagger:model ListPlayerResp
type ListPlayerResp struct {
	BaseDataInfo
	// Data
	Data ListPlayerInfo `json:"data"`
}

// swagger:model ListPlayerInfo
type ListPlayerInfo struct {
	BaseListInfo
	Data []PlayerInfo `json:"data"`
}

// BannerInfo
// swagger:model BannerInfo
type BannerInfo struct {
	// Id
	Id uint64 `json:"id"`
	// CreatedAt
	CreatedAt int64 `json:"createdAt"`
	// UpdatedAt
	UpdatedAt int64 `json:"updatedAt"`
	// Status
	Status uint32 `json:"status"`
	// Endpoint
	Endpoint string `json:"endpoint"`
	// Module
	Module string `json:"module"`
	// FileType
	FileType uint32 `json:"fileType"`
	// Path
	Path string `json:"path"`
	// JumpUrl
	JumpUrl string    `json:"jumpUrl"`
	File    *FileInfo `json:"file"`
}

// CreateBannerReq
// swagger:model CreateBannerReq
type CreateBannerReq struct {
	// Status
	Status uint32 `json:"status"`
	// Endpoint
	Endpoint string `json:"endpoint"`
	// Module
	Module string `json:"module"`
	// FileType
	FileType uint32 `json:"fileType"`
	// Path
	Path string `json:"path"`
	// JumpUrl
	JumpUrl string `json:"jumpUrl"`
	// FileId
	FileId string `json:"fileId"`
}

// UpdateBannerReq
// swagger:model UpdateBannerReq
type UpdateBannerReq struct {
	// Id
	Id uint64 `json:"id"`
	// Status
	Status *uint32 `json:"status,optional"`
	// Endpoint
	Endpoint *string `json:"endpoint,optional"`
	// Module
	Module *string `json:"module,optional"`
	// FileType
	FileType *uint32 `json:"fileType,optional"`
	// Path
	Path *string `json:"path,optional"`
	// JumpUrl
	JumpUrl *string `json:"jumpUrl,optional"`
	// FileId
	FileId *string `json:"fileId,optional"`
}

// DeleteBannerReq
// swagger:model DeleteBannerReq
type DeleteBannerReq struct {
	// Id
	Id uint64 `json:"id"`
}

// FindBannerReq
// swagger:model FindBannerReq
type FindBannerReq struct {
	// Id
	Id uint64 `json:"id"`
}

// swagger:model FindBannerResp
type FindBannerResp struct {
	BaseMsgResp
	// Data
	Data BannerInfo `json:"data"`
}

// ListBannerReq
// swagger:model ListBannerReq
type ListBannerReq struct {
	// Page
	Page uint64 `json:"page"`
	// PageSize
	PageSize uint64 `json:"pageSize"`
	// Status
	Status *uint32 `json:"status,optional"`
	// Endpoint
	Endpoint *string `json:"endpoint,optional"`
	// Module
	Module *string `json:"module,optional"`
	// FileType
	FileType *uint32 `json:"fileType,optional"`
	// Path
	Path *string `json:"path,optional"`
	// JumpUrl
	JumpUrl *string `json:"jumpUrl,optional"`
}

// ListBannerResp
// swagger:model ListBannerResp
type ListBannerResp struct {
	BaseDataInfo
	// Data
	Data ListBannerInfo `json:"data"`
}

// swagger:model ListBannerInfo
type ListBannerInfo struct {
	BaseListInfo
	Data []BannerInfo `json:"data"`
}

// OrderInfo
// swagger:model OrderInfo
type OrderInfo struct {
	// Id
	Id uint64 `json:"id"`
	// CreatedAt
	CreatedAt int64 `json:"createdAt"`
	// UpdatedAt
	UpdatedAt int64 `json:"updatedAt"`
	// Status
	Status uint32 `json:"status"`
	// StatusDesc
	StatusDesc string `json:"statusDesc"`
	// Type
	Type string `json:"type"`
	// Code
	Code string `json:"code"`
	// Num
	Num float64 `json:"num"`
	// ToAddress
	ToAddress string `json:"toAddress"`
	// FromAddress
	FromAddress string `json:"fromAddress"`
	// TransactionId
	TransactionId string `json:"transactionId"`
}

// CreateOrderReq
// swagger:model CreateOrderReq
type CreateOrderReq struct {
	// Status
	Status uint32 `json:"status"`
	// Num
	Num float64 `json:"num"`
	// Type
	Type string `json:"type"`
	// ToAddress
	ToAddress string `json:"toAddress"`
	// FromAddress
	FromAddress string `json:"fromAddress"`
}

// UpdateOrderReq
// swagger:model UpdateOrderReq
type UpdateOrderReq struct {
	// Id
	Id uint64 `json:"id"`
	// Status
	Status *uint32 `json:"status,optional"`
	// Num
	Num *float64 `json:"num,optional"`
	// ToAddress
	ToAddress *string `json:"toAddress,optional"`
	// FromAddress
	FromAddress *string `json:"fromAddress,optional"`
	// TransactionId
	TransactionId *string `json:"transactionId,optional"`
}

// DeleteOrderReq
// swagger:model DeleteOrderReq
type DeleteOrderReq struct {
	// Id
	Id uint64 `json:"id"`
}

// FindOrderReq
// swagger:model FindOrderReq
type FindOrderReq struct {
	// Id
	Id uint64 `json:"id"`
}

// swagger:model FindOrderResp
type FindOrderResp struct {
	BaseDataInfo
	// Data
	Data OrderInfo `json:"data"`
}

// ListOrderReq
// swagger:model ListOrderReq
type ListOrderReq struct {
	// Page
	Page uint64 `json:"page"`
	// PageSize
	PageSize uint64 `json:"pageSize"`
	// Status
	Status *uint32 `json:"status,optional"`
	// Type
	Type *string `json:"type,optional"`
	// Code
	Code *string `json:"code,optional"`
	// Num
	Num *float64 `json:"num,optional"`
	// ToAddress
	ToAddress *string `json:"toAddress,optional"`
	// FromAddress
	FromAddress *string `json:"fromAddress,optional"`
}

// ListBannerResp
// swagger:model ListOrderResp
type ListOrderResp struct {
	BaseDataInfo
	// Data
	Data ListOrderInfo `json:"data"`
}

// swagger:model ListOrderInfo
type ListOrderInfo struct {
	BaseListInfo
	Data []OrderInfo `json:"data"`
}

// 统揽数据
// swagger:model OverviewInfo
type OverviewInfo struct {
	// 今日参与游戏玩家数
	TodayParticipateCount uint64 `json:"todayParticipateCount"`
	// 今日新增游戏玩家数
	TodayNewPlayerCount uint64 `json:"todayNewPlayerCount"`
	// 今日游戏开局次数
	TodayRoundCount uint64 `json:"todayRoundCount"`
	// 今日累积吃羊数
	TodayEatCount uint64 `json:"todayEatCount"`
	// 今日平台收益
	TodayPlatformProfit uint64 `json:"todayPlatformProfit"`
	// 平台累积收益（可按月查询明细）
	TotalPlatformProfit uint64 `json:"totalPlatformProfit"`
	// 平台累积人数
	TotalPlayerCount uint64 `json:"totalPlayerCount"`
}

// 获取总揽数据请求体
// swagger:model GetOverviewReq
type GetOverviewReq struct {
	StartTime *int64 `json:"startTime,optional"`
	EndTime   *int64 `json:"endTime,optional"`
}

// 获取总揽数据响应体
// swagger:model GetOverviewResp
type GetOverviewResp struct {
	BaseDataInfo
	Data OverviewInfo `json:"data"`
}

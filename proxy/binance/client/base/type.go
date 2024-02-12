package base

// SideType define side type of order
type SideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderStatusType define order status type
type OrderStatusType string

// SymbolType define symbol type
type SymbolType string

// SymbolStatusType define symbol status type
type SymbolStatusType string

// SymbolFilterType define symbol filter type
type SymbolFilterType string

// UserDataEventType define spot user data event type
type UserDataEventType string

// MarginTransferType define margin transfer type
type MarginTransferType int

// MarginLoanStatusType define margin loan status type
type MarginLoanStatusType string

// MarginRepayStatusType define margin repay status type
type MarginRepayStatusType string

// FuturesTransferStatusType define futures transfer status type
type FuturesTransferStatusType string

// SideEffectType define side effect type for orders
type SideEffectType string

// FuturesTransferType define futures transfer type
type FuturesTransferType int

// TransactionType define transaction type
type TransactionType string

// LendingType define the type of lending (flexible saving, activity, ...)
type LendingType string

// StakingProduct define the staking product (locked staking, flexible defi staking, locked defi staking, ...)
type StakingProduct string

// StakingTransactionType define the staking transaction type (subscription, redemption, interest)
type StakingTransactionType string

// LiquidityOperationType define the type of adding/removing liquidity to a liquidity pool(COMBINATION, SINGLE)
type LiquidityOperationType string

// SwappingStatus define the status of swap when querying the swap history
type SwappingStatus int

// LiquidityRewardType define the type of reward we'd claim
type LiquidityRewardType int

// RewardClaimStatus define the status of claiming a reward
type RewardClaimStatus int

// RateLimitType define the rate limitation types
// see https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md#enum-definitions
type RateLimitType string

// RateLimitInterval define the rate limitation intervals
type RateLimitInterval string

// AccountType define the account types
type AccountType string

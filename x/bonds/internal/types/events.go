package types

const (
	EventTypeCreateBond         = "create_bond"
	EventTypeEditBond           = "edit_bond"
	EventTypeSetNextAlpha       = "set_next_alpha"
	EventTypeEditAlphaSuccess   = "edit_alpha_success"
	EventTypeEditAlphaFailed    = "edit_alpha_failed"
	EventTypeInitSwapper        = "init_swapper"
	EventTypeBuy                = "buy"
	EventTypeSell               = "sell"
	EventTypeSwap               = "swap"
	EventTypeMakeOutcomePayment = "make_outcome_payment"
	EventTypeWithdrawShare      = "withdraw_share"
	EventTypeOrderCancel        = "order_cancel"
	EventTypeOrderFulfill       = "order_fulfill"
	EventTypeStateChange        = "state_change"

	AttributeKeyBondDid                = "bond_did"
	AttributeKeyToken                  = "token"
	AttributeKeyName                   = "name"
	AttributeKeyDescription            = "description"
	AttributeKeyFunctionType           = "function_type"
	AttributeKeyFunctionParameters     = "function_parameters"
	AttributeKeyReserveTokens          = "reserve_tokens"
	AttributeKeyTxFeePercentage        = "tx_fee_percentage"
	AttributeKeyExitFeePercentage      = "exit_fee_percentage"
	AttributeKeyFeeAddress             = "fee_address"
	AttributeKeyMaxSupply              = "max_supply"
	AttributeKeyOrderQuantityLimits    = "order_quantity_limits"
	AttributeKeySanityRate             = "sanity_rate"
	AttributeKeySanityMarginPercentage = "sanity_margin_percentage"
	AttributeKeyAllowSells             = "allow_sells"
	AttributeKeyBatchBlocks            = "batch_blocks"
	AttributeKeyOutcomePayment         = "outcome_payment"
	AttributeKeyState                  = "state"
	AttributeKeyMaxPrices              = "max_prices"
	AttributeKeySwapFromToken          = "from_token"
	AttributeKeySwapToToken            = "to_token"
	AttributeKeyOrderType              = "order_type"
	AttributeKeyAddress                = "address"
	AttributeKeyCancelReason           = "cancel_reason"
	AttributeKeyTokensMinted           = "tokens_minted"
	AttributeKeyTokensBurned           = "tokens_burned"
	AttributeKeyTokensSwapped          = "tokens_swapped"
	AttributeKeyChargedPrices          = "charged_prices"
	AttributeKeyChargedPricesReserve   = "charged_prices_of_which_reserve"
	AttributeKeyChargedPricesFunding   = "charged_prices_of_which_funding"
	AttributeKeyChargedFees            = "charged_fees"
	AttributeKeyReturnedToAddress      = "returned_to_address"
	AttributeKeyNewBondTokenBalance    = "new_bond_token_balance"
	AttributeKeyOldState               = "old_state"
	AttributeKeyNewState               = "new_state"
	AttributeKeyAlpha                  = "alpha"

	AttributeValueBuyOrder  = "buy"
	AttributeValueSellOrder = "sell"
	AttributeValueSwapOrder = "swap"
	AttributeValueCategory  = ModuleName
)

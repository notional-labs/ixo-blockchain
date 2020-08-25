package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryParams          = "queryParams"
	QueryPaymentTemplate = "queryPaymentTemplate"
	QueryPaymentContract = "queryPaymentContract"
	QuerySubscription    = "querySubscription"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryParams:
			return queryParams(ctx, k)
		case QueryPaymentTemplate:
			return queryPaymentTemplate(ctx, path[1:], k)
		case QueryPaymentContract:
			return queryPaymentContract(ctx, path[1:], k)
		case QuerySubscription:
			return querySubscription(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown payments query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInternal, "failed to marshal JSON")
	}

	return res, nil
}

func queryPaymentTemplate(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	templateId := path[0]

	template, err := k.GetPaymentTemplate(ctx, templateId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "payment template '%s' does not exist", templateId)
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, template)
	if err2 != nil {
		return nil, sdkerrors.Wrapf(types.ErrInternal, "failed to marshal JSON", err2.Error())
	}

	return res, nil
}

func queryPaymentContract(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	contractId := path[0]

	contract, err := k.GetPaymentContract(ctx, contractId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "payment contract '%s' does not exist", contractId)
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, contract)
	if err2 != nil {
		return nil, sdkerrors.Wrapf(types.ErrInternal, "failed to marshal JSON", err2.Error())
	}

	return res, nil
}

func querySubscription(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	subscriptionId := path[0]

	subscription, err := k.GetSubscription(ctx, subscriptionId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "subscription '%s' does not exist", subscriptionId)
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, subscription)
	if err2 != nil {
		return nil, sdkerrors.Wrapf(types.ErrInternal, "failed to marshal JSON", err2.Error())
	}

	return res, nil
}

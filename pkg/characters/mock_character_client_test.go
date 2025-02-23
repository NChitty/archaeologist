package characters_test

import (
	"context"
	"io"

	"github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/mock"
)

type mockClient struct {
	mock.Mock
}

// ActionAcceptNewTaskMyNameActionTaskNewPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionAcceptNewTaskMyNameActionTaskNewPostWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.ActionAcceptNewTaskMyNameActionTaskNewPostResponse, error) {
	panic("unimplemented")
}

// ActionBuyBankExpansionMyNameActionBankBuyExpansionPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionBuyBankExpansionMyNameActionBankBuyExpansionPostWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.ActionBuyBankExpansionMyNameActionBankBuyExpansionPostResponse, error) {
	panic("unimplemented")
}

// ActionCompleteTaskMyNameActionTaskCompletePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionCompleteTaskMyNameActionTaskCompletePostWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.ActionCompleteTaskMyNameActionTaskCompletePostResponse, error) {
	panic("unimplemented")
}

// ActionCraftingMyNameActionCraftingPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionCraftingMyNameActionCraftingPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionCraftingMyNameActionCraftingPostResponse, error) {
	panic("unimplemented")
}

// ActionCraftingMyNameActionCraftingPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionCraftingMyNameActionCraftingPostWithResponse(ctx context.Context, name string, body client.CraftingSchema, reqEditors ...client.RequestEditorFn) (*client.ActionCraftingMyNameActionCraftingPostResponse, error) {
	panic("unimplemented")
}

// ActionDeleteItemMyNameActionDeletePostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionDeleteItemMyNameActionDeletePostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionDeleteItemMyNameActionDeletePostResponse, error) {
	panic("unimplemented")
}

// ActionDeleteItemMyNameActionDeletePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionDeleteItemMyNameActionDeletePostWithResponse(ctx context.Context, name string, body client.SimpleItemSchema, reqEditors ...client.RequestEditorFn) (*client.ActionDeleteItemMyNameActionDeletePostResponse, error) {
	panic("unimplemented")
}

// ActionDepositBankGoldMyNameActionBankDepositGoldPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionDepositBankGoldMyNameActionBankDepositGoldPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionDepositBankGoldMyNameActionBankDepositGoldPostResponse, error) {
	panic("unimplemented")
}

// ActionDepositBankGoldMyNameActionBankDepositGoldPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionDepositBankGoldMyNameActionBankDepositGoldPostWithResponse(ctx context.Context, name string, body client.DepositWithdrawGoldSchema, reqEditors ...client.RequestEditorFn) (*client.ActionDepositBankGoldMyNameActionBankDepositGoldPostResponse, error) {
	panic("unimplemented")
}

// ActionDepositBankMyNameActionBankDepositPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionDepositBankMyNameActionBankDepositPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionDepositBankMyNameActionBankDepositPostResponse, error) {
	panic("unimplemented")
}

// ActionDepositBankMyNameActionBankDepositPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionDepositBankMyNameActionBankDepositPostWithResponse(ctx context.Context, name string, body client.SimpleItemSchema, reqEditors ...client.RequestEditorFn) (*client.ActionDepositBankMyNameActionBankDepositPostResponse, error) {
	panic("unimplemented")
}

// ActionEquipItemMyNameActionEquipPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionEquipItemMyNameActionEquipPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionEquipItemMyNameActionEquipPostResponse, error) {
	panic("unimplemented")
}

// ActionEquipItemMyNameActionEquipPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionEquipItemMyNameActionEquipPostWithResponse(ctx context.Context, name string, body client.EquipSchema, reqEditors ...client.RequestEditorFn) (*client.ActionEquipItemMyNameActionEquipPostResponse, error) {
	panic("unimplemented")
}

// ActionFightMyNameActionFightPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionFightMyNameActionFightPostWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.ActionFightMyNameActionFightPostResponse, error) {
	panic("unimplemented")
}

// ActionGatheringMyNameActionGatheringPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionGatheringMyNameActionGatheringPostWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.ActionGatheringMyNameActionGatheringPostResponse, error) {
	panic("unimplemented")
}

// ActionGeBuyItemMyNameActionGrandexchangeBuyPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionGeBuyItemMyNameActionGrandexchangeBuyPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionGeBuyItemMyNameActionGrandexchangeBuyPostResponse, error) {
	panic("unimplemented")
}

// ActionGeBuyItemMyNameActionGrandexchangeBuyPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionGeBuyItemMyNameActionGrandexchangeBuyPostWithResponse(ctx context.Context, name string, body client.GEBuyOrderSchema, reqEditors ...client.RequestEditorFn) (*client.ActionGeBuyItemMyNameActionGrandexchangeBuyPostResponse, error) {
	panic("unimplemented")
}

// ActionGeCancelSellOrderMyNameActionGrandexchangeCancelPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionGeCancelSellOrderMyNameActionGrandexchangeCancelPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionGeCancelSellOrderMyNameActionGrandexchangeCancelPostResponse, error) {
	panic("unimplemented")
}

// ActionGeCancelSellOrderMyNameActionGrandexchangeCancelPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionGeCancelSellOrderMyNameActionGrandexchangeCancelPostWithResponse(ctx context.Context, name string, body client.GECancelOrderSchema, reqEditors ...client.RequestEditorFn) (*client.ActionGeCancelSellOrderMyNameActionGrandexchangeCancelPostResponse, error) {
	panic("unimplemented")
}

// ActionGeCreateSellOrderMyNameActionGrandexchangeSellPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionGeCreateSellOrderMyNameActionGrandexchangeSellPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionGeCreateSellOrderMyNameActionGrandexchangeSellPostResponse, error) {
	panic("unimplemented")
}

// ActionGeCreateSellOrderMyNameActionGrandexchangeSellPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionGeCreateSellOrderMyNameActionGrandexchangeSellPostWithResponse(ctx context.Context, name string, body client.GEOrderCreationrSchema, reqEditors ...client.RequestEditorFn) (*client.ActionGeCreateSellOrderMyNameActionGrandexchangeSellPostResponse, error) {
	panic("unimplemented")
}

// ActionMoveMyNameActionMovePostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionMoveMyNameActionMovePostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionMoveMyNameActionMovePostResponse, error) {
	panic("unimplemented")
}

// ActionMoveMyNameActionMovePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionMoveMyNameActionMovePostWithResponse(ctx context.Context, name string, body client.DestinationSchema, reqEditors ...client.RequestEditorFn) (*client.ActionMoveMyNameActionMovePostResponse, error) {
	panic("unimplemented")
}

// ActionNpcBuyItemMyNameActionNpcBuyPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionNpcBuyItemMyNameActionNpcBuyPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionNpcBuyItemMyNameActionNpcBuyPostResponse, error) {
	panic("unimplemented")
}

// ActionNpcBuyItemMyNameActionNpcBuyPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionNpcBuyItemMyNameActionNpcBuyPostWithResponse(ctx context.Context, name string, body client.NpcMerchantBuySchema, reqEditors ...client.RequestEditorFn) (*client.ActionNpcBuyItemMyNameActionNpcBuyPostResponse, error) {
	panic("unimplemented")
}

// ActionNpcSellItemMyNameActionNpcSellPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionNpcSellItemMyNameActionNpcSellPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionNpcSellItemMyNameActionNpcSellPostResponse, error) {
	panic("unimplemented")
}

// ActionNpcSellItemMyNameActionNpcSellPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionNpcSellItemMyNameActionNpcSellPostWithResponse(ctx context.Context, name string, body client.NpcMerchantBuySchema, reqEditors ...client.RequestEditorFn) (*client.ActionNpcSellItemMyNameActionNpcSellPostResponse, error) {
	panic("unimplemented")
}

// ActionRecyclingMyNameActionRecyclingPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionRecyclingMyNameActionRecyclingPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionRecyclingMyNameActionRecyclingPostResponse, error) {
	panic("unimplemented")
}

// ActionRecyclingMyNameActionRecyclingPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionRecyclingMyNameActionRecyclingPostWithResponse(ctx context.Context, name string, body client.RecyclingSchema, reqEditors ...client.RequestEditorFn) (*client.ActionRecyclingMyNameActionRecyclingPostResponse, error) {
	panic("unimplemented")
}

// ActionRestMyNameActionRestPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionRestMyNameActionRestPostWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.ActionRestMyNameActionRestPostResponse, error) {
	panic("unimplemented")
}

// ActionTaskCancelMyNameActionTaskCancelPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionTaskCancelMyNameActionTaskCancelPostWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.ActionTaskCancelMyNameActionTaskCancelPostResponse, error) {
	panic("unimplemented")
}

// ActionTaskExchangeMyNameActionTaskExchangePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionTaskExchangeMyNameActionTaskExchangePostWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.ActionTaskExchangeMyNameActionTaskExchangePostResponse, error) {
	panic("unimplemented")
}

// ActionTaskTradeMyNameActionTaskTradePostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionTaskTradeMyNameActionTaskTradePostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionTaskTradeMyNameActionTaskTradePostResponse, error) {
	panic("unimplemented")
}

// ActionTaskTradeMyNameActionTaskTradePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionTaskTradeMyNameActionTaskTradePostWithResponse(ctx context.Context, name string, body client.SimpleItemSchema, reqEditors ...client.RequestEditorFn) (*client.ActionTaskTradeMyNameActionTaskTradePostResponse, error) {
	panic("unimplemented")
}

// ActionUnequipItemMyNameActionUnequipPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionUnequipItemMyNameActionUnequipPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionUnequipItemMyNameActionUnequipPostResponse, error) {
	panic("unimplemented")
}

// ActionUnequipItemMyNameActionUnequipPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionUnequipItemMyNameActionUnequipPostWithResponse(ctx context.Context, name string, body client.UnequipSchema, reqEditors ...client.RequestEditorFn) (*client.ActionUnequipItemMyNameActionUnequipPostResponse, error) {
	panic("unimplemented")
}

// ActionUseItemMyNameActionUsePostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionUseItemMyNameActionUsePostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionUseItemMyNameActionUsePostResponse, error) {
	panic("unimplemented")
}

// ActionUseItemMyNameActionUsePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionUseItemMyNameActionUsePostWithResponse(ctx context.Context, name string, body client.SimpleItemSchema, reqEditors ...client.RequestEditorFn) (*client.ActionUseItemMyNameActionUsePostResponse, error) {
	panic("unimplemented")
}

// ActionWithdrawBankGoldMyNameActionBankWithdrawGoldPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionWithdrawBankGoldMyNameActionBankWithdrawGoldPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionWithdrawBankGoldMyNameActionBankWithdrawGoldPostResponse, error) {
	panic("unimplemented")
}

// ActionWithdrawBankGoldMyNameActionBankWithdrawGoldPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionWithdrawBankGoldMyNameActionBankWithdrawGoldPostWithResponse(ctx context.Context, name string, body client.DepositWithdrawGoldSchema, reqEditors ...client.RequestEditorFn) (*client.ActionWithdrawBankGoldMyNameActionBankWithdrawGoldPostResponse, error) {
	panic("unimplemented")
}

// ActionWithdrawBankMyNameActionBankWithdrawPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionWithdrawBankMyNameActionBankWithdrawPostWithBodyWithResponse(ctx context.Context, name string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ActionWithdrawBankMyNameActionBankWithdrawPostResponse, error) {
	panic("unimplemented")
}

// ActionWithdrawBankMyNameActionBankWithdrawPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ActionWithdrawBankMyNameActionBankWithdrawPostWithResponse(ctx context.Context, name string, body client.SimpleItemSchema, reqEditors ...client.RequestEditorFn) (*client.ActionWithdrawBankMyNameActionBankWithdrawPostResponse, error) {
	panic("unimplemented")
}

// ChangePasswordMyChangePasswordPostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ChangePasswordMyChangePasswordPostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.ChangePasswordMyChangePasswordPostResponse, error) {
	panic("unimplemented")
}

// ChangePasswordMyChangePasswordPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) ChangePasswordMyChangePasswordPostWithResponse(ctx context.Context, body client.ChangePassword, reqEditors ...client.RequestEditorFn) (*client.ChangePasswordMyChangePasswordPostResponse, error) {
	panic("unimplemented")
}

// CreateAccountAccountsCreatePostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) CreateAccountAccountsCreatePostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.CreateAccountAccountsCreatePostResponse, error) {
	panic("unimplemented")
}

// CreateAccountAccountsCreatePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) CreateAccountAccountsCreatePostWithResponse(ctx context.Context, body client.AddAccountSchema, reqEditors ...client.RequestEditorFn) (*client.CreateAccountAccountsCreatePostResponse, error) {
	panic("unimplemented")
}

// CreateCharacterCharactersCreatePostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) CreateCharacterCharactersCreatePostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.CreateCharacterCharactersCreatePostResponse, error) {
	panic("unimplemented")
}

// CreateCharacterCharactersCreatePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) CreateCharacterCharactersCreatePostWithResponse(ctx context.Context, body client.AddCharacterSchema, reqEditors ...client.RequestEditorFn) (*client.CreateCharacterCharactersCreatePostResponse, error) {
	panic("unimplemented")
}

// DeleteCharacterCharactersDeletePostWithBodyWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) DeleteCharacterCharactersDeletePostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*client.DeleteCharacterCharactersDeletePostResponse, error) {
	panic("unimplemented")
}

// DeleteCharacterCharactersDeletePostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) DeleteCharacterCharactersDeletePostWithResponse(ctx context.Context, body client.DeleteCharacterSchema, reqEditors ...client.RequestEditorFn) (*client.DeleteCharacterCharactersDeletePostResponse, error) {
	panic("unimplemented")
}

// GenerateTokenTokenPostWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GenerateTokenTokenPostWithResponse(ctx context.Context, reqEditors ...client.RequestEditorFn) (*client.GenerateTokenTokenPostResponse, error) {
	panic("unimplemented")
}

// GetAccountAccountsAccountGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAccountAccountsAccountGetWithResponse(ctx context.Context, account string, reqEditors ...client.RequestEditorFn) (*client.GetAccountAccountsAccountGetResponse, error) {
	panic("unimplemented")
}

// GetAccountAchievementsAccountsAccountAchievementsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAccountAchievementsAccountsAccountAchievementsGetWithResponse(ctx context.Context, account string, params *client.GetAccountAchievementsAccountsAccountAchievementsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAccountAchievementsAccountsAccountAchievementsGetResponse, error) {
	panic("unimplemented")
}

// GetAccountDetailsMyDetailsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAccountDetailsMyDetailsGetWithResponse(ctx context.Context, reqEditors ...client.RequestEditorFn) (*client.GetAccountDetailsMyDetailsGetResponse, error) {
	panic("unimplemented")
}

// GetAccountsLeaderboardLeaderboardAccountsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAccountsLeaderboardLeaderboardAccountsGetWithResponse(ctx context.Context, params *client.GetAccountsLeaderboardLeaderboardAccountsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAccountsLeaderboardLeaderboardAccountsGetResponse, error) {
	panic("unimplemented")
}

// GetAchievementAchievementsCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAchievementAchievementsCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetAchievementAchievementsCodeGetResponse, error) {
	panic("unimplemented")
}

// GetAllAchievementsAchievementsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllAchievementsAchievementsGetWithResponse(ctx context.Context, params *client.GetAllAchievementsAchievementsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllAchievementsAchievementsGetResponse, error) {
	panic("unimplemented")
}

// GetAllActiveEventsEventsActiveGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllActiveEventsEventsActiveGetWithResponse(ctx context.Context, params *client.GetAllActiveEventsEventsActiveGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllActiveEventsEventsActiveGetResponse, error) {
	panic("unimplemented")
}

// GetAllBadgesBadgesGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllBadgesBadgesGetWithResponse(ctx context.Context, params *client.GetAllBadgesBadgesGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllBadgesBadgesGetResponse, error) {
	panic("unimplemented")
}

// GetAllCharactersLogsMyLogsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllCharactersLogsMyLogsGetWithResponse(ctx context.Context, params *client.GetAllCharactersLogsMyLogsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllCharactersLogsMyLogsGetResponse, error) {
	panic("unimplemented")
}

// GetAllEffectsEffectsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllEffectsEffectsGetWithResponse(ctx context.Context, params *client.GetAllEffectsEffectsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllEffectsEffectsGetResponse, error) {
	panic("unimplemented")
}

// GetAllEventsEventsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllEventsEventsGetWithResponse(ctx context.Context, params *client.GetAllEventsEventsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllEventsEventsGetResponse, error) {
	panic("unimplemented")
}

// GetAllItemsItemsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllItemsItemsGetWithResponse(ctx context.Context, params *client.GetAllItemsItemsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllItemsItemsGetResponse, error) {
	panic("unimplemented")
}

// GetAllMapsMapsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllMapsMapsGetWithResponse(ctx context.Context, params *client.GetAllMapsMapsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllMapsMapsGetResponse, error) {
	panic("unimplemented")
}

// GetAllMonstersMonstersGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllMonstersMonstersGetWithResponse(ctx context.Context, params *client.GetAllMonstersMonstersGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllMonstersMonstersGetResponse, error) {
	panic("unimplemented")
}

// GetAllNpcsNpcsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllNpcsNpcsGetWithResponse(ctx context.Context, params *client.GetAllNpcsNpcsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllNpcsNpcsGetResponse, error) {
	panic("unimplemented")
}

// GetAllResourcesResourcesGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllResourcesResourcesGetWithResponse(ctx context.Context, params *client.GetAllResourcesResourcesGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllResourcesResourcesGetResponse, error) {
	panic("unimplemented")
}

// GetAllTasksRewardsTasksRewardsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllTasksRewardsTasksRewardsGetWithResponse(ctx context.Context, params *client.GetAllTasksRewardsTasksRewardsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllTasksRewardsTasksRewardsGetResponse, error) {
	panic("unimplemented")
}

// GetAllTasksTasksListGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetAllTasksTasksListGetWithResponse(ctx context.Context, params *client.GetAllTasksTasksListGetParams, reqEditors ...client.RequestEditorFn) (*client.GetAllTasksTasksListGetResponse, error) {
	panic("unimplemented")
}

// GetBadgeBadgesCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetBadgeBadgesCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetBadgeBadgesCodeGetResponse, error) {
	panic("unimplemented")
}

// GetBankDetailsMyBankGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetBankDetailsMyBankGetWithResponse(ctx context.Context, reqEditors ...client.RequestEditorFn) (*client.GetBankDetailsMyBankGetResponse, error) {
	panic("unimplemented")
}

// GetBankItemsMyBankItemsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetBankItemsMyBankItemsGetWithResponse(ctx context.Context, params *client.GetBankItemsMyBankItemsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetBankItemsMyBankItemsGetResponse, error) {
	panic("unimplemented")
}

// GetCharacterCharactersNameGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetCharacterCharactersNameGetWithResponse(ctx context.Context, name string, reqEditors ...client.RequestEditorFn) (*client.GetCharacterCharactersNameGetResponse, error) {
  args := m.Called(ctx, name, reqEditors)
  return args.Get(0).(*client.GetCharacterCharactersNameGetResponse), args.Error(1)
}

// GetCharactersLeaderboardLeaderboardCharactersGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetCharactersLeaderboardLeaderboardCharactersGetWithResponse(ctx context.Context, params *client.GetCharactersLeaderboardLeaderboardCharactersGetParams, reqEditors ...client.RequestEditorFn) (*client.GetCharactersLeaderboardLeaderboardCharactersGetResponse, error) {
	panic("unimplemented")
}

// GetEffectEffectsCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetEffectEffectsCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetEffectEffectsCodeGetResponse, error) {
	panic("unimplemented")
}

// GetGeSellHistoryGrandexchangeHistoryCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetGeSellHistoryGrandexchangeHistoryCodeGetWithResponse(ctx context.Context, code string, params *client.GetGeSellHistoryGrandexchangeHistoryCodeGetParams, reqEditors ...client.RequestEditorFn) (*client.GetGeSellHistoryGrandexchangeHistoryCodeGetResponse, error) {
	panic("unimplemented")
}

// GetGeSellHistoryMyGrandexchangeHistoryGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetGeSellHistoryMyGrandexchangeHistoryGetWithResponse(ctx context.Context, params *client.GetGeSellHistoryMyGrandexchangeHistoryGetParams, reqEditors ...client.RequestEditorFn) (*client.GetGeSellHistoryMyGrandexchangeHistoryGetResponse, error) {
	panic("unimplemented")
}

// GetGeSellOrderGrandexchangeOrdersIdGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetGeSellOrderGrandexchangeOrdersIdGetWithResponse(ctx context.Context, id string, reqEditors ...client.RequestEditorFn) (*client.GetGeSellOrderGrandexchangeOrdersIdGetResponse, error) {
	panic("unimplemented")
}

// GetGeSellOrdersGrandexchangeOrdersGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetGeSellOrdersGrandexchangeOrdersGetWithResponse(ctx context.Context, params *client.GetGeSellOrdersGrandexchangeOrdersGetParams, reqEditors ...client.RequestEditorFn) (*client.GetGeSellOrdersGrandexchangeOrdersGetResponse, error) {
	panic("unimplemented")
}

// GetGeSellOrdersMyGrandexchangeOrdersGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetGeSellOrdersMyGrandexchangeOrdersGetWithResponse(ctx context.Context, params *client.GetGeSellOrdersMyGrandexchangeOrdersGetParams, reqEditors ...client.RequestEditorFn) (*client.GetGeSellOrdersMyGrandexchangeOrdersGetResponse, error) {
	panic("unimplemented")
}

// GetItemItemsCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetItemItemsCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetItemItemsCodeGetResponse, error) {
	panic("unimplemented")
}

// GetMapMapsXYGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetMapMapsXYGetWithResponse(ctx context.Context, x int, y int, reqEditors ...client.RequestEditorFn) (*client.GetMapMapsXYGetResponse, error) {
	panic("unimplemented")
}

// GetMonsterMonstersCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetMonsterMonstersCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetMonsterMonstersCodeGetResponse, error) {
	panic("unimplemented")
}

// GetMyCharactersMyCharactersGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetMyCharactersMyCharactersGetWithResponse(ctx context.Context, reqEditors ...client.RequestEditorFn) (*client.GetMyCharactersMyCharactersGetResponse, error) {
	panic("unimplemented")
}

// GetNpcItemsNpcsCodeItemsGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetNpcItemsNpcsCodeItemsGetWithResponse(ctx context.Context, code string, params *client.GetNpcItemsNpcsCodeItemsGetParams, reqEditors ...client.RequestEditorFn) (*client.GetNpcItemsNpcsCodeItemsGetResponse, error) {
	panic("unimplemented")
}

// GetNpcNpcsCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetNpcNpcsCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetNpcNpcsCodeGetResponse, error) {
	panic("unimplemented")
}

// GetResourceResourcesCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetResourceResourcesCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetResourceResourcesCodeGetResponse, error) {
	panic("unimplemented")
}

// GetStatusGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetStatusGetWithResponse(ctx context.Context, reqEditors ...client.RequestEditorFn) (*client.GetStatusGetResponse, error) {
	panic("unimplemented")
}

// GetTaskTasksListCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetTaskTasksListCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetTaskTasksListCodeGetResponse, error) {
	panic("unimplemented")
}

// GetTasksRewardTasksRewardsCodeGetWithResponse implements client.ClientWithResponsesInterface.
func (m *mockClient) GetTasksRewardTasksRewardsCodeGetWithResponse(ctx context.Context, code string, reqEditors ...client.RequestEditorFn) (*client.GetTasksRewardTasksRewardsCodeGetResponse, error) {
	panic("unimplemented")
}

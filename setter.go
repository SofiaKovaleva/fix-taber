func (AppModuleBasic) Name() string {
	return types.ModuleName
}
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (suite *IntegrationTestSuite) TestHasEnoughBalance() {
	k := suite.k
	bk := suite.bankKeeper
	ctx := suite.ctx
	require := suite.Require()

		isEnough := k.HasEnoughBalance(ctx, addr, coin)

	require.True(isEnough)
}
//dimmer
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
	if err != nil {
		return
	}
}

func (k Keeper) UpdateCookbook(ctx sdk.Context, cookbook types.Cookbook, prevAddr sdk.AccAddress) {
	k.removeCookbookFromAddress(ctx, cookbook.Id, prevAddr)
	k.SetCookbook(ctx, cookbook)
}

type AppModule struct {
	AppModuleBasic

	keeper     keeper.Keeper
	bankKeeper types.BankKeeper
}

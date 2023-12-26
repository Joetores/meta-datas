func (suite *IntegrationTestSuite) TestHasEnoughBalance() {
	k := suite.k
	bk := suite.bankKeeper
	ctx := suite.ctx
	require := suite.Require()

  	// CoinsIssuer module account has minter permissions
	isEnough := k.HasEnoughBalance(ctx, addr, coin)

	require.True(isEnough)
}

func (suite *IntegrationTestSuite) TestHasEnoughIBCDenomBalance() {
	k := suite.k
	bk := suite.bankKeeper
	ctx := suite.ctx
	require := suite.Require()
	require.False(isEnough)
}

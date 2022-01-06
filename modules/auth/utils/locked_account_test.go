package utils

import (
	"github.com/forbole/egldjuno/types"
)

func (suite *AuthProxyTestSuite) TestProxy_GetLockedAccountAddress() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)
	addresses := []string{
		"f1830cb81484659a",
	}
	lockedAccount, err := GetLockedAccount(addresses, height, proxy)
	suite.Require().NoError(err)

	expected := []types.LockedAccount{
		types.NewLockedAccount("f1830cb81484659a", "0x15e4b565057e6545"),
	}

	suite.Require().Equal(lockedAccount[0], expected[0])
}

func (suite *AuthProxyTestSuite) TestProxy_GetLockedAccountBalance() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)

	balance, err := getLockedTokenAccountBalance("f1830cb81484659a", height, proxy)
	suite.Require().NoError(err)

	suite.Require().Equal(uint64(0), balance)
}

func (suite *AuthProxyTestSuite) TestProxy_getLockedTokenAccountUnlockLimit() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)

	balance, err := getLockedTokenAccountUnlockLimit("f1830cb81484659a", height, proxy)
	suite.Require().NoError(err)

	suite.Require().Equal(uint64(100000), balance)
}

func (suite *AuthProxyTestSuite) TestProxy_getDelegatorNodeInfo() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)

	nodeInfo, err := getDelegatorNodeInfo("808b03495a0408bb", height, proxy)
	suite.Require().NoError(err)

	suite.Require().Equal("2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e", nodeInfo[0].NodeID)
	suite.Require().Equal(uint32(3905), nodeInfo[0].Id)

}

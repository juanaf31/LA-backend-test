package models

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteTestAccount struct {
    suite.Suite
	VariableAccountNumber string
	VariableCustName string
	VaribleBalance int
}
type SuiteTestTransfer struct {
    suite.Suite
	VariableReceiver string
	VaribleAmount int
}

func (suite *SuiteTestAccount) SetupTest() {
	suite.VariableAccountNumber = "abc"
	suite.VariableCustName = "def"
    suite.VaribleBalance = 5
}
func (suite *SuiteTestTransfer) SetupTest() {
	suite.VariableReceiver = "abc"
    suite.VaribleAmount = 5
}

func (suite *SuiteTestAccount) TestAccount() {
	suite.Equal(suite.VariableAccountNumber, "abc")
    suite.Equal(suite.VariableCustName, "def")
    suite.Equal(suite.VaribleBalance, 5)
}
func (suite *SuiteTestTransfer) TestTransfer() {
	suite.Equal(suite.VariableReceiver, "abc")
    suite.Equal(suite.VaribleAmount, 5)
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(SuiteTestAccount))
	suite.Run(t, new(SuiteTestTransfer))
}
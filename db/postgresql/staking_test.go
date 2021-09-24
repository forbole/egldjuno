package postgresql_test

import (
	"fmt"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/types"
)

func (suite *DbTestSuite) TestBigDipperDb_TotalStakeByType() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := []types.TotalStakeByType{
		types.NewTotalStakeByType(1, 2, 3),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveTotalStakeByType(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewTotalStakeByTypeRow(1, 2, 3)
	var outputs []dbtypes.TotalStakeByTypeRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM total_stake_by_type`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	fmt.Println(outputs[0])
	fmt.Println(expectedRow)

	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_StakeRequirements(){

	// ------------------------------
	  // --- Prepare the data
	  // ------------------------------

  
	/*  TODO: Prepare parameter    */
  
	input:=[]types.StakeRequirements{
		  types.NewStakeRequirements(1,2,3),
	  }
  
  
	  // ------------------------------
	  // --- Save the data
	  // ------------------------------
  
	err := suite.database.SaveStakeRequirements(input)
	suite.Require().NoError(err)
  
	// ------------------------------
	  // --- Verify the data
	  // ------------------------------
	  expectedRow := dbtypes.NewStakeRequirementsRow(1,2,3)
	var outputs []dbtypes.StakeRequirementsRow
	  err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM stake_requirements`)
	  suite.Require().NoError(err)
	  suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))
  
  }

  func (suite *DbTestSuite) TestBigDipperDb_WeeklyPayout(){

	// ------------------------------
	  // --- Prepare the data
	  // ------------------------------
  
	/*  TODO: Prepare parameter    */
  
	input:=types.NewWeeklyPayout(1,2)
  
  
	  // ------------------------------
	  // --- Save the data
	  // ------------------------------
  
	err := suite.database.SaveWeeklyPayout(input)
	suite.Require().NoError(err)
  
	// ------------------------------
	  // --- Verify the data
	  // ------------------------------
	  expectedRow := dbtypes.NewWeeklyPayoutRow(1,2)
	var outputs []dbtypes.WeeklyPayoutRow
	  err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM weekly_payout`)
	  suite.Require().NoError(err)
	  suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))
  
  }
  
package postgresql_test

import (
	"fmt"

	dbtypes "github.com/forbole/egldjuno/db/types"
	"github.com/forbole/egldjuno/types"
	"github.com/lib/pq"
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

func (suite *DbTestSuite) TestBigDipperDb_StakeRequirements() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.StakeRequirements{
		types.NewStakeRequirements(1, 2, 3),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveStakeRequirements(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewStakeRequirementsRow(1, 2, 3)
	var outputs []dbtypes.StakeRequirementsRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM stake_requirements`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_WeeklyPayout() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := types.NewWeeklyPayout(1, 2)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveWeeklyPayout(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewWeeklyPayoutRow(1, 2)
	var outputs []dbtypes.WeeklyPayoutRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM weekly_payout`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_TotalStake() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewTotalStake(1, 2)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveTotalStake(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewTotalStakeRow(1, 2)
	var outputs []dbtypes.TotalStakeRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM total_stake`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))
}

func (suite *DbTestSuite) TestBigDipperDb_StakingTable() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewStakingTable(10, []string{"abc", "efg"})

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveStakingTable(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := []dbtypes.StakingTableRow{
		dbtypes.NewStakingTableRow("abc"),
		dbtypes.NewStakingTableRow("efg"),
	}
	var outputs []dbtypes.StakingTableRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM staking_table`)
	suite.Require().NoError(err)
	for i, row := range expectedRow {
		suite.Require().True(row.Equal(outputs[i]))
	}

}

func (suite *DbTestSuite) TestBigDipperDb_ProposedTable() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewProposedTable(10, []string{"abc", "def"})

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveProposedTable(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := []dbtypes.ProposedTableRow{
		dbtypes.NewProposedTableRow(10, "abc"),
		dbtypes.NewProposedTableRow(10, "def"),
	}
	var outputs []dbtypes.ProposedTableRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM proposed_table`)
	suite.Require().NoError(err)
	for i, rows := range expectedRow {
		suite.Require().True(rows.Equal(outputs[i]))
	}
}

func (suite *DbTestSuite) insertCurrentTable(height uint64, nodeId []string) {
	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewCurrentTable(10, nodeId)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveCurrentTable(input)
	suite.Require().NoError(err)
}

func (suite *DbTestSuite) TestBigDipperDb_CurrentTable() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewCurrentTable(10, []string([]string{"abc", "def"}))

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveCurrentTable(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRows := []dbtypes.CurrentTableRow{
		dbtypes.NewCurrentTableRow(10, "abc"),
		dbtypes.NewCurrentTableRow(10, "def"),
	}
	var outputs []dbtypes.CurrentTableRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM current_table`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 2, "should contain two rows")
	for i, row := range expectedRows {
		suite.Require().True(row.Equal(outputs[i]))
	}

}

func (suite *DbTestSuite) TestBigDipperDb_NodeTotalCommitment() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeTotalCommitment{
		types.NewNodeTotalCommitment("0x1", 100000008, 1),
	}

	err := suite.InsertIntoStakingTable(1, "0x1")
	suite.Require().NoError(err)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err = suite.database.SaveNodeTotalCommitment(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeTotalCommitmentRow("0x1", 100000008, 1)
	var outputs []dbtypes.NodeTotalCommitmentRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_total_commitment`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeTotalCommitmentWithoutDelegators() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeTotalCommitmentWithoutDelegators{
		types.NewNodeTotalCommitmentWithoutDelegators("0x1", 100000008, 1),
	}

	err := suite.InsertIntoStakingTable(1, "0x1")
	suite.Require().NoError(err)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err = suite.database.SaveNodeTotalCommitmentWithoutDelegators(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeTotalCommitmentWithoutDelegatorsRow("0x1", 100000008, 1)
	var outputs []dbtypes.NodeTotalCommitmentWithoutDelegatorsRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_total_commitment_without_delegators`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_SaveNodeInfosFromTable() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	id := "e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6"

	stakerNodeInfo := types.NewStakerNodeInfo(id, 5, "34.211.45.12:3569",
		"3fa19db960a86a1722a2d8ffa9563bd1e7d905c91536860c64f9e808ef88862639112a1e872d7eef93dd91207d07dd7c043e2a1e80077b109290682250429f1f",
		"87d827de3e1b3541c394dcbbb6d76d98e3a7710d6740d28122c468b83f41002625e7a5788cabfcd6ce76b188f7f60de614364d4ab2932dfe0ed6f2d602bd551606ea31045ca2ccde9658a175ccd73da859ab17e56ad81ca4f6ef982c5968a7cb",
		0, 20000000, 0, 0, 0, []uint32{1}, 0, 0, 0)

	input := []types.StakerNodeInfo{
		stakerNodeInfo,
	}

	err := suite.InsertIntoStakingTable(1, id)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err = suite.database.SaveNodeInfosFromTable(input, 1)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------

	expectedRow := dbtypes.NewNodeInfosFromTableRow(id, 5, "34.211.45.12:3569",
		"3fa19db960a86a1722a2d8ffa9563bd1e7d905c91536860c64f9e808ef88862639112a1e872d7eef93dd91207d07dd7c043e2a1e80077b109290682250429f1f",
		"87d827de3e1b3541c394dcbbb6d76d98e3a7710d6740d28122c468b83f41002625e7a5788cabfcd6ce76b188f7f60de614364d4ab2932dfe0ed6f2d602bd551606ea31045ca2ccde9658a175ccd73da859ab17e56ad81ca4f6ef982c5968a7cb",
		0, 20000000, 0, 0, 0, pq.Int32Array{1}, 0, 0, 0, 1)
	var outputs []dbtypes.NodeInfosFromTableRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_infos_from_table`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only 1 row")
	fmt.Println(outputs[0])
	suite.Require().True(expectedRow.Equal(outputs[0]))
}

func (suite *DbTestSuite) TestBigDipperDb_CutPercentage() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := types.NewCutPercentage(1, 1)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveCutPercentage(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewCutPercentageRow(1, 1)
	var outputs []dbtypes.CutPercentageRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM cut_percentage`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	fmt.Println(outputs[0])
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) InsertIntoStakingTable(height uint64, nodeId string) error {
	_, err := suite.database.Sqlx.Exec(
		`INSERT INTO staking_table(node_id) VALUES ($1)`, nodeId)
	return err
}
func (suite *DbTestSuite) TestBigDipperDb_DelegatorInfo() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */
	delegatorId := int64(8411)
	delegatorNodeId := "2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e"
	nodeInfo := types.NewDelegatorNodeInfo(uint32(delegatorId), delegatorNodeId, 0, 0, 0, 0, 0, 0)

	input := []types.DelegatorNodeInfo{
		nodeInfo,
	}

	err := suite.InsertIntoStakingTable(1, delegatorNodeId)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err = suite.database.SaveDelegatorInfo(input, 1)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewDelegatorInfoRow(8411, "2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e", 0, 0, 0, 0, 0, 0, 1)
	var outputs []dbtypes.DelegatorInfoRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM delegator_info`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))
}

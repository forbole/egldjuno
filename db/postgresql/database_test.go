package postgresql_test

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/forbole/egldjuno/types/config"

	database "github.com/forbole/egldjuno/db/postgresql"
	"github.com/forbole/egldjuno/types"

	juno "github.com/desmos-labs/juno/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/stretchr/testify/suite"

	_ "github.com/proullon/ramsql/driver"
)

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

type DbTestSuite struct {
	suite.Suite

	database *database.Db
}

func (suite *DbTestSuite) SetupTest() {
	// Create the codec
	codec := simapp.MakeTestEncodingConfig()

	// Build the database
	cfg := types.NewConfig(
		nil, nil, nil,
		config.NewDatabaseConfig(
			juno.NewDatabaseConfig(
				"bdjuno",
				"localhost",
				5433,
				"bdjuno",
				"password",
				"",
				"public",
				-1,
				-1,
			),
			true,
		),
		nil, nil, nil, nil,
	)

	db, err := database.Builder(cfg, &codec)
	suite.Require().NoError(err)

	bigDipperDb, ok := (db).(*database.Db)
	suite.Require().True(ok)

	// Delete the public schema
	_, err = bigDipperDb.Sql.Exec(`DROP SCHEMA public CASCADE;`)
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = bigDipperDb.Sql.Exec(`CREATE SCHEMA public;`)
	suite.Require().NoError(err)

	dirPath := path.Join(".", "schema")
	dir, err := ioutil.ReadDir(dirPath)
	suite.Require().NoError(err)

	for _, fileInfo := range dir {
		file, err := ioutil.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		suite.Require().NoError(err)

		commentsRegExp := regexp.MustCompile(`/\*.*\*/`)
		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err := bigDipperDb.Sql.Exec(commentsRegExp.ReplaceAllString(request, ""))
			suite.Require().NoError(err)
		}
	}

	suite.database = bigDipperDb
}

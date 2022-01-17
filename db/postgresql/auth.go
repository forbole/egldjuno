package postgresql

import (
	"fmt"

	"github.com/forbole/egldjuno/types"
)

func (db *Db) SaveAccount(account []types.Account) error {
	stmt := `INSERT INTO account(address,nonce,balance,root_hash,tx_count,scr_count,shard,developer_reward) VALUES `

	var params []interface{}

	for i, rows := range account {
		ai := i * 8
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5, ai+6, ai+7, ai+8)

		params = append(params, rows.Address, rows.Nonce, rows.Balance, rows.RootHash, rows.TxCount, rows.ScrCount, rows.Shard, rows.DeveloperReward)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

package impl

import (
	"context"
	"database/sql"
	"errors"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/mysqllib"
	"math/big"
	"time"
)

type IDGenRepositoryImpl struct {
}

func (repositoryImpl *IDGenRepositoryImpl) NextIDGet(ctx context.Context,
	tableName string) ([]byte, error) {
	tx, err := databaseInst.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	nextID, err := repositoryImpl.nextIDGetWithRetries(tx, ctx, tableName, 0)
	if err != nil {
		return nil, mysqllib.RollbackTx(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, mysqllib.RollbackTx(tx, err)
	}
	return nextID, nil
}

/*
  	The below function creates the nextID for the given table with multiple retries  for the case
	where two or more threads try to create id for a table for the first time i.e if a table name
	is not present in the id_generator table and multiple threads try to create the id at the same time,
	we get a race condition, this race condition is resolved by multiple retries

	The max retries are set to 10.
*/

const (
	maxRetries = 10
)

func (repositoryImpl *IDGenRepositoryImpl) nextIDGetWithRetries(tx *sql.Tx,
	ctx context.Context,
	tableName string,
	retryCount uint8) ([]byte, error) {
	if retryCount > maxRetries {
		return nil, errorlib.NewInternalServerError("retries-exceeded-for-table=" + tableName)
	}
	qry := "SELECT id FROM id_generator WHERE table_name = ? FOR UPDATE;"
	row := tx.QueryRowContext(ctx, qry, tableName)
	/*  row is always non-nil value */
	var idBytes []byte
	err := row.Scan(&idBytes)
	if errors.Is(err, sql.ErrNoRows) {
		return repositoryImpl.insertID(tx, ctx, tableName, retryCount)
	}
	if err != nil {
		return nil, err
	}

	nextIDBytes := getNextID(idBytes)
	err = repositoryImpl.updateNextID(tx, ctx, tableName, nextIDBytes)
	if err != nil {
		return nil, err
	}
	return nextIDBytes, nil
}

func (repositoryImpl *IDGenRepositoryImpl) insertID(tx *sql.Tx,
	ctx context.Context,
	tableName string,
	retryCount uint8) ([]byte, error) {
	id := new(big.Int)
	id = id.SetInt64(1)
	newID := id.Bytes()
	qry := `INSERT INTO id_generator (table_name, id) 
			VALUES (?, ?);
			`
	result, err := tx.ExecContext(ctx, qry, tableName, newID)
	if err != nil && mysqllib.IsConflictError(err) {
		/* Recursively Retry Again */
		time.Sleep(time.Duration(30) * time.Microsecond) // wait and retry
		return repositoryImpl.nextIDGetWithRetries(tx, ctx, tableName, retryCount+1)
	}
	if err != nil {
		return nil, err
	}
	affRows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affRows != 1 {
		return nil, errors.New("rows-affected!=1")
	}
	return newID, nil
}

func (repositoryImpl *IDGenRepositoryImpl) updateNextID(tx *sql.Tx,
	ctx context.Context,
	tableName string, newID []byte) error {
	qry := "UPDATE id_generator SET id = ? WHERE table_name =?;"
	result, err := tx.ExecContext(ctx, qry, newID, tableName)
	if err != nil {
		return err
	}
	affRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affRows != 1 {
		return errorlib.NewInternalServerError("affRows!=1")
	}
	return nil
}

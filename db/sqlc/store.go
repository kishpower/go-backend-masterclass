package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// provides capabilities to execute db queries and transaction
type Store struct {
	*Queries
	connPool *pgxpool.Pool
}

// new store creates new store
func NewStore(connPool *pgxpool.Pool) *Store {
	return &Store{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

// executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rberr := tx.Rollback(ctx); rberr != nil {
			return fmt.Errorf("tx error: %v , rb error: %v", err, rberr)
		}
		return err
	}

	return tx.Commit(ctx)
}

// contains all necessary input params of transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// contains the result of the Transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// to avoid deadlock , creating empty struct
var txKey = struct{}{}

// transferTx : perform a money transaction from one account to another.
// it creates transfer record , add account entries and updaate account balance within a single database transaction.
func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		fmt.Println(txName, "create entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Amount:        args.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		})

		if err != nil {
			return err
		}

		// TODO : update account balance
		fmt.Println(txName, "get account 1 for update")
		acc1, err := q.GetAccountForUpdate(ctx, args.FromAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 1")
		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      acc1.ID,
			Balance: acc1.Balance - args.Amount,
		})

		fmt.Println(txName, "get account 2 for update")
		acc2, err := q.GetAccountForUpdate(ctx, args.ToAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 2")
		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      acc2.ID,
			Balance: acc2.Balance + args.Amount,
		})

		return nil
	})

	return result, err
}

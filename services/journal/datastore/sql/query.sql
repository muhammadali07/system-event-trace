
-- name: getListTransaction :many
SELECT * FROM journal order by created_at asc;

-- name: getTransactionByCondition :one
SELECT * FROM journal
WHERE account_number_credit = $1 or account_number_debit = $2;

-- name: createJournalTransaction :one
INSERT INTO journal (
  id, transcation_date, account_number_credit, account_number_debit, 
  amount_credit, amount_debit, type_transaction, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;



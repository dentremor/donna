// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: debts.sql

package models

import (
	"context"
)

const createDebt = `-- name: CreateDebt :one
with contact as (
    select id
    from contacts
    where contacts.id = $1
        and namespace = $2
),
insertion as (
    insert into debts (amount, currency, description, contact_id)
    select $3,
        $4,
        $5,
        $1
    from contact
    where exists (
            select 1
            from contact
        )
    returning debts.id
)
select id
from insertion
`

type CreateDebtParams struct {
	ID          int32
	Namespace   string
	Amount      float64
	Currency    string
	Description string
}

func (q *Queries) CreateDebt(ctx context.Context, arg CreateDebtParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createDebt,
		arg.ID,
		arg.Namespace,
		arg.Amount,
		arg.Currency,
		arg.Description,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteDebtsForContact = `-- name: DeleteDebtsForContact :exec
delete from debts using contacts
where debts.contact_id = contacts.id
    and contacts.id = $1
    and contacts.namespace = $2
`

type DeleteDebtsForContactParams struct {
	ID        int32
	Namespace string
}

func (q *Queries) DeleteDebtsForContact(ctx context.Context, arg DeleteDebtsForContactParams) error {
	_, err := q.db.ExecContext(ctx, deleteDebtsForContact, arg.ID, arg.Namespace)
	return err
}

const getDebtAndContact = `-- name: GetDebtAndContact :one
select debts.id as debt_id,
    debts.amount,
    debts.currency,
    debts.description,
    contacts.id as contact_id,
    contacts.first_name,
    contacts.last_name
from contacts
    inner join debts on debts.contact_id = contacts.id
where contacts.id = $1
    and contacts.namespace = $2
    and debts.id = $3
`

type GetDebtAndContactParams struct {
	ID        int32
	Namespace string
	ID_2      int32
}

type GetDebtAndContactRow struct {
	DebtID      int32
	Amount      float64
	Currency    string
	Description string
	ContactID   int32
	FirstName   string
	LastName    string
}

func (q *Queries) GetDebtAndContact(ctx context.Context, arg GetDebtAndContactParams) (GetDebtAndContactRow, error) {
	row := q.db.QueryRowContext(ctx, getDebtAndContact, arg.ID, arg.Namespace, arg.ID_2)
	var i GetDebtAndContactRow
	err := row.Scan(
		&i.DebtID,
		&i.Amount,
		&i.Currency,
		&i.Description,
		&i.ContactID,
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}

const getDebts = `-- name: GetDebts :many
select debts.id,
    debts.amount,
    debts.currency,
    debts.description
from contacts
    right join debts on debts.contact_id = contacts.id
where contacts.id = $1
    and contacts.namespace = $2
`

type GetDebtsParams struct {
	ID        int32
	Namespace string
}

type GetDebtsRow struct {
	ID          int32
	Amount      float64
	Currency    string
	Description string
}

func (q *Queries) GetDebts(ctx context.Context, arg GetDebtsParams) ([]GetDebtsRow, error) {
	rows, err := q.db.QueryContext(ctx, getDebts, arg.ID, arg.Namespace)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDebtsRow
	for rows.Next() {
		var i GetDebtsRow
		if err := rows.Scan(
			&i.ID,
			&i.Amount,
			&i.Currency,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const settleDebt = `-- name: SettleDebt :exec
delete from debts using contacts
where debts.id = $3
    and debts.contact_id = contacts.id
    and contacts.id = $1
    and contacts.namespace = $2
`

type SettleDebtParams struct {
	ID        int32
	Namespace string
	ID_2      int32
}

func (q *Queries) SettleDebt(ctx context.Context, arg SettleDebtParams) error {
	_, err := q.db.ExecContext(ctx, settleDebt, arg.ID, arg.Namespace, arg.ID_2)
	return err
}

const updateDebt = `-- name: UpdateDebt :exec
update debts
set amount = $4,
    currency = $5,
    description = $6
from contacts
where contacts.id = $1
    and contacts.namespace = $2
    and debts.id = $3
    and debts.contact_id = contacts.id
`

type UpdateDebtParams struct {
	ID          int32
	Namespace   string
	ID_2        int32
	Amount      float64
	Currency    string
	Description string
}

func (q *Queries) UpdateDebt(ctx context.Context, arg UpdateDebtParams) error {
	_, err := q.db.ExecContext(ctx, updateDebt,
		arg.ID,
		arg.Namespace,
		arg.ID_2,
		arg.Amount,
		arg.Currency,
		arg.Description,
	)
	return err
}
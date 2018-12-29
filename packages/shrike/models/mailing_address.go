// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// MailingAddress is an object representing the database table.
type MailingAddress struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *mailingAddressR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mailingAddressL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MailingAddressColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// MailingAddressRels is where relationship names are stored.
var MailingAddressRels = struct {
}{}

// mailingAddressR is where relationships are stored.
type mailingAddressR struct {
}

// NewStruct creates a new relationship struct
func (*mailingAddressR) NewStruct() *mailingAddressR {
	return &mailingAddressR{}
}

// mailingAddressL is where Load methods for each relationship are stored.
type mailingAddressL struct{}

var (
	mailingAddressColumns               = []string{"id", "created_at", "updated_at"}
	mailingAddressColumnsWithoutDefault = []string{"created_at", "updated_at"}
	mailingAddressColumnsWithDefault    = []string{"id"}
	mailingAddressPrimaryKeyColumns     = []string{"id"}
)

type (
	// MailingAddressSlice is an alias for a slice of pointers to MailingAddress.
	// This should generally be used opposed to []MailingAddress.
	MailingAddressSlice []*MailingAddress
	// MailingAddressHook is the signature for custom MailingAddress hook methods
	MailingAddressHook func(context.Context, boil.ContextExecutor, *MailingAddress) error

	mailingAddressQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mailingAddressType                 = reflect.TypeOf(&MailingAddress{})
	mailingAddressMapping              = queries.MakeStructMapping(mailingAddressType)
	mailingAddressPrimaryKeyMapping, _ = queries.BindMapping(mailingAddressType, mailingAddressMapping, mailingAddressPrimaryKeyColumns)
	mailingAddressInsertCacheMut       sync.RWMutex
	mailingAddressInsertCache          = make(map[string]insertCache)
	mailingAddressUpdateCacheMut       sync.RWMutex
	mailingAddressUpdateCache          = make(map[string]updateCache)
	mailingAddressUpsertCacheMut       sync.RWMutex
	mailingAddressUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var mailingAddressBeforeInsertHooks []MailingAddressHook
var mailingAddressBeforeUpdateHooks []MailingAddressHook
var mailingAddressBeforeDeleteHooks []MailingAddressHook
var mailingAddressBeforeUpsertHooks []MailingAddressHook

var mailingAddressAfterInsertHooks []MailingAddressHook
var mailingAddressAfterSelectHooks []MailingAddressHook
var mailingAddressAfterUpdateHooks []MailingAddressHook
var mailingAddressAfterDeleteHooks []MailingAddressHook
var mailingAddressAfterUpsertHooks []MailingAddressHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MailingAddress) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MailingAddress) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MailingAddress) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MailingAddress) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MailingAddress) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MailingAddress) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MailingAddress) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MailingAddress) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MailingAddress) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range mailingAddressAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMailingAddressHook registers your hook function for all future operations.
func AddMailingAddressHook(hookPoint boil.HookPoint, mailingAddressHook MailingAddressHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		mailingAddressBeforeInsertHooks = append(mailingAddressBeforeInsertHooks, mailingAddressHook)
	case boil.BeforeUpdateHook:
		mailingAddressBeforeUpdateHooks = append(mailingAddressBeforeUpdateHooks, mailingAddressHook)
	case boil.BeforeDeleteHook:
		mailingAddressBeforeDeleteHooks = append(mailingAddressBeforeDeleteHooks, mailingAddressHook)
	case boil.BeforeUpsertHook:
		mailingAddressBeforeUpsertHooks = append(mailingAddressBeforeUpsertHooks, mailingAddressHook)
	case boil.AfterInsertHook:
		mailingAddressAfterInsertHooks = append(mailingAddressAfterInsertHooks, mailingAddressHook)
	case boil.AfterSelectHook:
		mailingAddressAfterSelectHooks = append(mailingAddressAfterSelectHooks, mailingAddressHook)
	case boil.AfterUpdateHook:
		mailingAddressAfterUpdateHooks = append(mailingAddressAfterUpdateHooks, mailingAddressHook)
	case boil.AfterDeleteHook:
		mailingAddressAfterDeleteHooks = append(mailingAddressAfterDeleteHooks, mailingAddressHook)
	case boil.AfterUpsertHook:
		mailingAddressAfterUpsertHooks = append(mailingAddressAfterUpsertHooks, mailingAddressHook)
	}
}

// One returns a single mailingAddress record from the query.
func (q mailingAddressQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MailingAddress, error) {
	o := &MailingAddress{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mailing_address")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MailingAddress records from the query.
func (q mailingAddressQuery) All(ctx context.Context, exec boil.ContextExecutor) (MailingAddressSlice, error) {
	var o []*MailingAddress

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MailingAddress slice")
	}

	if len(mailingAddressAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MailingAddress records in the query.
func (q mailingAddressQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mailing_address rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mailingAddressQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mailing_address exists")
	}

	return count > 0, nil
}

// MailingAddresses retrieves all the records using an executor.
func MailingAddresses(mods ...qm.QueryMod) mailingAddressQuery {
	mods = append(mods, qm.From("\"mailing_address\""))
	return mailingAddressQuery{NewQuery(mods...)}
}

// FindMailingAddress retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMailingAddress(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*MailingAddress, error) {
	mailingAddressObj := &MailingAddress{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mailing_address\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mailingAddressObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mailing_address")
	}

	return mailingAddressObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MailingAddress) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mailing_address provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mailingAddressColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mailingAddressInsertCacheMut.RLock()
	cache, cached := mailingAddressInsertCache[key]
	mailingAddressInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mailingAddressColumns,
			mailingAddressColumnsWithDefault,
			mailingAddressColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mailingAddressType, mailingAddressMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mailingAddressType, mailingAddressMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mailing_address\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mailing_address\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into mailing_address")
	}

	if !cached {
		mailingAddressInsertCacheMut.Lock()
		mailingAddressInsertCache[key] = cache
		mailingAddressInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MailingAddress.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MailingAddress) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mailingAddressUpdateCacheMut.RLock()
	cache, cached := mailingAddressUpdateCache[key]
	mailingAddressUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mailingAddressColumns,
			mailingAddressPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mailing_address, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mailing_address\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mailingAddressPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mailingAddressType, mailingAddressMapping, append(wl, mailingAddressPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update mailing_address row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mailing_address")
	}

	if !cached {
		mailingAddressUpdateCacheMut.Lock()
		mailingAddressUpdateCache[key] = cache
		mailingAddressUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mailingAddressQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mailing_address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mailing_address")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MailingAddressSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailingAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mailing_address\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mailingAddressPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mailingAddress slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mailingAddress")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MailingAddress) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mailing_address provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mailingAddressColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	mailingAddressUpsertCacheMut.RLock()
	cache, cached := mailingAddressUpsertCache[key]
	mailingAddressUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mailingAddressColumns,
			mailingAddressColumnsWithDefault,
			mailingAddressColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mailingAddressColumns,
			mailingAddressPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert mailing_address, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mailingAddressPrimaryKeyColumns))
			copy(conflict, mailingAddressPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mailing_address\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mailingAddressType, mailingAddressMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mailingAddressType, mailingAddressMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert mailing_address")
	}

	if !cached {
		mailingAddressUpsertCacheMut.Lock()
		mailingAddressUpsertCache[key] = cache
		mailingAddressUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MailingAddress record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MailingAddress) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MailingAddress provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mailingAddressPrimaryKeyMapping)
	sql := "DELETE FROM \"mailing_address\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mailing_address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mailing_address")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mailingAddressQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mailingAddressQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mailing_address")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mailing_address")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MailingAddressSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MailingAddress slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(mailingAddressBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailingAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mailing_address\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mailingAddressPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mailingAddress slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mailing_address")
	}

	if len(mailingAddressAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *MailingAddress) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMailingAddress(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MailingAddressSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MailingAddressSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailingAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mailing_address\".* FROM \"mailing_address\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mailingAddressPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MailingAddressSlice")
	}

	*o = slice

	return nil
}

// MailingAddressExists checks if the MailingAddress row exists.
func MailingAddressExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mailing_address\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mailing_address exists")
	}

	return exists, nil
}

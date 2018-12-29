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

// ActivityType is an object representing the database table.
type ActivityType struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *activityTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L activityTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ActivityTypeColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// ActivityTypeRels is where relationship names are stored.
var ActivityTypeRels = struct {
}{}

// activityTypeR is where relationships are stored.
type activityTypeR struct {
}

// NewStruct creates a new relationship struct
func (*activityTypeR) NewStruct() *activityTypeR {
	return &activityTypeR{}
}

// activityTypeL is where Load methods for each relationship are stored.
type activityTypeL struct{}

var (
	activityTypeColumns               = []string{"id", "created_at", "updated_at"}
	activityTypeColumnsWithoutDefault = []string{"created_at", "updated_at"}
	activityTypeColumnsWithDefault    = []string{"id"}
	activityTypePrimaryKeyColumns     = []string{"id"}
)

type (
	// ActivityTypeSlice is an alias for a slice of pointers to ActivityType.
	// This should generally be used opposed to []ActivityType.
	ActivityTypeSlice []*ActivityType
	// ActivityTypeHook is the signature for custom ActivityType hook methods
	ActivityTypeHook func(context.Context, boil.ContextExecutor, *ActivityType) error

	activityTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	activityTypeType                 = reflect.TypeOf(&ActivityType{})
	activityTypeMapping              = queries.MakeStructMapping(activityTypeType)
	activityTypePrimaryKeyMapping, _ = queries.BindMapping(activityTypeType, activityTypeMapping, activityTypePrimaryKeyColumns)
	activityTypeInsertCacheMut       sync.RWMutex
	activityTypeInsertCache          = make(map[string]insertCache)
	activityTypeUpdateCacheMut       sync.RWMutex
	activityTypeUpdateCache          = make(map[string]updateCache)
	activityTypeUpsertCacheMut       sync.RWMutex
	activityTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var activityTypeBeforeInsertHooks []ActivityTypeHook
var activityTypeBeforeUpdateHooks []ActivityTypeHook
var activityTypeBeforeDeleteHooks []ActivityTypeHook
var activityTypeBeforeUpsertHooks []ActivityTypeHook

var activityTypeAfterInsertHooks []ActivityTypeHook
var activityTypeAfterSelectHooks []ActivityTypeHook
var activityTypeAfterUpdateHooks []ActivityTypeHook
var activityTypeAfterDeleteHooks []ActivityTypeHook
var activityTypeAfterUpsertHooks []ActivityTypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ActivityType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ActivityType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ActivityType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ActivityType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ActivityType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ActivityType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ActivityType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ActivityType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ActivityType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range activityTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddActivityTypeHook registers your hook function for all future operations.
func AddActivityTypeHook(hookPoint boil.HookPoint, activityTypeHook ActivityTypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		activityTypeBeforeInsertHooks = append(activityTypeBeforeInsertHooks, activityTypeHook)
	case boil.BeforeUpdateHook:
		activityTypeBeforeUpdateHooks = append(activityTypeBeforeUpdateHooks, activityTypeHook)
	case boil.BeforeDeleteHook:
		activityTypeBeforeDeleteHooks = append(activityTypeBeforeDeleteHooks, activityTypeHook)
	case boil.BeforeUpsertHook:
		activityTypeBeforeUpsertHooks = append(activityTypeBeforeUpsertHooks, activityTypeHook)
	case boil.AfterInsertHook:
		activityTypeAfterInsertHooks = append(activityTypeAfterInsertHooks, activityTypeHook)
	case boil.AfterSelectHook:
		activityTypeAfterSelectHooks = append(activityTypeAfterSelectHooks, activityTypeHook)
	case boil.AfterUpdateHook:
		activityTypeAfterUpdateHooks = append(activityTypeAfterUpdateHooks, activityTypeHook)
	case boil.AfterDeleteHook:
		activityTypeAfterDeleteHooks = append(activityTypeAfterDeleteHooks, activityTypeHook)
	case boil.AfterUpsertHook:
		activityTypeAfterUpsertHooks = append(activityTypeAfterUpsertHooks, activityTypeHook)
	}
}

// One returns a single activityType record from the query.
func (q activityTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ActivityType, error) {
	o := &ActivityType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for activity_type")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ActivityType records from the query.
func (q activityTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (ActivityTypeSlice, error) {
	var o []*ActivityType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ActivityType slice")
	}

	if len(activityTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ActivityType records in the query.
func (q activityTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count activity_type rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q activityTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if activity_type exists")
	}

	return count > 0, nil
}

// ActivityTypes retrieves all the records using an executor.
func ActivityTypes(mods ...qm.QueryMod) activityTypeQuery {
	mods = append(mods, qm.From("\"activity_type\""))
	return activityTypeQuery{NewQuery(mods...)}
}

// FindActivityType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindActivityType(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ActivityType, error) {
	activityTypeObj := &ActivityType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"activity_type\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, activityTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from activity_type")
	}

	return activityTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ActivityType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no activity_type provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(activityTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	activityTypeInsertCacheMut.RLock()
	cache, cached := activityTypeInsertCache[key]
	activityTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			activityTypeColumns,
			activityTypeColumnsWithDefault,
			activityTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(activityTypeType, activityTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(activityTypeType, activityTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"activity_type\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"activity_type\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into activity_type")
	}

	if !cached {
		activityTypeInsertCacheMut.Lock()
		activityTypeInsertCache[key] = cache
		activityTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ActivityType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ActivityType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	activityTypeUpdateCacheMut.RLock()
	cache, cached := activityTypeUpdateCache[key]
	activityTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			activityTypeColumns,
			activityTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update activity_type, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"activity_type\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, activityTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(activityTypeType, activityTypeMapping, append(wl, activityTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update activity_type row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for activity_type")
	}

	if !cached {
		activityTypeUpdateCacheMut.Lock()
		activityTypeUpdateCache[key] = cache
		activityTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q activityTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for activity_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for activity_type")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ActivityTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), activityTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"activity_type\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, activityTypePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in activityType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all activityType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ActivityType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no activity_type provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(activityTypeColumnsWithDefault, o)

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

	activityTypeUpsertCacheMut.RLock()
	cache, cached := activityTypeUpsertCache[key]
	activityTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			activityTypeColumns,
			activityTypeColumnsWithDefault,
			activityTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			activityTypeColumns,
			activityTypePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert activity_type, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(activityTypePrimaryKeyColumns))
			copy(conflict, activityTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"activity_type\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(activityTypeType, activityTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(activityTypeType, activityTypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert activity_type")
	}

	if !cached {
		activityTypeUpsertCacheMut.Lock()
		activityTypeUpsertCache[key] = cache
		activityTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ActivityType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ActivityType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ActivityType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), activityTypePrimaryKeyMapping)
	sql := "DELETE FROM \"activity_type\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from activity_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for activity_type")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q activityTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no activityTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from activity_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for activity_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ActivityTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ActivityType slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(activityTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), activityTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"activity_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, activityTypePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from activityType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for activity_type")
	}

	if len(activityTypeAfterDeleteHooks) != 0 {
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
func (o *ActivityType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindActivityType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ActivityTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ActivityTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), activityTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"activity_type\".* FROM \"activity_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, activityTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ActivityTypeSlice")
	}

	*o = slice

	return nil
}

// ActivityTypeExists checks if the ActivityType row exists.
func ActivityTypeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"activity_type\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if activity_type exists")
	}

	return exists, nil
}
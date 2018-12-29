// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testElections(t *testing.T) {
	t.Parallel()

	query := Elections()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testElectionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testElectionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Elections().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testElectionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ElectionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testElectionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ElectionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Election exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ElectionExists to return true, but got false.")
	}
}

func testElectionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	electionFound, err := FindElection(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if electionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testElectionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Elections().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testElectionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Elections().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testElectionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	electionOne := &Election{}
	electionTwo := &Election{}
	if err = randomize.Struct(seed, electionOne, electionDBTypes, false, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}
	if err = randomize.Struct(seed, electionTwo, electionDBTypes, false, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = electionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = electionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Elections().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testElectionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	electionOne := &Election{}
	electionTwo := &Election{}
	if err = randomize.Struct(seed, electionOne, electionDBTypes, false, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}
	if err = randomize.Struct(seed, electionTwo, electionDBTypes, false, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = electionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = electionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func electionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func electionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func electionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func electionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func electionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func electionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func electionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func electionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func electionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Election) error {
	*o = Election{}
	return nil
}

func testElectionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Election{}
	o := &Election{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, electionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Election object: %s", err)
	}

	AddElectionHook(boil.BeforeInsertHook, electionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	electionBeforeInsertHooks = []ElectionHook{}

	AddElectionHook(boil.AfterInsertHook, electionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	electionAfterInsertHooks = []ElectionHook{}

	AddElectionHook(boil.AfterSelectHook, electionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	electionAfterSelectHooks = []ElectionHook{}

	AddElectionHook(boil.BeforeUpdateHook, electionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	electionBeforeUpdateHooks = []ElectionHook{}

	AddElectionHook(boil.AfterUpdateHook, electionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	electionAfterUpdateHooks = []ElectionHook{}

	AddElectionHook(boil.BeforeDeleteHook, electionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	electionBeforeDeleteHooks = []ElectionHook{}

	AddElectionHook(boil.AfterDeleteHook, electionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	electionAfterDeleteHooks = []ElectionHook{}

	AddElectionHook(boil.BeforeUpsertHook, electionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	electionBeforeUpsertHooks = []ElectionHook{}

	AddElectionHook(boil.AfterUpsertHook, electionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	electionAfterUpsertHooks = []ElectionHook{}
}

func testElectionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testElectionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(electionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testElectionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testElectionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ElectionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testElectionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Elections().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	electionDBTypes = map[string]string{`CreatedAt`: `timestamp without time zone`, `ID`: `integer`, `UpdatedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testElectionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(electionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(electionColumns) == len(electionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, electionDBTypes, true, electionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testElectionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(electionColumns) == len(electionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Election{}
	if err = randomize.Struct(seed, o, electionDBTypes, true, electionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, electionDBTypes, true, electionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(electionColumns, electionPrimaryKeyColumns) {
		fields = electionColumns
	} else {
		fields = strmangle.SetComplement(
			electionColumns,
			electionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ElectionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testElectionsUpsert(t *testing.T) {
	t.Parallel()

	if len(electionColumns) == len(electionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Election{}
	if err = randomize.Struct(seed, &o, electionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Election: %s", err)
	}

	count, err := Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, electionDBTypes, false, electionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Election struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Election: %s", err)
	}

	count, err = Elections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
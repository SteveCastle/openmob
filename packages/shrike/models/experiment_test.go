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

func testExperiments(t *testing.T) {
	t.Parallel()

	query := Experiments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testExperimentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
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

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExperimentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Experiments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExperimentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ExperimentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExperimentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ExperimentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Experiment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ExperimentExists to return true, but got false.")
	}
}

func testExperimentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	experimentFound, err := FindExperiment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if experimentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testExperimentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Experiments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testExperimentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Experiments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testExperimentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	experimentOne := &Experiment{}
	experimentTwo := &Experiment{}
	if err = randomize.Struct(seed, experimentOne, experimentDBTypes, false, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}
	if err = randomize.Struct(seed, experimentTwo, experimentDBTypes, false, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = experimentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = experimentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Experiments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testExperimentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	experimentOne := &Experiment{}
	experimentTwo := &Experiment{}
	if err = randomize.Struct(seed, experimentOne, experimentDBTypes, false, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}
	if err = randomize.Struct(seed, experimentTwo, experimentDBTypes, false, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = experimentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = experimentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func experimentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func experimentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func experimentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func experimentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func experimentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func experimentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func experimentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func experimentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func experimentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Experiment) error {
	*o = Experiment{}
	return nil
}

func testExperimentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Experiment{}
	o := &Experiment{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, experimentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Experiment object: %s", err)
	}

	AddExperimentHook(boil.BeforeInsertHook, experimentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	experimentBeforeInsertHooks = []ExperimentHook{}

	AddExperimentHook(boil.AfterInsertHook, experimentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	experimentAfterInsertHooks = []ExperimentHook{}

	AddExperimentHook(boil.AfterSelectHook, experimentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	experimentAfterSelectHooks = []ExperimentHook{}

	AddExperimentHook(boil.BeforeUpdateHook, experimentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	experimentBeforeUpdateHooks = []ExperimentHook{}

	AddExperimentHook(boil.AfterUpdateHook, experimentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	experimentAfterUpdateHooks = []ExperimentHook{}

	AddExperimentHook(boil.BeforeDeleteHook, experimentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	experimentBeforeDeleteHooks = []ExperimentHook{}

	AddExperimentHook(boil.AfterDeleteHook, experimentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	experimentAfterDeleteHooks = []ExperimentHook{}

	AddExperimentHook(boil.BeforeUpsertHook, experimentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	experimentBeforeUpsertHooks = []ExperimentHook{}

	AddExperimentHook(boil.AfterUpsertHook, experimentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	experimentAfterUpsertHooks = []ExperimentHook{}
}

func testExperimentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExperimentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(experimentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExperimentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
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

func testExperimentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ExperimentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testExperimentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Experiments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	experimentDBTypes = map[string]string{`CreatedAt`: `timestamp without time zone`, `ID`: `integer`, `UpdatedAt`: `timestamp without time zone`}
	_                 = bytes.MinRead
)

func testExperimentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(experimentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(experimentColumns) == len(experimentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testExperimentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(experimentColumns) == len(experimentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Experiment{}
	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, experimentDBTypes, true, experimentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(experimentColumns, experimentPrimaryKeyColumns) {
		fields = experimentColumns
	} else {
		fields = strmangle.SetComplement(
			experimentColumns,
			experimentPrimaryKeyColumns,
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

	slice := ExperimentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testExperimentsUpsert(t *testing.T) {
	t.Parallel()

	if len(experimentColumns) == len(experimentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Experiment{}
	if err = randomize.Struct(seed, &o, experimentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Experiment: %s", err)
	}

	count, err := Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, experimentDBTypes, false, experimentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Experiment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Experiment: %s", err)
	}

	count, err = Experiments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
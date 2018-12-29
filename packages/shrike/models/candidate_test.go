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

func testCandidates(t *testing.T) {
	t.Parallel()

	query := Candidates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCandidatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
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

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCandidatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Candidates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCandidatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CandidateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCandidatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CandidateExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Candidate exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CandidateExists to return true, but got false.")
	}
}

func testCandidatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	candidateFound, err := FindCandidate(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if candidateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCandidatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Candidates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCandidatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Candidates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCandidatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	candidateOne := &Candidate{}
	candidateTwo := &Candidate{}
	if err = randomize.Struct(seed, candidateOne, candidateDBTypes, false, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}
	if err = randomize.Struct(seed, candidateTwo, candidateDBTypes, false, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = candidateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = candidateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Candidates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCandidatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	candidateOne := &Candidate{}
	candidateTwo := &Candidate{}
	if err = randomize.Struct(seed, candidateOne, candidateDBTypes, false, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}
	if err = randomize.Struct(seed, candidateTwo, candidateDBTypes, false, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = candidateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = candidateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func candidateBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func candidateAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func candidateAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func candidateBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func candidateAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func candidateBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func candidateAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func candidateBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func candidateAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Candidate) error {
	*o = Candidate{}
	return nil
}

func testCandidatesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Candidate{}
	o := &Candidate{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, candidateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Candidate object: %s", err)
	}

	AddCandidateHook(boil.BeforeInsertHook, candidateBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	candidateBeforeInsertHooks = []CandidateHook{}

	AddCandidateHook(boil.AfterInsertHook, candidateAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	candidateAfterInsertHooks = []CandidateHook{}

	AddCandidateHook(boil.AfterSelectHook, candidateAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	candidateAfterSelectHooks = []CandidateHook{}

	AddCandidateHook(boil.BeforeUpdateHook, candidateBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	candidateBeforeUpdateHooks = []CandidateHook{}

	AddCandidateHook(boil.AfterUpdateHook, candidateAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	candidateAfterUpdateHooks = []CandidateHook{}

	AddCandidateHook(boil.BeforeDeleteHook, candidateBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	candidateBeforeDeleteHooks = []CandidateHook{}

	AddCandidateHook(boil.AfterDeleteHook, candidateAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	candidateAfterDeleteHooks = []CandidateHook{}

	AddCandidateHook(boil.BeforeUpsertHook, candidateBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	candidateBeforeUpsertHooks = []CandidateHook{}

	AddCandidateHook(boil.AfterUpsertHook, candidateAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	candidateAfterUpsertHooks = []CandidateHook{}
}

func testCandidatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCandidatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(candidateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCandidatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
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

func testCandidatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CandidateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCandidatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Candidates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	candidateDBTypes = map[string]string{`CreatedAt`: `timestamp without time zone`, `ID`: `integer`, `UpdatedAt`: `timestamp without time zone`}
	_                = bytes.MinRead
)

func testCandidatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(candidatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(candidateColumns) == len(candidatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCandidatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(candidateColumns) == len(candidatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Candidate{}
	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, candidateDBTypes, true, candidatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(candidateColumns, candidatePrimaryKeyColumns) {
		fields = candidateColumns
	} else {
		fields = strmangle.SetComplement(
			candidateColumns,
			candidatePrimaryKeyColumns,
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

	slice := CandidateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCandidatesUpsert(t *testing.T) {
	t.Parallel()

	if len(candidateColumns) == len(candidatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Candidate{}
	if err = randomize.Struct(seed, &o, candidateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Candidate: %s", err)
	}

	count, err := Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, candidateDBTypes, false, candidatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Candidate struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Candidate: %s", err)
	}

	count, err = Candidates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

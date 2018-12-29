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

func testNotes(t *testing.T) {
	t.Parallel()

	query := Notes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNotesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
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

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Notes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NoteSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NoteExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Note exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NoteExists to return true, but got false.")
	}
}

func testNotesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	noteFound, err := FindNote(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if noteFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNotesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Notes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNotesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Notes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNotesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	noteOne := &Note{}
	noteTwo := &Note{}
	if err = randomize.Struct(seed, noteOne, noteDBTypes, false, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}
	if err = randomize.Struct(seed, noteTwo, noteDBTypes, false, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = noteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = noteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Notes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNotesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	noteOne := &Note{}
	noteTwo := &Note{}
	if err = randomize.Struct(seed, noteOne, noteDBTypes, false, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}
	if err = randomize.Struct(seed, noteTwo, noteDBTypes, false, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = noteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = noteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func noteBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func noteAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func noteAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func noteBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func noteAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func noteBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func noteAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func noteBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func noteAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Note) error {
	*o = Note{}
	return nil
}

func testNotesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Note{}
	o := &Note{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, noteDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Note object: %s", err)
	}

	AddNoteHook(boil.BeforeInsertHook, noteBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	noteBeforeInsertHooks = []NoteHook{}

	AddNoteHook(boil.AfterInsertHook, noteAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	noteAfterInsertHooks = []NoteHook{}

	AddNoteHook(boil.AfterSelectHook, noteAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	noteAfterSelectHooks = []NoteHook{}

	AddNoteHook(boil.BeforeUpdateHook, noteBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	noteBeforeUpdateHooks = []NoteHook{}

	AddNoteHook(boil.AfterUpdateHook, noteAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	noteAfterUpdateHooks = []NoteHook{}

	AddNoteHook(boil.BeforeDeleteHook, noteBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	noteBeforeDeleteHooks = []NoteHook{}

	AddNoteHook(boil.AfterDeleteHook, noteAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	noteAfterDeleteHooks = []NoteHook{}

	AddNoteHook(boil.BeforeUpsertHook, noteBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	noteBeforeUpsertHooks = []NoteHook{}

	AddNoteHook(boil.AfterUpsertHook, noteAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	noteAfterUpsertHooks = []NoteHook{}
}

func testNotesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(noteColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
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

func testNotesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NoteSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNotesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Notes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	noteDBTypes = map[string]string{`CreatedAt`: `timestamp without time zone`, `ID`: `integer`, `UpdatedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testNotesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(notePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(noteColumns) == len(notePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, noteDBTypes, true, notePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNotesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(noteColumns) == len(notePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Note{}
	if err = randomize.Struct(seed, o, noteDBTypes, true, noteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, noteDBTypes, true, notePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(noteColumns, notePrimaryKeyColumns) {
		fields = noteColumns
	} else {
		fields = strmangle.SetComplement(
			noteColumns,
			notePrimaryKeyColumns,
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

	slice := NoteSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNotesUpsert(t *testing.T) {
	t.Parallel()

	if len(noteColumns) == len(notePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Note{}
	if err = randomize.Struct(seed, &o, noteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Note: %s", err)
	}

	count, err := Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, noteDBTypes, false, notePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Note struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Note: %s", err)
	}

	count, err = Notes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
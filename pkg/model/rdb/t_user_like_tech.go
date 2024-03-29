// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TUserLikeTech is an object representing the database table.
type TUserLikeTech struct { // id
	ID int `boil:"id" json:"id" toml:"id" yaml:"id"`
	// user id
	UserID int `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	// tech id
	TechID int `boil:"tech_id" json:"tech_id" toml:"tech_id" yaml:"tech_id"`
	// delete flag
	IsDeleted null.String `boil:"is_deleted" json:"is_deleted,omitempty" toml:"is_deleted" yaml:"is_deleted,omitempty"`
	// created date
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	// updated date
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *tUserLikeTechR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tUserLikeTechL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TUserLikeTechColumns = struct {
	ID        string
	UserID    string
	TechID    string
	IsDeleted string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	TechID:    "tech_id",
	IsDeleted: "is_deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var TUserLikeTechTableColumns = struct {
	ID        string
	UserID    string
	TechID    string
	IsDeleted string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "t_user_like_tech.id",
	UserID:    "t_user_like_tech.user_id",
	TechID:    "t_user_like_tech.tech_id",
	IsDeleted: "t_user_like_tech.is_deleted",
	CreatedAt: "t_user_like_tech.created_at",
	UpdatedAt: "t_user_like_tech.updated_at",
}

// Generated where

var TUserLikeTechWhere = struct {
	ID        whereHelperint
	UserID    whereHelperint
	TechID    whereHelperint
	IsDeleted whereHelpernull_String
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "`t_user_like_tech`.`id`"},
	UserID:    whereHelperint{field: "`t_user_like_tech`.`user_id`"},
	TechID:    whereHelperint{field: "`t_user_like_tech`.`tech_id`"},
	IsDeleted: whereHelpernull_String{field: "`t_user_like_tech`.`is_deleted`"},
	CreatedAt: whereHelpernull_Time{field: "`t_user_like_tech`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`t_user_like_tech`.`updated_at`"},
}

// TUserLikeTechRels is where relationship names are stored.
var TUserLikeTechRels = struct {
}{}

// tUserLikeTechR is where relationships are stored.
type tUserLikeTechR struct {
}

// NewStruct creates a new relationship struct
func (*tUserLikeTechR) NewStruct() *tUserLikeTechR {
	return &tUserLikeTechR{}
}

// tUserLikeTechL is where Load methods for each relationship are stored.
type tUserLikeTechL struct{}

var (
	tUserLikeTechAllColumns            = []string{"id", "user_id", "tech_id", "is_deleted", "created_at", "updated_at"}
	tUserLikeTechColumnsWithoutDefault = []string{"user_id", "tech_id"}
	tUserLikeTechColumnsWithDefault    = []string{"id", "is_deleted", "created_at", "updated_at"}
	tUserLikeTechPrimaryKeyColumns     = []string{"id"}
	tUserLikeTechGeneratedColumns      = []string{}
)

type (
	// TUserLikeTechSlice is an alias for a slice of pointers to TUserLikeTech.
	// This should almost always be used instead of []TUserLikeTech.
	TUserLikeTechSlice []*TUserLikeTech

	tUserLikeTechQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tUserLikeTechType                 = reflect.TypeOf(&TUserLikeTech{})
	tUserLikeTechMapping              = queries.MakeStructMapping(tUserLikeTechType)
	tUserLikeTechPrimaryKeyMapping, _ = queries.BindMapping(tUserLikeTechType, tUserLikeTechMapping, tUserLikeTechPrimaryKeyColumns)
	tUserLikeTechInsertCacheMut       sync.RWMutex
	tUserLikeTechInsertCache          = make(map[string]insertCache)
	tUserLikeTechUpdateCacheMut       sync.RWMutex
	tUserLikeTechUpdateCache          = make(map[string]updateCache)
	tUserLikeTechUpsertCacheMut       sync.RWMutex
	tUserLikeTechUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single tUserLikeTech record from the query.
func (q tUserLikeTechQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TUserLikeTech, error) {
	o := &TUserLikeTech{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for t_user_like_tech")
	}

	return o, nil
}

// All returns all TUserLikeTech records from the query.
func (q tUserLikeTechQuery) All(ctx context.Context, exec boil.ContextExecutor) (TUserLikeTechSlice, error) {
	var o []*TUserLikeTech

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TUserLikeTech slice")
	}

	return o, nil
}

// Count returns the count of all TUserLikeTech records in the query.
func (q tUserLikeTechQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count t_user_like_tech rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tUserLikeTechQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if t_user_like_tech exists")
	}

	return count > 0, nil
}

// TUserLikeTeches retrieves all the records using an executor.
func TUserLikeTeches(mods ...qm.QueryMod) tUserLikeTechQuery {
	mods = append(mods, qm.From("`t_user_like_tech`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`t_user_like_tech`.*"})
	}

	return tUserLikeTechQuery{q}
}

// FindTUserLikeTech retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTUserLikeTech(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TUserLikeTech, error) {
	tUserLikeTechObj := &TUserLikeTech{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `t_user_like_tech` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tUserLikeTechObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from t_user_like_tech")
	}

	return tUserLikeTechObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TUserLikeTech) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no t_user_like_tech provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(tUserLikeTechColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tUserLikeTechInsertCacheMut.RLock()
	cache, cached := tUserLikeTechInsertCache[key]
	tUserLikeTechInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tUserLikeTechAllColumns,
			tUserLikeTechColumnsWithDefault,
			tUserLikeTechColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tUserLikeTechType, tUserLikeTechMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tUserLikeTechType, tUserLikeTechMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `t_user_like_tech` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `t_user_like_tech` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `t_user_like_tech` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tUserLikeTechPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into t_user_like_tech")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tUserLikeTechMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for t_user_like_tech")
	}

CacheNoHooks:
	if !cached {
		tUserLikeTechInsertCacheMut.Lock()
		tUserLikeTechInsertCache[key] = cache
		tUserLikeTechInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the TUserLikeTech.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TUserLikeTech) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	tUserLikeTechUpdateCacheMut.RLock()
	cache, cached := tUserLikeTechUpdateCache[key]
	tUserLikeTechUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tUserLikeTechAllColumns,
			tUserLikeTechPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update t_user_like_tech, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `t_user_like_tech` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tUserLikeTechPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tUserLikeTechType, tUserLikeTechMapping, append(wl, tUserLikeTechPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update t_user_like_tech row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for t_user_like_tech")
	}

	if !cached {
		tUserLikeTechUpdateCacheMut.Lock()
		tUserLikeTechUpdateCache[key] = cache
		tUserLikeTechUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q tUserLikeTechQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for t_user_like_tech")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for t_user_like_tech")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TUserLikeTechSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tUserLikeTechPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `t_user_like_tech` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tUserLikeTechPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tUserLikeTech slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tUserLikeTech")
	}
	return rowsAff, nil
}

var mySQLTUserLikeTechUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TUserLikeTech) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no t_user_like_tech provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(tUserLikeTechColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTUserLikeTechUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tUserLikeTechUpsertCacheMut.RLock()
	cache, cached := tUserLikeTechUpsertCache[key]
	tUserLikeTechUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tUserLikeTechAllColumns,
			tUserLikeTechColumnsWithDefault,
			tUserLikeTechColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tUserLikeTechAllColumns,
			tUserLikeTechPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert t_user_like_tech, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`t_user_like_tech`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `t_user_like_tech` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tUserLikeTechType, tUserLikeTechMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tUserLikeTechType, tUserLikeTechMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for t_user_like_tech")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tUserLikeTechMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tUserLikeTechType, tUserLikeTechMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for t_user_like_tech")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for t_user_like_tech")
	}

CacheNoHooks:
	if !cached {
		tUserLikeTechUpsertCacheMut.Lock()
		tUserLikeTechUpsertCache[key] = cache
		tUserLikeTechUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single TUserLikeTech record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TUserLikeTech) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TUserLikeTech provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tUserLikeTechPrimaryKeyMapping)
	sql := "DELETE FROM `t_user_like_tech` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from t_user_like_tech")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for t_user_like_tech")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tUserLikeTechQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tUserLikeTechQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from t_user_like_tech")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for t_user_like_tech")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TUserLikeTechSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tUserLikeTechPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `t_user_like_tech` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tUserLikeTechPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tUserLikeTech slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for t_user_like_tech")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TUserLikeTech) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTUserLikeTech(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TUserLikeTechSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TUserLikeTechSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tUserLikeTechPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `t_user_like_tech`.* FROM `t_user_like_tech` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tUserLikeTechPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TUserLikeTechSlice")
	}

	*o = slice

	return nil
}

// TUserLikeTechExists checks if the TUserLikeTech row exists.
func TUserLikeTechExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `t_user_like_tech` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if t_user_like_tech exists")
	}

	return exists, nil
}

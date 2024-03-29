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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// TUserWorkHistory is an object representing the database table.
type TUserWorkHistory struct { // id
	ID int `boil:"id" json:"id" toml:"id" yaml:"id"`
	// user id
	UserID int `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	// company id
	CompanyID int `boil:"company_id" json:"company_id" toml:"company_id" yaml:"company_id"`
	// title
	Title string `boil:"title" json:"title" toml:"title" yaml:"title"`
	// description
	Description types.JSON `boil:"description" json:"description" toml:"description" yaml:"description"`
	// tech ids
	TechIds types.JSON `boil:"tech_ids" json:"tech_ids" toml:"tech_ids" yaml:"tech_ids"`
	// started date
	StartedAt null.Time `boil:"started_at" json:"started_at,omitempty" toml:"started_at" yaml:"started_at,omitempty"`
	// ended date
	EndedAt null.Time `boil:"ended_at" json:"ended_at,omitempty" toml:"ended_at" yaml:"ended_at,omitempty"`
	// delete flag
	IsDeleted null.String `boil:"is_deleted" json:"is_deleted,omitempty" toml:"is_deleted" yaml:"is_deleted,omitempty"`
	// created date
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	// updated date
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *tUserWorkHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tUserWorkHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TUserWorkHistoryColumns = struct {
	ID          string
	UserID      string
	CompanyID   string
	Title       string
	Description string
	TechIds     string
	StartedAt   string
	EndedAt     string
	IsDeleted   string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	UserID:      "user_id",
	CompanyID:   "company_id",
	Title:       "title",
	Description: "description",
	TechIds:     "tech_ids",
	StartedAt:   "started_at",
	EndedAt:     "ended_at",
	IsDeleted:   "is_deleted",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var TUserWorkHistoryTableColumns = struct {
	ID          string
	UserID      string
	CompanyID   string
	Title       string
	Description string
	TechIds     string
	StartedAt   string
	EndedAt     string
	IsDeleted   string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "t_user_work_history.id",
	UserID:      "t_user_work_history.user_id",
	CompanyID:   "t_user_work_history.company_id",
	Title:       "t_user_work_history.title",
	Description: "t_user_work_history.description",
	TechIds:     "t_user_work_history.tech_ids",
	StartedAt:   "t_user_work_history.started_at",
	EndedAt:     "t_user_work_history.ended_at",
	IsDeleted:   "t_user_work_history.is_deleted",
	CreatedAt:   "t_user_work_history.created_at",
	UpdatedAt:   "t_user_work_history.updated_at",
}

// Generated where

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var TUserWorkHistoryWhere = struct {
	ID          whereHelperint
	UserID      whereHelperint
	CompanyID   whereHelperint
	Title       whereHelperstring
	Description whereHelpertypes_JSON
	TechIds     whereHelpertypes_JSON
	StartedAt   whereHelpernull_Time
	EndedAt     whereHelpernull_Time
	IsDeleted   whereHelpernull_String
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
}{
	ID:          whereHelperint{field: "`t_user_work_history`.`id`"},
	UserID:      whereHelperint{field: "`t_user_work_history`.`user_id`"},
	CompanyID:   whereHelperint{field: "`t_user_work_history`.`company_id`"},
	Title:       whereHelperstring{field: "`t_user_work_history`.`title`"},
	Description: whereHelpertypes_JSON{field: "`t_user_work_history`.`description`"},
	TechIds:     whereHelpertypes_JSON{field: "`t_user_work_history`.`tech_ids`"},
	StartedAt:   whereHelpernull_Time{field: "`t_user_work_history`.`started_at`"},
	EndedAt:     whereHelpernull_Time{field: "`t_user_work_history`.`ended_at`"},
	IsDeleted:   whereHelpernull_String{field: "`t_user_work_history`.`is_deleted`"},
	CreatedAt:   whereHelpernull_Time{field: "`t_user_work_history`.`created_at`"},
	UpdatedAt:   whereHelpernull_Time{field: "`t_user_work_history`.`updated_at`"},
}

// TUserWorkHistoryRels is where relationship names are stored.
var TUserWorkHistoryRels = struct {
}{}

// tUserWorkHistoryR is where relationships are stored.
type tUserWorkHistoryR struct {
}

// NewStruct creates a new relationship struct
func (*tUserWorkHistoryR) NewStruct() *tUserWorkHistoryR {
	return &tUserWorkHistoryR{}
}

// tUserWorkHistoryL is where Load methods for each relationship are stored.
type tUserWorkHistoryL struct{}

var (
	tUserWorkHistoryAllColumns            = []string{"id", "user_id", "company_id", "title", "description", "tech_ids", "started_at", "ended_at", "is_deleted", "created_at", "updated_at"}
	tUserWorkHistoryColumnsWithoutDefault = []string{"user_id", "company_id", "title", "description", "tech_ids", "started_at", "ended_at"}
	tUserWorkHistoryColumnsWithDefault    = []string{"id", "is_deleted", "created_at", "updated_at"}
	tUserWorkHistoryPrimaryKeyColumns     = []string{"id"}
	tUserWorkHistoryGeneratedColumns      = []string{}
)

type (
	// TUserWorkHistorySlice is an alias for a slice of pointers to TUserWorkHistory.
	// This should almost always be used instead of []TUserWorkHistory.
	TUserWorkHistorySlice []*TUserWorkHistory

	tUserWorkHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tUserWorkHistoryType                 = reflect.TypeOf(&TUserWorkHistory{})
	tUserWorkHistoryMapping              = queries.MakeStructMapping(tUserWorkHistoryType)
	tUserWorkHistoryPrimaryKeyMapping, _ = queries.BindMapping(tUserWorkHistoryType, tUserWorkHistoryMapping, tUserWorkHistoryPrimaryKeyColumns)
	tUserWorkHistoryInsertCacheMut       sync.RWMutex
	tUserWorkHistoryInsertCache          = make(map[string]insertCache)
	tUserWorkHistoryUpdateCacheMut       sync.RWMutex
	tUserWorkHistoryUpdateCache          = make(map[string]updateCache)
	tUserWorkHistoryUpsertCacheMut       sync.RWMutex
	tUserWorkHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single tUserWorkHistory record from the query.
func (q tUserWorkHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TUserWorkHistory, error) {
	o := &TUserWorkHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for t_user_work_history")
	}

	return o, nil
}

// All returns all TUserWorkHistory records from the query.
func (q tUserWorkHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (TUserWorkHistorySlice, error) {
	var o []*TUserWorkHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TUserWorkHistory slice")
	}

	return o, nil
}

// Count returns the count of all TUserWorkHistory records in the query.
func (q tUserWorkHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count t_user_work_history rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tUserWorkHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if t_user_work_history exists")
	}

	return count > 0, nil
}

// TUserWorkHistories retrieves all the records using an executor.
func TUserWorkHistories(mods ...qm.QueryMod) tUserWorkHistoryQuery {
	mods = append(mods, qm.From("`t_user_work_history`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`t_user_work_history`.*"})
	}

	return tUserWorkHistoryQuery{q}
}

// FindTUserWorkHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTUserWorkHistory(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TUserWorkHistory, error) {
	tUserWorkHistoryObj := &TUserWorkHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `t_user_work_history` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tUserWorkHistoryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from t_user_work_history")
	}

	return tUserWorkHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TUserWorkHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no t_user_work_history provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(tUserWorkHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tUserWorkHistoryInsertCacheMut.RLock()
	cache, cached := tUserWorkHistoryInsertCache[key]
	tUserWorkHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tUserWorkHistoryAllColumns,
			tUserWorkHistoryColumnsWithDefault,
			tUserWorkHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tUserWorkHistoryType, tUserWorkHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tUserWorkHistoryType, tUserWorkHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `t_user_work_history` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `t_user_work_history` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `t_user_work_history` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tUserWorkHistoryPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into t_user_work_history")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tUserWorkHistoryMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for t_user_work_history")
	}

CacheNoHooks:
	if !cached {
		tUserWorkHistoryInsertCacheMut.Lock()
		tUserWorkHistoryInsertCache[key] = cache
		tUserWorkHistoryInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the TUserWorkHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TUserWorkHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	tUserWorkHistoryUpdateCacheMut.RLock()
	cache, cached := tUserWorkHistoryUpdateCache[key]
	tUserWorkHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tUserWorkHistoryAllColumns,
			tUserWorkHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update t_user_work_history, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `t_user_work_history` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tUserWorkHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tUserWorkHistoryType, tUserWorkHistoryMapping, append(wl, tUserWorkHistoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update t_user_work_history row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for t_user_work_history")
	}

	if !cached {
		tUserWorkHistoryUpdateCacheMut.Lock()
		tUserWorkHistoryUpdateCache[key] = cache
		tUserWorkHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q tUserWorkHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for t_user_work_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for t_user_work_history")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TUserWorkHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tUserWorkHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `t_user_work_history` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tUserWorkHistoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tUserWorkHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tUserWorkHistory")
	}
	return rowsAff, nil
}

var mySQLTUserWorkHistoryUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TUserWorkHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no t_user_work_history provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(tUserWorkHistoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTUserWorkHistoryUniqueColumns, o)

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

	tUserWorkHistoryUpsertCacheMut.RLock()
	cache, cached := tUserWorkHistoryUpsertCache[key]
	tUserWorkHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tUserWorkHistoryAllColumns,
			tUserWorkHistoryColumnsWithDefault,
			tUserWorkHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tUserWorkHistoryAllColumns,
			tUserWorkHistoryPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert t_user_work_history, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`t_user_work_history`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `t_user_work_history` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tUserWorkHistoryType, tUserWorkHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tUserWorkHistoryType, tUserWorkHistoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for t_user_work_history")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tUserWorkHistoryMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tUserWorkHistoryType, tUserWorkHistoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for t_user_work_history")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for t_user_work_history")
	}

CacheNoHooks:
	if !cached {
		tUserWorkHistoryUpsertCacheMut.Lock()
		tUserWorkHistoryUpsertCache[key] = cache
		tUserWorkHistoryUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single TUserWorkHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TUserWorkHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TUserWorkHistory provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tUserWorkHistoryPrimaryKeyMapping)
	sql := "DELETE FROM `t_user_work_history` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from t_user_work_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for t_user_work_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tUserWorkHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tUserWorkHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from t_user_work_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for t_user_work_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TUserWorkHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tUserWorkHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `t_user_work_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tUserWorkHistoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tUserWorkHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for t_user_work_history")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TUserWorkHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTUserWorkHistory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TUserWorkHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TUserWorkHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tUserWorkHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `t_user_work_history`.* FROM `t_user_work_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tUserWorkHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TUserWorkHistorySlice")
	}

	*o = slice

	return nil
}

// TUserWorkHistoryExists checks if the TUserWorkHistory row exists.
func TUserWorkHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `t_user_work_history` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if t_user_work_history exists")
	}

	return exists, nil
}

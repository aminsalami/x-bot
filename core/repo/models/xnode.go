// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Xnode is an object representing the database table.
type Xnode struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Address   string    `boil:"address" json:"address" toml:"address" yaml:"address"`
	PanelType string    `boil:"panel_type" json:"panel_type" toml:"panel_type" yaml:"panel_type"`
	Active    null.Bool `boil:"active" json:"active,omitempty" toml:"active" yaml:"active,omitempty"`

	R *xnodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L xnodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var XnodeColumns = struct {
	ID        string
	Address   string
	PanelType string
	Active    string
}{
	ID:        "id",
	Address:   "address",
	PanelType: "panel_type",
	Active:    "active",
}

var XnodeTableColumns = struct {
	ID        string
	Address   string
	PanelType string
	Active    string
}{
	ID:        "xnode.id",
	Address:   "xnode.address",
	PanelType: "xnode.panel_type",
	Active:    "xnode.active",
}

// Generated where

var XnodeWhere = struct {
	ID        whereHelperint64
	Address   whereHelperstring
	PanelType whereHelperstring
	Active    whereHelpernull_Bool
}{
	ID:        whereHelperint64{field: "\"xnode\".\"id\""},
	Address:   whereHelperstring{field: "\"xnode\".\"address\""},
	PanelType: whereHelperstring{field: "\"xnode\".\"panel_type\""},
	Active:    whereHelpernull_Bool{field: "\"xnode\".\"active\""},
}

// XnodeRels is where relationship names are stored.
var XnodeRels = struct {
}{}

// xnodeR is where relationships are stored.
type xnodeR struct {
}

// NewStruct creates a new relationship struct
func (*xnodeR) NewStruct() *xnodeR {
	return &xnodeR{}
}

// xnodeL is where Load methods for each relationship are stored.
type xnodeL struct{}

var (
	xnodeAllColumns            = []string{"id", "address", "panel_type", "active"}
	xnodeColumnsWithoutDefault = []string{"address", "panel_type"}
	xnodeColumnsWithDefault    = []string{"id", "active"}
	xnodePrimaryKeyColumns     = []string{"id"}
	xnodeGeneratedColumns      = []string{"id"}
)

type (
	// XnodeSlice is an alias for a slice of pointers to Xnode.
	// This should almost always be used instead of []Xnode.
	XnodeSlice []*Xnode
	// XnodeHook is the signature for custom Xnode hook methods
	XnodeHook func(context.Context, boil.ContextExecutor, *Xnode) error

	xnodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	xnodeType                 = reflect.TypeOf(&Xnode{})
	xnodeMapping              = queries.MakeStructMapping(xnodeType)
	xnodePrimaryKeyMapping, _ = queries.BindMapping(xnodeType, xnodeMapping, xnodePrimaryKeyColumns)
	xnodeInsertCacheMut       sync.RWMutex
	xnodeInsertCache          = make(map[string]insertCache)
	xnodeUpdateCacheMut       sync.RWMutex
	xnodeUpdateCache          = make(map[string]updateCache)
	xnodeUpsertCacheMut       sync.RWMutex
	xnodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var xnodeAfterSelectHooks []XnodeHook

var xnodeBeforeInsertHooks []XnodeHook
var xnodeAfterInsertHooks []XnodeHook

var xnodeBeforeUpdateHooks []XnodeHook
var xnodeAfterUpdateHooks []XnodeHook

var xnodeBeforeDeleteHooks []XnodeHook
var xnodeAfterDeleteHooks []XnodeHook

var xnodeBeforeUpsertHooks []XnodeHook
var xnodeAfterUpsertHooks []XnodeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Xnode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Xnode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Xnode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Xnode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Xnode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Xnode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Xnode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Xnode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Xnode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range xnodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddXnodeHook registers your hook function for all future operations.
func AddXnodeHook(hookPoint boil.HookPoint, xnodeHook XnodeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		xnodeAfterSelectHooks = append(xnodeAfterSelectHooks, xnodeHook)
	case boil.BeforeInsertHook:
		xnodeBeforeInsertHooks = append(xnodeBeforeInsertHooks, xnodeHook)
	case boil.AfterInsertHook:
		xnodeAfterInsertHooks = append(xnodeAfterInsertHooks, xnodeHook)
	case boil.BeforeUpdateHook:
		xnodeBeforeUpdateHooks = append(xnodeBeforeUpdateHooks, xnodeHook)
	case boil.AfterUpdateHook:
		xnodeAfterUpdateHooks = append(xnodeAfterUpdateHooks, xnodeHook)
	case boil.BeforeDeleteHook:
		xnodeBeforeDeleteHooks = append(xnodeBeforeDeleteHooks, xnodeHook)
	case boil.AfterDeleteHook:
		xnodeAfterDeleteHooks = append(xnodeAfterDeleteHooks, xnodeHook)
	case boil.BeforeUpsertHook:
		xnodeBeforeUpsertHooks = append(xnodeBeforeUpsertHooks, xnodeHook)
	case boil.AfterUpsertHook:
		xnodeAfterUpsertHooks = append(xnodeAfterUpsertHooks, xnodeHook)
	}
}

// One returns a single xnode record from the query.
func (q xnodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Xnode, error) {
	o := &Xnode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for xnode")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Xnode records from the query.
func (q xnodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (XnodeSlice, error) {
	var o []*Xnode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Xnode slice")
	}

	if len(xnodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Xnode records in the query.
func (q xnodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count xnode rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q xnodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if xnode exists")
	}

	return count > 0, nil
}

// Xnodes retrieves all the records using an executor.
func Xnodes(mods ...qm.QueryMod) xnodeQuery {
	mods = append(mods, qm.From("\"xnode\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"xnode\".*"})
	}

	return xnodeQuery{q}
}

// FindXnode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindXnode(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Xnode, error) {
	xnodeObj := &Xnode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"xnode\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, xnodeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from xnode")
	}

	if err = xnodeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return xnodeObj, err
	}

	return xnodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Xnode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no xnode provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(xnodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	xnodeInsertCacheMut.RLock()
	cache, cached := xnodeInsertCache[key]
	xnodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			xnodeAllColumns,
			xnodeColumnsWithDefault,
			xnodeColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, xnodeGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(xnodeType, xnodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(xnodeType, xnodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"xnode\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"xnode\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into xnode")
	}

	if !cached {
		xnodeInsertCacheMut.Lock()
		xnodeInsertCache[key] = cache
		xnodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Xnode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Xnode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	xnodeUpdateCacheMut.RLock()
	cache, cached := xnodeUpdateCache[key]
	xnodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			xnodeAllColumns,
			xnodePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, xnodeGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update xnode, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"xnode\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, xnodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(xnodeType, xnodeMapping, append(wl, xnodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update xnode row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for xnode")
	}

	if !cached {
		xnodeUpdateCacheMut.Lock()
		xnodeUpdateCache[key] = cache
		xnodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q xnodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for xnode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for xnode")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o XnodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xnodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"xnode\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xnodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in xnode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all xnode")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Xnode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no xnode provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(xnodeColumnsWithDefault, o)

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

	xnodeUpsertCacheMut.RLock()
	cache, cached := xnodeUpsertCache[key]
	xnodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			xnodeAllColumns,
			xnodeColumnsWithDefault,
			xnodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			xnodeAllColumns,
			xnodePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert xnode, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(xnodePrimaryKeyColumns))
			copy(conflict, xnodePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"xnode\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(xnodeType, xnodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(xnodeType, xnodeMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert xnode")
	}

	if !cached {
		xnodeUpsertCacheMut.Lock()
		xnodeUpsertCache[key] = cache
		xnodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Xnode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Xnode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Xnode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), xnodePrimaryKeyMapping)
	sql := "DELETE FROM \"xnode\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from xnode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for xnode")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q xnodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no xnodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from xnode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for xnode")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o XnodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(xnodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xnodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"xnode\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xnodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from xnode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for xnode")
	}

	if len(xnodeAfterDeleteHooks) != 0 {
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
func (o *Xnode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindXnode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *XnodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := XnodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xnodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"xnode\".* FROM \"xnode\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xnodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in XnodeSlice")
	}

	*o = slice

	return nil
}

// XnodeExists checks if the Xnode row exists.
func XnodeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"xnode\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if xnode exists")
	}

	return exists, nil
}

// Exists checks if the Xnode row exists.
func (o *Xnode) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return XnodeExists(ctx, exec, o.ID)
}

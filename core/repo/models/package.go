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

// Package is an object representing the database table.
type Package struct {
	ID             int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name           string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Duration       int64     `boil:"duration" json:"duration" toml:"duration" yaml:"duration"`
	Price          int64     `boil:"price" json:"price" toml:"price" yaml:"price"`
	TrafficAllowed float32   `boil:"traffic_allowed" json:"traffic_allowed" toml:"traffic_allowed" yaml:"traffic_allowed"`
	ResetMode      string    `boil:"reset_mode" json:"reset_mode" toml:"reset_mode" yaml:"reset_mode"`
	Active         null.Bool `boil:"active" json:"active,omitempty" toml:"active" yaml:"active,omitempty"`

	R *packageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L packageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PackageColumns = struct {
	ID             string
	Name           string
	Duration       string
	Price          string
	TrafficAllowed string
	ResetMode      string
	Active         string
}{
	ID:             "id",
	Name:           "name",
	Duration:       "duration",
	Price:          "price",
	TrafficAllowed: "traffic_allowed",
	ResetMode:      "reset_mode",
	Active:         "active",
}

var PackageTableColumns = struct {
	ID             string
	Name           string
	Duration       string
	Price          string
	TrafficAllowed string
	ResetMode      string
	Active         string
}{
	ID:             "package.id",
	Name:           "package.name",
	Duration:       "package.duration",
	Price:          "package.price",
	TrafficAllowed: "package.traffic_allowed",
	ResetMode:      "package.reset_mode",
	Active:         "package.active",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperfloat32 struct{ field string }

func (w whereHelperfloat32) EQ(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat32) NEQ(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat32) LT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat32) LTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat32) GT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat32) GTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat32) IN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat32) NIN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Bool struct{ field string }

func (w whereHelpernull_Bool) EQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bool) NEQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bool) LT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bool) LTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bool) GT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bool) GTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Bool) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bool) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var PackageWhere = struct {
	ID             whereHelperint64
	Name           whereHelperstring
	Duration       whereHelperint64
	Price          whereHelperint64
	TrafficAllowed whereHelperfloat32
	ResetMode      whereHelperstring
	Active         whereHelpernull_Bool
}{
	ID:             whereHelperint64{field: "\"package\".\"id\""},
	Name:           whereHelperstring{field: "\"package\".\"name\""},
	Duration:       whereHelperint64{field: "\"package\".\"duration\""},
	Price:          whereHelperint64{field: "\"package\".\"price\""},
	TrafficAllowed: whereHelperfloat32{field: "\"package\".\"traffic_allowed\""},
	ResetMode:      whereHelperstring{field: "\"package\".\"reset_mode\""},
	Active:         whereHelpernull_Bool{field: "\"package\".\"active\""},
}

// PackageRels is where relationship names are stored.
var PackageRels = struct {
	Tusers string
}{
	Tusers: "Tusers",
}

// packageR is where relationships are stored.
type packageR struct {
	Tusers TuserSlice `boil:"Tusers" json:"Tusers" toml:"Tusers" yaml:"Tusers"`
}

// NewStruct creates a new relationship struct
func (*packageR) NewStruct() *packageR {
	return &packageR{}
}

func (r *packageR) GetTusers() TuserSlice {
	if r == nil {
		return nil
	}
	return r.Tusers
}

// packageL is where Load methods for each relationship are stored.
type packageL struct{}

var (
	packageAllColumns            = []string{"id", "name", "duration", "price", "traffic_allowed", "reset_mode", "active"}
	packageColumnsWithoutDefault = []string{"name", "duration", "price", "reset_mode"}
	packageColumnsWithDefault    = []string{"id", "traffic_allowed", "active"}
	packagePrimaryKeyColumns     = []string{"id"}
	packageGeneratedColumns      = []string{"id"}
)

type (
	// PackageSlice is an alias for a slice of pointers to Package.
	// This should almost always be used instead of []Package.
	PackageSlice []*Package
	// PackageHook is the signature for custom Package hook methods
	PackageHook func(context.Context, boil.ContextExecutor, *Package) error

	packageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	packageType                 = reflect.TypeOf(&Package{})
	packageMapping              = queries.MakeStructMapping(packageType)
	packagePrimaryKeyMapping, _ = queries.BindMapping(packageType, packageMapping, packagePrimaryKeyColumns)
	packageInsertCacheMut       sync.RWMutex
	packageInsertCache          = make(map[string]insertCache)
	packageUpdateCacheMut       sync.RWMutex
	packageUpdateCache          = make(map[string]updateCache)
	packageUpsertCacheMut       sync.RWMutex
	packageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var packageAfterSelectHooks []PackageHook

var packageBeforeInsertHooks []PackageHook
var packageAfterInsertHooks []PackageHook

var packageBeforeUpdateHooks []PackageHook
var packageAfterUpdateHooks []PackageHook

var packageBeforeDeleteHooks []PackageHook
var packageAfterDeleteHooks []PackageHook

var packageBeforeUpsertHooks []PackageHook
var packageAfterUpsertHooks []PackageHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Package) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Package) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Package) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Package) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Package) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Package) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Package) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Package) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Package) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range packageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPackageHook registers your hook function for all future operations.
func AddPackageHook(hookPoint boil.HookPoint, packageHook PackageHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		packageAfterSelectHooks = append(packageAfterSelectHooks, packageHook)
	case boil.BeforeInsertHook:
		packageBeforeInsertHooks = append(packageBeforeInsertHooks, packageHook)
	case boil.AfterInsertHook:
		packageAfterInsertHooks = append(packageAfterInsertHooks, packageHook)
	case boil.BeforeUpdateHook:
		packageBeforeUpdateHooks = append(packageBeforeUpdateHooks, packageHook)
	case boil.AfterUpdateHook:
		packageAfterUpdateHooks = append(packageAfterUpdateHooks, packageHook)
	case boil.BeforeDeleteHook:
		packageBeforeDeleteHooks = append(packageBeforeDeleteHooks, packageHook)
	case boil.AfterDeleteHook:
		packageAfterDeleteHooks = append(packageAfterDeleteHooks, packageHook)
	case boil.BeforeUpsertHook:
		packageBeforeUpsertHooks = append(packageBeforeUpsertHooks, packageHook)
	case boil.AfterUpsertHook:
		packageAfterUpsertHooks = append(packageAfterUpsertHooks, packageHook)
	}
}

// One returns a single package record from the query.
func (q packageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Package, error) {
	o := &Package{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for package")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Package records from the query.
func (q packageQuery) All(ctx context.Context, exec boil.ContextExecutor) (PackageSlice, error) {
	var o []*Package

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Package slice")
	}

	if len(packageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Package records in the query.
func (q packageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count package rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q packageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if package exists")
	}

	return count > 0, nil
}

// Tusers retrieves all the tuser's Tusers with an executor.
func (o *Package) Tusers(mods ...qm.QueryMod) tuserQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"tuser\".\"package_id\"=?", o.ID),
	)

	return Tusers(queryMods...)
}

// LoadTusers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (packageL) LoadTusers(ctx context.Context, e boil.ContextExecutor, singular bool, maybePackage interface{}, mods queries.Applicator) error {
	var slice []*Package
	var object *Package

	if singular {
		var ok bool
		object, ok = maybePackage.(*Package)
		if !ok {
			object = new(Package)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePackage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePackage))
			}
		}
	} else {
		s, ok := maybePackage.(*[]*Package)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePackage)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePackage))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &packageR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &packageR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tuser`),
		qm.WhereIn(`tuser.package_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load tuser")
	}

	var resultSlice []*Tuser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice tuser")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on tuser")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tuser")
	}

	if len(tuserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Tusers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &tuserR{}
			}
			foreign.R.Package = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PackageID {
				local.R.Tusers = append(local.R.Tusers, foreign)
				if foreign.R == nil {
					foreign.R = &tuserR{}
				}
				foreign.R.Package = local
				break
			}
		}
	}

	return nil
}

// AddTusers adds the given related objects to the existing relationships
// of the package, optionally inserting them as new records.
// Appends related to o.R.Tusers.
// Sets related.R.Package appropriately.
func (o *Package) AddTusers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Tuser) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PackageID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"tuser\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"package_id"}),
				strmangle.WhereClause("\"", "\"", 0, tuserPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PackageID = o.ID
		}
	}

	if o.R == nil {
		o.R = &packageR{
			Tusers: related,
		}
	} else {
		o.R.Tusers = append(o.R.Tusers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tuserR{
				Package: o,
			}
		} else {
			rel.R.Package = o
		}
	}
	return nil
}

// Packages retrieves all the records using an executor.
func Packages(mods ...qm.QueryMod) packageQuery {
	mods = append(mods, qm.From("\"package\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"package\".*"})
	}

	return packageQuery{q}
}

// FindPackage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPackage(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Package, error) {
	packageObj := &Package{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"package\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, packageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from package")
	}

	if err = packageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return packageObj, err
	}

	return packageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Package) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no package provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(packageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	packageInsertCacheMut.RLock()
	cache, cached := packageInsertCache[key]
	packageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			packageAllColumns,
			packageColumnsWithDefault,
			packageColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, packageGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(packageType, packageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(packageType, packageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"package\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"package\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into package")
	}

	if !cached {
		packageInsertCacheMut.Lock()
		packageInsertCache[key] = cache
		packageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Package.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Package) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	packageUpdateCacheMut.RLock()
	cache, cached := packageUpdateCache[key]
	packageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			packageAllColumns,
			packagePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, packageGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update package, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"package\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, packagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(packageType, packageMapping, append(wl, packagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update package row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for package")
	}

	if !cached {
		packageUpdateCacheMut.Lock()
		packageUpdateCache[key] = cache
		packageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q packageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for package")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for package")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PackageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), packagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"package\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, packagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in package slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all package")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Package) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no package provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(packageColumnsWithDefault, o)

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

	packageUpsertCacheMut.RLock()
	cache, cached := packageUpsertCache[key]
	packageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			packageAllColumns,
			packageColumnsWithDefault,
			packageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			packageAllColumns,
			packagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert package, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(packagePrimaryKeyColumns))
			copy(conflict, packagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"package\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(packageType, packageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(packageType, packageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert package")
	}

	if !cached {
		packageUpsertCacheMut.Lock()
		packageUpsertCache[key] = cache
		packageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Package record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Package) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Package provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), packagePrimaryKeyMapping)
	sql := "DELETE FROM \"package\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from package")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for package")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q packageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no packageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from package")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for package")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PackageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(packageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), packagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"package\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, packagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from package slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for package")
	}

	if len(packageAfterDeleteHooks) != 0 {
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
func (o *Package) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPackage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PackageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PackageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), packagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"package\".* FROM \"package\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, packagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PackageSlice")
	}

	*o = slice

	return nil
}

// PackageExists checks if the Package row exists.
func PackageExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"package\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if package exists")
	}

	return exists, nil
}

// Exists checks if the Package row exists.
func (o *Package) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PackageExists(ctx, exec, o.ID)
}

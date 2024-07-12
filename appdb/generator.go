// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package appdb

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Generator is an object representing the database table.
type Generator struct {
	ID              int64   `boil:"id" json:"id" toml:"id" yaml:"id"`
	AssetID         int32   `boil:"asset_id" json:"asset_id" toml:"asset_id" yaml:"asset_id"`
	Attribute       string  `boil:"attribute" json:"attribute" toml:"attribute" yaml:"attribute"`
	Subtype         string  `boil:"subtype" json:"subtype" toml:"subtype" yaml:"subtype"`
	FunctionType    string  `boil:"function_type" json:"function_type" toml:"function_type" yaml:"function_type"`
	MinValue        float64 `boil:"min_value" json:"min_value" toml:"min_value" yaml:"min_value"`
	MaxValue        float64 `boil:"max_value" json:"max_value" toml:"max_value" yaml:"max_value"`
	IntervalSeconds int32   `boil:"interval_seconds" json:"interval_seconds" toml:"interval_seconds" yaml:"interval_seconds"`
	Frequency       float64 `boil:"frequency" json:"frequency" toml:"frequency" yaml:"frequency"`

	R *generatorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L generatorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GeneratorColumns = struct {
	ID              string
	AssetID         string
	Attribute       string
	Subtype         string
	FunctionType    string
	MinValue        string
	MaxValue        string
	IntervalSeconds string
	Frequency       string
}{
	ID:              "id",
	AssetID:         "asset_id",
	Attribute:       "attribute",
	Subtype:         "subtype",
	FunctionType:    "function_type",
	MinValue:        "min_value",
	MaxValue:        "max_value",
	IntervalSeconds: "interval_seconds",
	Frequency:       "frequency",
}

var GeneratorTableColumns = struct {
	ID              string
	AssetID         string
	Attribute       string
	Subtype         string
	FunctionType    string
	MinValue        string
	MaxValue        string
	IntervalSeconds string
	Frequency       string
}{
	ID:              "generator.id",
	AssetID:         "generator.asset_id",
	Attribute:       "generator.attribute",
	Subtype:         "generator.subtype",
	FunctionType:    "generator.function_type",
	MinValue:        "generator.min_value",
	MaxValue:        "generator.max_value",
	IntervalSeconds: "generator.interval_seconds",
	Frequency:       "generator.frequency",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint32 struct{ field string }

func (w whereHelperint32) EQ(x int32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint32) NEQ(x int32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint32) LT(x int32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint32) LTE(x int32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint32) GT(x int32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint32) GTE(x int32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint32) IN(slice []int32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint32) NIN(slice []int32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
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

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var GeneratorWhere = struct {
	ID              whereHelperint64
	AssetID         whereHelperint32
	Attribute       whereHelperstring
	Subtype         whereHelperstring
	FunctionType    whereHelperstring
	MinValue        whereHelperfloat64
	MaxValue        whereHelperfloat64
	IntervalSeconds whereHelperint32
	Frequency       whereHelperfloat64
}{
	ID:              whereHelperint64{field: "\"device_simulator\".\"generator\".\"id\""},
	AssetID:         whereHelperint32{field: "\"device_simulator\".\"generator\".\"asset_id\""},
	Attribute:       whereHelperstring{field: "\"device_simulator\".\"generator\".\"attribute\""},
	Subtype:         whereHelperstring{field: "\"device_simulator\".\"generator\".\"subtype\""},
	FunctionType:    whereHelperstring{field: "\"device_simulator\".\"generator\".\"function_type\""},
	MinValue:        whereHelperfloat64{field: "\"device_simulator\".\"generator\".\"min_value\""},
	MaxValue:        whereHelperfloat64{field: "\"device_simulator\".\"generator\".\"max_value\""},
	IntervalSeconds: whereHelperint32{field: "\"device_simulator\".\"generator\".\"interval_seconds\""},
	Frequency:       whereHelperfloat64{field: "\"device_simulator\".\"generator\".\"frequency\""},
}

// GeneratorRels is where relationship names are stored.
var GeneratorRels = struct {
}{}

// generatorR is where relationships are stored.
type generatorR struct {
}

// NewStruct creates a new relationship struct
func (*generatorR) NewStruct() *generatorR {
	return &generatorR{}
}

// generatorL is where Load methods for each relationship are stored.
type generatorL struct{}

var (
	generatorAllColumns            = []string{"id", "asset_id", "attribute", "subtype", "function_type", "min_value", "max_value", "interval_seconds", "frequency"}
	generatorColumnsWithoutDefault = []string{"asset_id", "attribute", "subtype", "function_type", "min_value", "max_value", "interval_seconds", "frequency"}
	generatorColumnsWithDefault    = []string{"id"}
	generatorPrimaryKeyColumns     = []string{"id"}
	generatorGeneratedColumns      = []string{}
)

type (
	// GeneratorSlice is an alias for a slice of pointers to Generator.
	// This should almost always be used instead of []Generator.
	GeneratorSlice []*Generator
	// GeneratorHook is the signature for custom Generator hook methods
	GeneratorHook func(context.Context, boil.ContextExecutor, *Generator) error

	generatorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	generatorType                 = reflect.TypeOf(&Generator{})
	generatorMapping              = queries.MakeStructMapping(generatorType)
	generatorPrimaryKeyMapping, _ = queries.BindMapping(generatorType, generatorMapping, generatorPrimaryKeyColumns)
	generatorInsertCacheMut       sync.RWMutex
	generatorInsertCache          = make(map[string]insertCache)
	generatorUpdateCacheMut       sync.RWMutex
	generatorUpdateCache          = make(map[string]updateCache)
	generatorUpsertCacheMut       sync.RWMutex
	generatorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var generatorAfterSelectMu sync.Mutex
var generatorAfterSelectHooks []GeneratorHook

var generatorBeforeInsertMu sync.Mutex
var generatorBeforeInsertHooks []GeneratorHook
var generatorAfterInsertMu sync.Mutex
var generatorAfterInsertHooks []GeneratorHook

var generatorBeforeUpdateMu sync.Mutex
var generatorBeforeUpdateHooks []GeneratorHook
var generatorAfterUpdateMu sync.Mutex
var generatorAfterUpdateHooks []GeneratorHook

var generatorBeforeDeleteMu sync.Mutex
var generatorBeforeDeleteHooks []GeneratorHook
var generatorAfterDeleteMu sync.Mutex
var generatorAfterDeleteHooks []GeneratorHook

var generatorBeforeUpsertMu sync.Mutex
var generatorBeforeUpsertHooks []GeneratorHook
var generatorAfterUpsertMu sync.Mutex
var generatorAfterUpsertHooks []GeneratorHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Generator) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Generator) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Generator) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Generator) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Generator) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Generator) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Generator) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Generator) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Generator) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range generatorAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGeneratorHook registers your hook function for all future operations.
func AddGeneratorHook(hookPoint boil.HookPoint, generatorHook GeneratorHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		generatorAfterSelectMu.Lock()
		generatorAfterSelectHooks = append(generatorAfterSelectHooks, generatorHook)
		generatorAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		generatorBeforeInsertMu.Lock()
		generatorBeforeInsertHooks = append(generatorBeforeInsertHooks, generatorHook)
		generatorBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		generatorAfterInsertMu.Lock()
		generatorAfterInsertHooks = append(generatorAfterInsertHooks, generatorHook)
		generatorAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		generatorBeforeUpdateMu.Lock()
		generatorBeforeUpdateHooks = append(generatorBeforeUpdateHooks, generatorHook)
		generatorBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		generatorAfterUpdateMu.Lock()
		generatorAfterUpdateHooks = append(generatorAfterUpdateHooks, generatorHook)
		generatorAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		generatorBeforeDeleteMu.Lock()
		generatorBeforeDeleteHooks = append(generatorBeforeDeleteHooks, generatorHook)
		generatorBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		generatorAfterDeleteMu.Lock()
		generatorAfterDeleteHooks = append(generatorAfterDeleteHooks, generatorHook)
		generatorAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		generatorBeforeUpsertMu.Lock()
		generatorBeforeUpsertHooks = append(generatorBeforeUpsertHooks, generatorHook)
		generatorBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		generatorAfterUpsertMu.Lock()
		generatorAfterUpsertHooks = append(generatorAfterUpsertHooks, generatorHook)
		generatorAfterUpsertMu.Unlock()
	}
}

// OneG returns a single generator record from the query using the global executor.
func (q generatorQuery) OneG(ctx context.Context) (*Generator, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single generator record from the query.
func (q generatorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Generator, error) {
	o := &Generator{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "appdb: failed to execute a one query for generator")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Generator records from the query using the global executor.
func (q generatorQuery) AllG(ctx context.Context) (GeneratorSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Generator records from the query.
func (q generatorQuery) All(ctx context.Context, exec boil.ContextExecutor) (GeneratorSlice, error) {
	var o []*Generator

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "appdb: failed to assign all query results to Generator slice")
	}

	if len(generatorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Generator records in the query using the global executor
func (q generatorQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Generator records in the query.
func (q generatorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to count generator rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q generatorQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q generatorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "appdb: failed to check if generator exists")
	}

	return count > 0, nil
}

// Generators retrieves all the records using an executor.
func Generators(mods ...qm.QueryMod) generatorQuery {
	mods = append(mods, qm.From("\"device_simulator\".\"generator\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"device_simulator\".\"generator\".*"})
	}

	return generatorQuery{q}
}

// FindGeneratorG retrieves a single record by ID.
func FindGeneratorG(ctx context.Context, iD int64, selectCols ...string) (*Generator, error) {
	return FindGenerator(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindGenerator retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGenerator(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Generator, error) {
	generatorObj := &Generator{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"device_simulator\".\"generator\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, generatorObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "appdb: unable to select from generator")
	}

	if err = generatorObj.doAfterSelectHooks(ctx, exec); err != nil {
		return generatorObj, err
	}

	return generatorObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Generator) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Generator) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("appdb: no generator provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(generatorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	generatorInsertCacheMut.RLock()
	cache, cached := generatorInsertCache[key]
	generatorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			generatorAllColumns,
			generatorColumnsWithDefault,
			generatorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(generatorType, generatorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(generatorType, generatorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"device_simulator\".\"generator\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"device_simulator\".\"generator\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "appdb: unable to insert into generator")
	}

	if !cached {
		generatorInsertCacheMut.Lock()
		generatorInsertCache[key] = cache
		generatorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Generator record using the global executor.
// See Update for more documentation.
func (o *Generator) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Generator.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Generator) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	generatorUpdateCacheMut.RLock()
	cache, cached := generatorUpdateCache[key]
	generatorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			generatorAllColumns,
			generatorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("appdb: unable to update generator, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"device_simulator\".\"generator\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, generatorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(generatorType, generatorMapping, append(wl, generatorPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "appdb: unable to update generator row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by update for generator")
	}

	if !cached {
		generatorUpdateCacheMut.Lock()
		generatorUpdateCache[key] = cache
		generatorUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q generatorQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q generatorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to update all for generator")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to retrieve rows affected for generator")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GeneratorSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GeneratorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("appdb: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), generatorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"device_simulator\".\"generator\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, generatorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to update all in generator slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to retrieve rows affected all in update all generator")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Generator) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Generator) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("appdb: no generator provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(generatorColumnsWithDefault, o)

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

	generatorUpsertCacheMut.RLock()
	cache, cached := generatorUpsertCache[key]
	generatorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			generatorAllColumns,
			generatorColumnsWithDefault,
			generatorColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			generatorAllColumns,
			generatorPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("appdb: unable to upsert generator, could not build update column list")
		}

		ret := strmangle.SetComplement(generatorAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(generatorPrimaryKeyColumns) == 0 {
				return errors.New("appdb: unable to upsert generator, could not build conflict column list")
			}

			conflict = make([]string, len(generatorPrimaryKeyColumns))
			copy(conflict, generatorPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"device_simulator\".\"generator\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(generatorType, generatorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(generatorType, generatorMapping, ret)
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
		return errors.Wrap(err, "appdb: unable to upsert generator")
	}

	if !cached {
		generatorUpsertCacheMut.Lock()
		generatorUpsertCache[key] = cache
		generatorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Generator record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Generator) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Generator record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Generator) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("appdb: no Generator provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), generatorPrimaryKeyMapping)
	sql := "DELETE FROM \"device_simulator\".\"generator\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete from generator")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by delete for generator")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q generatorQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q generatorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("appdb: no generatorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete all from generator")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by deleteall for generator")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o GeneratorSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GeneratorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(generatorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), generatorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"device_simulator\".\"generator\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, generatorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete all from generator slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by deleteall for generator")
	}

	if len(generatorAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Generator) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("appdb: no Generator provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Generator) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGenerator(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GeneratorSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("appdb: empty GeneratorSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GeneratorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GeneratorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), generatorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"device_simulator\".\"generator\".* FROM \"device_simulator\".\"generator\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, generatorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "appdb: unable to reload all in GeneratorSlice")
	}

	*o = slice

	return nil
}

// GeneratorExistsG checks if the Generator row exists.
func GeneratorExistsG(ctx context.Context, iD int64) (bool, error) {
	return GeneratorExists(ctx, boil.GetContextDB(), iD)
}

// GeneratorExists checks if the Generator row exists.
func GeneratorExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"device_simulator\".\"generator\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "appdb: unable to check if generator exists")
	}

	return exists, nil
}

// Exists checks if the Generator row exists.
func (o *Generator) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GeneratorExists(ctx, exec, o.ID)
}

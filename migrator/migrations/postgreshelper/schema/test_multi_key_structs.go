// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"
	"time"

	"github.com/lib/pq"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

var (
	// CreateTableTestMultiKeyStructsStmt holds the create statement for table `test_multi_key_structs`.
	CreateTableTestMultiKeyStructsStmt = &postgres.CreateStmts{
		GormModel: (*TestMultiKeyStructs)(nil),
		Children: []*postgres.CreateStmts{
			&postgres.CreateStmts{
				GormModel: (*TestMultiKeyStructsNesteds)(nil),
				Children:  []*postgres.CreateStmts{},
			},
		},
	}

	// TestMultiKeyStructsSchema is the go schema for table `test_multi_key_structs`.
	TestMultiKeyStructsSchema = func() *walker.Schema {
		schema := walker.Walk(reflect.TypeOf((*storage.TestMultiKeyStruct)(nil)), "test_multi_key_structs")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_SEARCH_UNSET, "testmultikeystruct", (*storage.TestMultiKeyStruct)(nil)))
		return schema
	}()
)

const (
	TestMultiKeyStructsTableName        = "test_multi_key_structs"
	TestMultiKeyStructsNestedsTableName = "test_multi_key_structs_nesteds"
)

// TestMultiKeyStructs holds the Gorm model for Postgres table `test_multi_key_structs`.
type TestMultiKeyStructs struct {
	Key1              string                          `gorm:"column:key1;type:varchar;primaryKey"`
	Key2              string                          `gorm:"column:key2;type:varchar;primaryKey"`
	StringSlice       *pq.StringArray                 `gorm:"column:stringslice;type:text[]"`
	Bool              bool                            `gorm:"column:bool;type:bool"`
	Uint64            uint64                          `gorm:"column:uint64;type:integer"`
	Int64             int64                           `gorm:"column:int64;type:integer"`
	Float             float32                         `gorm:"column:float;type:numeric"`
	Labels            map[string]string               `gorm:"column:labels;type:jsonb"`
	Timestamp         *time.Time                      `gorm:"column:timestamp;type:timestamp"`
	Enum              storage.TestMultiKeyStruct_Enum `gorm:"column:enum;type:integer"`
	Enums             *pq.Int32Array                  `gorm:"column:enums;type:int[]"`
	String            string                          `gorm:"column:string_;type:varchar"`
	Int32Slice        *pq.Int32Array                  `gorm:"column:int32slice;type:int[]"`
	OneofnestedNested string                          `gorm:"column:oneofnested_nested;type:varchar"`
	Serialized        []byte                          `gorm:"column:serialized;type:bytea"`
}

// TestMultiKeyStructsNesteds holds the Gorm model for Postgres table `test_multi_key_structs_nesteds`.
type TestMultiKeyStructsNesteds struct {
	TestMultiKeyStructsKey1 string              `gorm:"column:test_multi_key_structs_key1;type:varchar;primaryKey"`
	TestMultiKeyStructsKey2 string              `gorm:"column:test_multi_key_structs_key2;type:varchar;primaryKey"`
	Idx                     int                 `gorm:"column:idx;type:integer;primaryKey;index:testmultikeystructsnesteds_idx,type:btree"`
	Nested                  string              `gorm:"column:nested;type:varchar"`
	IsNested                bool                `gorm:"column:isnested;type:bool"`
	Int64                   int64               `gorm:"column:int64;type:integer"`
	Nested2Nested2          string              `gorm:"column:nested2_nested2;type:varchar"`
	Nested2IsNested         bool                `gorm:"column:nested2_isnested;type:bool"`
	Nested2Int64            int64               `gorm:"column:nested2_int64;type:integer"`
	TestMultiKeyStructsRef  TestMultiKeyStructs `gorm:"foreignKey:test_multi_key_structs_key1,test_multi_key_structs_key2;references:key1,key2;belongsTo;constraint:OnDelete:CASCADE"`
}
/*
Copyright 2021 The Vitess Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fuzzing

import (
	"context"
	"sync"
	"testing"
	"time"

	"mdibaiee/vitess/go/mysql/collations"
	"mdibaiee/vitess/go/mysql/fakesqldb"
	"mdibaiee/vitess/go/sqltypes"
	"mdibaiee/vitess/go/vt/dbconfigs"
	"mdibaiee/vitess/go/vt/vtenv"
	"mdibaiee/vitess/go/vt/vttablet/tabletserver/connpool"
	"mdibaiee/vitess/go/vt/vttablet/tabletserver/schema"
	"mdibaiee/vitess/go/vt/vttablet/tabletserver/tabletenv"

	fuzz "github.com/AdaLogics/go-fuzz-headers"
)

var initter sync.Once

func FuzzLoadTable(data []byte) int {
	initter.Do(initTesting)
	f := fuzz.NewConsumer(data)
	tableName, err := f.GetString()
	if err != nil {
		return 0
	}
	comment, err := f.GetString()
	if err != nil {
		return 0
	}
	query, err := f.GetSQLString()
	if err != nil {
		return 0
	}

	t := &testing.T{}

	db := fakesqldb.New(t)
	defer db.Close()
	db.AddQuery(query, &sqltypes.Result{})

	_, _ = newTestLoadTable(tableName, comment, db)
	return 1
}

func newTestLoadTable(tableName, comment string, db *fakesqldb.DB) (*schema.Table, error) {
	ctx := context.Background()
	appParams := dbconfigs.New(db.ConnParams())
	dbaParams := dbconfigs.New(db.ConnParams())
	cfg := tabletenv.ConnPoolConfig{
		Size:        2,
		IdleTimeout: 10 * time.Second,
	}

	connPool := connpool.NewPool(tabletenv.NewEnv(vtenv.NewTestEnv(), nil, "SchemaTest"), "", cfg)
	connPool.Open(appParams, dbaParams, appParams)
	conn, err := connPool.Get(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Recycle()

	return schema.LoadTable(conn, "fakesqldb", tableName, "BASE_TABLE", comment, collations.MySQL8())
}

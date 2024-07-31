package schemadiff

import (
	"github.com/estuary/vitess/go/mysql/collations"
	"github.com/estuary/vitess/go/vt/vtenv"
)

type Environment struct {
	*vtenv.Environment
	DefaultColl collations.ID
}

func NewTestEnv() *Environment {
	return &Environment{
		Environment: vtenv.NewTestEnv(),
		DefaultColl: collations.MySQL8().DefaultConnectionCharset(),
	}
}

func NewEnv(env *vtenv.Environment, defaultColl collations.ID) *Environment {
	return &Environment{
		Environment: env,
		DefaultColl: defaultColl,
	}
}

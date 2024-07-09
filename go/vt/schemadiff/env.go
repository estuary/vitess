package schemadiff

import (
	"mdibaiee/vitess/oracle/go/mysql/collations"
	"mdibaiee/vitess/oracle/go/vt/vtenv"
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

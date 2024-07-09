/*
Copyright 2020 The Vitess Authors.

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

package main

import (
	"flag"

	"mdibaiee/vitess/oracle/go/acl"
	"mdibaiee/vitess/oracle/go/cmd/vtctldclient/command"
	"mdibaiee/vitess/oracle/go/exit"
	"mdibaiee/vitess/oracle/go/vt/grpcclient"
	"mdibaiee/vitess/oracle/go/vt/grpccommon"
	"mdibaiee/vitess/oracle/go/vt/log"
	"mdibaiee/vitess/oracle/go/vt/logutil"
	"mdibaiee/vitess/oracle/go/vt/servenv"
	"mdibaiee/vitess/oracle/go/vt/vtctl/grpcclientcommon"
	"mdibaiee/vitess/oracle/go/vt/vtctl/vtctlclient"

	_flag "mdibaiee/vitess/oracle/go/internal/flag"
)

func main() {
	defer exit.Recover()

	// Grab all those global flags across the codebase and shove 'em on in.
	// (TODO|andrew) remove this line after the migration to pflag is complete.
	command.Root.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	log.RegisterFlags(command.Root.PersistentFlags())
	logutil.RegisterFlags(command.Root.PersistentFlags())
	grpcclient.RegisterFlags(command.Root.PersistentFlags())
	grpccommon.RegisterFlags(command.Root.PersistentFlags())
	grpcclientcommon.RegisterFlags(command.Root.PersistentFlags())
	servenv.RegisterMySQLServerFlags(command.Root.PersistentFlags())
	vtctlclient.RegisterFlags(command.Root.PersistentFlags())
	acl.RegisterFlags(command.Root.PersistentFlags())

	// hack to get rid of an "ERROR: logging before flag.Parse"
	_flag.TrickGlog()

	// back to your regularly scheduled cobra programming
	if err := command.Root.Execute(); err != nil {
		log.Error(err)
		exit.Return(1)
	}
}

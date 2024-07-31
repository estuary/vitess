package main

import (
	"log"

	"github.com/estuary/vitess/go/acl"
	"github.com/estuary/vitess/go/cmd/rulesctl/cmd"
	vtlog "github.com/estuary/vitess/go/vt/log"
	"github.com/estuary/vitess/go/vt/logutil"
	"github.com/estuary/vitess/go/vt/servenv"
)

func main() {
	rootCmd := cmd.Main()
	vtlog.RegisterFlags(rootCmd.PersistentFlags())
	logutil.RegisterFlags(rootCmd.PersistentFlags())
	acl.RegisterFlags(rootCmd.PersistentFlags())
	servenv.RegisterMySQLServerFlags(rootCmd.PersistentFlags())
	if err := rootCmd.Execute(); err != nil {
		log.Printf("%v", err)
	}
}

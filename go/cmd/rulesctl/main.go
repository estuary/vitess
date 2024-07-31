package main

import (
	"log"

	"github.com/vitess/vitess/go/acl"
	"github.com/vitess/vitess/go/cmd/rulesctl/cmd"
	vtlog "github.com/vitess/vitess/go/vt/log"
	"github.com/vitess/vitess/go/vt/logutil"
	"github.com/vitess/vitess/go/vt/servenv"
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

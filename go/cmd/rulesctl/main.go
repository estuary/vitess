package main

import (
	"log"

	"mdibaiee/vitess/oracle/go/acl"
	"mdibaiee/vitess/oracle/go/cmd/rulesctl/cmd"
	vtlog "mdibaiee/vitess/oracle/go/vt/log"
	"mdibaiee/vitess/oracle/go/vt/logutil"
	"mdibaiee/vitess/oracle/go/vt/servenv"
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

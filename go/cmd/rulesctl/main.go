package main

import (
	"log"

	"mdibaiee/vitess/go/acl"
	"mdibaiee/vitess/go/cmd/rulesctl/cmd"
	vtlog "mdibaiee/vitess/go/vt/log"
	"mdibaiee/vitess/go/vt/logutil"
	"mdibaiee/vitess/go/vt/servenv"
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

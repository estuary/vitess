package main

import (
	"log"

	"github.com/mdibaiee/vitess/go/acl"
	"github.com/mdibaiee/vitess/go/cmd/rulesctl/cmd"
	vtlog "github.com/mdibaiee/vitess/go/vt/log"
	"github.com/mdibaiee/vitess/go/vt/logutil"
	"github.com/mdibaiee/vitess/go/vt/servenv"
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

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	/*
		goto https://app.supabase.com/project/<project-id>/settings/database
		database URI will be like postgresql://<user-name>:<password>@db.<project-id>.supabase.co:5432/postgres
	*/
	var dbFlag string
	flag.StringVar(&dbFlag, "db", "", "database URI")
	flag.Parse()

	conn, err := pgx.Connect(context.Background(), dbFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var userid int
	var query string
	err = conn.QueryRow(context.Background(), "select userid, query from pg_stat_statements limit 1").Scan(&userid, &query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(userid, query)
}

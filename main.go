package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func main() {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://admin:1234@localhost:5432/burndown_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name int64
	var weight string
	err = conn.QueryRow(context.Background(), "SELECT c_id, c_name FROM public.t_ref_tasks_source WHERE c_id=$1", 30).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}

type taskSource struct {
	id   int64
	name string
}

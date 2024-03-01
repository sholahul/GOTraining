//go:build ignore
// +build ignore

// Generate the table of OID values
// Run with 'go run gen.go'.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	_ "github.com/lib/pq"
)

// OID represent a postgres Object Identifier Type.
type OID struct {
	ID   int
	Type string
}

// Name returns an upper case version of the oid type.
func (o OID) Name() string {
	return strings.ToUpper(o.Type)
}

func main() {
	datname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")

	if datname == "" {
		os.Setenv("PGDATABASE", "pqgotest")
	}

	if sslmode == "" {
		os.Setenv("PGSSLMODE", "disable")
	}

	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(`
		SELECT typname, oid
		FROM pg_type WHERE oid < 10000
		ORDER BY oid;
	`)
	if err != nil {
		log.Fatal(err)
	}
	oids := make([]*OID, 0)
	for rows.Next() {
		var oid OID
		if err = rows.Scan(&oid.Type, &oid.ID); err != nil {
			log.Fatal(err)
		}
		oids = append(oids, &oid)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("gofmt")
	cmd.Stderr = os.Stderr
	w, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("types.go")
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stdout = f
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "// Code generated by gen.go. DO NOT EDIT.")
	fmt.Fprintln(w, "\npackage oid")
	fmt.Fprintln(w, "const (")
	for _, oid := range oids {
		fmt.Fprintf(w, "T_%s Oid = %d\n", oid.Type, oid.ID)
	}
	fmt.Fprintln(w, ")")
	fmt.Fprintln(w, "var TypeName = map[Oid]string{")
	for _, oid := range oids {
		fmt.Fprintf(w, "T_%s: \"%s\",\n", oid.Type, oid.Name())
	}
	fmt.Fprintln(w, "}")
	w.Close()
	cmd.Wait()
}

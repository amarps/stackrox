package postgres

//go:generate pg-table-bindings-wrapper --type=storage.Group --table=groups --get-all-func --migration-seq 56 --migrate-from boltdb

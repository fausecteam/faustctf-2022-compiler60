module checker

// version on debian bullseye
go 1.15

replace executor => ../executor

require (
	executor v0.0.0-00010101000000-000000000000
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220415214452-44d87bc1d130
	github.com/fausecteam/ctf-gameserver/go/checkerlib v0.0.0-20220430151133-71c92e339f2b
	golang.org/x/sys v0.0.0-20220429121018-84afa8d3f7b3
)

replace github.com/criyle/go-sandbox => github.com/nename0/go-sandbox v0.9.7-0.20220530075807-010449658a9d

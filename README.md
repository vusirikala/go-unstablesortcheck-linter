### Go Linter
This `go` linter checks the occurences of `sort.Sort` and `sort.Slice` in `go` code and suggests to replace them with more stable versions `sort.Stable` and `sort.SliceStable`.

### Example
Run the below command to test the linter on `testdata/foo.go` file. The linter should detect an occurence of `sort.Sort` and `sort.Slice`. 
```
go run ./cmd/unstablesortlint/main.go -- ./testdata/foo.go
```
# go-form-3-custom-linter (wip)

A go based library to enable custom linters.

### What is implemented?
Some example checks like. 
1. detect usages of `time.Sleep`
2. detect usages of `log.WithField` which don't have 'Field' argument as constant

These are just samples for understanding how go analysis works.

### Usage
1. Build the project:

   ```go build github.com/go-form3-custom-linter/cmd/customlint```

   This will generate a ***customlint*** executable in the root.
2. Execute the linter using `go vet` command:

   ```go vet -vettool=$(which ./customlint) ./testdata/timeusage```

### Wip:
1. Invoke "customlint" directly, without specifying `go vet` command. 
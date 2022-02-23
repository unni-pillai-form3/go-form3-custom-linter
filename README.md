# go-form-3-custom-linter (wip)

A go based library to enable custom linters.

### What is implemented?
Some example checks like. 
1. detect usages of `time.Sleep`
2. detect usages of `log.WithField` which don't have 'Field' argument as constant

These are just samples for understanding how go analysis works.
## builder-pattern-codegen

This tool generates builder-pattern code for golang.

### How to build
To build executable, execute following command:

`go build`

## How to generate builder-pattern code
To generate builder-pattern code for your structure, define your structure in following way:
```
package PACKAGE_NAME

type STRUCTURE_NAME struct {
    field1 type // comment
    field2 type //comment
}
```

Define your package name and license in following way:
```
//
// LICENSE 
//
```

To generate code, use following command:

`builder-pattern-codegen  -file=./example/structure  -dir=$PWD/generated -boilerplate=./example/boilerplate`

*Note: Make sure directory exists.*

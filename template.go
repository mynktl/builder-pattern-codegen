package main

var BuildVarTemplate = `
// With$Var method fills the $Var field of $NewObj object.
func ($obj *$newobj) With$Var($var $iType) *$NewObj {
	$obj.$var = $var
    return $obj
}
`
var BuildTemplate = `
// New$Struct returns new instance of object $Struct
func New$Struct() *$NewObj {
	return &$Struct{}
}
`

var SetVarTemplate = `
// Set$Var method set the $Var field of $NewObj object.
func ($obj *$newobj) Set$Var($var $iType) {
	$obj.$var = $var
}
`
var GetVarTemplate = `
// Get$Var method get the $Var field of $NewObj object.
func ($obj *$newobj) Get$Var() $iType {
	return $obj.$var
}
`
var PredicateVarTemplate = `
// IsSet$Var method check if the $Var field of $NewObj object is set.
func Is$VarSet() PredicateFunc {
	return func(c *$newobj) bool {
		return $cond
    }
}
`
var PredicateTemplate = `
// PredicateFunc defines data-type for validation function
type PredicateFunc func(*$newobj) bool

`

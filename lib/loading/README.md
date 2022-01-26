# Loading phase

## Public interface

- T_PkgLoad (Pkg) -> []LoadedRule
- T_RuleLoad (Label) -> LoadedRule

## Private

- T_BzlLoad (Label) -> Starlark objects
- T_GlobLoad (Pkg, string) -> []string

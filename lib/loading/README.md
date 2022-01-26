# Loading phase

## Public interface

- T_Pkg (Pkg) -> []string
- T_LoadTarget (Label) -> Target

## Private

- T_BzlLoad (Label) -> Starlark globals
- T_GlobLoad (Pkg, string) -> []string

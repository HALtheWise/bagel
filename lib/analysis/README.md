## Public interface

- T_Providers (clabel) -> []Provider
- T_Action (clabel, int) -> Action(cmd, input refs)
- T_File (clabel, int) -> File(refs generating action)

## Private tasks

- T_AnalyzeTarget (clabel) -> AnalyzedTarget(providers, actions, files)

using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("graph");
$Go.import("github.com/HALtheWise/bagel/internal/dcache/graph");

# When bagel is invoked:
# - Increment the database epoch
# 	- Get the list of changed files since the last epoch
# 	- Store the token returned by the file watcher
# 	- Find FileStatInfo for all those in the database and set maybe_dirty=true
# 	- Propagate maybe_dirty up the tree through dependents
# - Do some work (query, bUIld, etc.)
# 	- To get an output
# 		- Hash arguments. If in cache and result in cache
# 			- If maybe_dirty = false, return result
# 			- else

# Need dependents for everything
# Need dependencies for ... everything?
# Need input hash if you want cache hits after cache misses in the same function
# Need input values for things queried late in expensive functions (so we can re-execute without running the parent function)
# Need output hashes for things that might not change even when their input does, called by expensive things
# Need output values for things that are expensive to compute _and_ whose children are expensive to compute.

# Between every two expensive functions, need something with output value stored and at least hash of input
	# Unless the two functions always get invalidated together anyway
# Wherever possible, 

# Leaves: T_FileStats, T_FileContentsInfo (Label:data -> Proto:data)
# 	- GetFilepath/ReadFile() based on FileContents
# T_EvalStarlark, wrapped by:
# 	- T_StarlarkGetSimple (Label+str:data -> Proto:data) - gets rule definitions, constants, etc.
# 	- GetStarlark helper function, calls either GetSimple or EvalStarlark
# 	- T_StarlarkListRules (Package:none -> list(string):data)
#	- T_StarlarkGetRule (Label:data -> UnconfiguredRule:data)
# T_ConfigureRule (UnconfiguredRule+config:hash -> ConfiguredRule:hash)
# T_AnalyzeRule, wrapped by:
# 	- T_RuleGetProviderSimple (Label+ProviderID:data -> Proto:data)
# 	- RuleGetProvider helper function
# 	- T_RuleListProviders (Label:none -> list(string):data) - rare, do later
# 	- special something for depsets that are only embedded in other depsets...
# 	- T_RuleGetAction (Label+int: -> Action: )
# 	- T_RuleNumActions (Label: -> int: )

# Interning tables for Label, Config, depset, File


# For things that are weird Starlark objects but we want to cache (rule defs, providers, depsets, Args...)
# 	- When the object is constructed, it gives itself an ID containing the (label, etc) of the active thread, and saves to a dict in thread-local storage
# 	- Given an ID, you can _either_ 
#		- call a Task that gives a simplified Proto object with basic stuff
# 		- call a Task that returns the full Starlark object (with lambdas and shit)
# 	- The former can invoke the latter if something weird's going on

# RefData is an entry in the Refs table. See refs.go for information on interpretation.
struct RefData {
	left @0 :UInt32;
	right @1 :UInt32;
}

struct FuncObj {
	kind @1 :UInt32; # Kind of function call, determined from initialization order. Could probably be UInt8
	arg @0 :UInt32; # Index pointing into Refs table
	dependencies @2 :List(UInt32); # List of indexes into the Funcs table
	result @3 :AnyPointer; # Capnp object storing a *Result (see below). May be null if result is not stored in file cache.
}


struct DiskCache {
	version @0 :Text; # Code version that populated this file
	refs @1 :List(RefData); # Mapping from ref index to ref data
	strings @2 :List(UInt8); # Blob storing strings data accessed through StringRefs
	stringsUsed @4 :UInt32; # Number of bytes of strings that have been filled
	funcs @3 :List(FuncObj); # Mapping from func index to func data
}

## Cached result types for FuncObj

struct RefResult {
	ref @0 :UInt32; # Index into Refs table
}

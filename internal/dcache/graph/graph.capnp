using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("graph");
$Go.import("github.com/HALtheWise/balez/internal/dcache/graph");

# When balez is invoked:
# - Increment the database epoch
# 	- Get the list of changed files since the last epoch
# 	- Store the token returned by the file watcher
# 	- Find FileStatInfo for all those in the database and set maybe_dirty=true
# 	- Propagate maybe_dirty up the tree through dependents
# - Do some work (query, build, etc.)
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


struct CacheEntry {
	id @0 :Int32; 
	# A unique ID. Maybe can be deleted in the future, and just store indexes.

	maybeDirty @1 :Bool;
	# Set to True if this cache entry is potentially stale.

	changedAt @2 :Int32;
	# The cache epoch at which the cached value was most recently updated with a new value

	# recomputed_at @5 :Int32;
	# The cache epoch at which this value was most recently verified correct by recomputing it

	dependencies @4 :List(Int32);
	# Cache entries of the data that went into computing this

	dependents @5 :List(Int32);
	# Other cache entries that use the data from this
	# Needed for propagating maybeDirty

	argsHash @6 :Int64;
	# Trustworthy hash of arguments to this function

	result @7 :AnyPointer;
	# The result to return


	# Debugging fields, may be absent

	name @3 :Text;
	# Optional descriptor of generating function
}

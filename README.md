## What if I reimplemented [Bazel](https://bazel.build/) from scratch? How hard can it be?

### Design

Try to find the simple core of Bazel.

Build one big task graph lazily, using a recursive-function-call strategy like what rust-analyzer does.

Instead of an explicit graph representation, just stick with memoized function call results.

- This is what rust-analyzer does, and it seems effective there
- This _should_ simplify implementation a lot
- Targets, Rules, and Actions are all memoized function calls
- Maybe later have a graph too for garbage collection and stuff

### Milestones

- Start without incremental builds (but with diamond deduplication)
- Start without a bunch of other stuff, I guess?

### mmap cache

- Unlike Bazel, _no_ persistent server process
- Instead, store (much of) the fine-grained cache into a file upon exit
- When the program starts, mmap that file and read from it on demand
  - If the program is being run several times in quick succession, the OS will probably still have the file in memory and this will be like shared memory
  - Inspired by how git handles the index, and by [rsc's csearch](https://swtch.com/~rsc/regexp/regexp4.html)
- Each file is a "cache layer" and refers to the next file up the chain (from the previous execution, presumably)
- At some point, we need a compaction algorithm to drop data from old files that has been invalidated by new layers. This can maybe happen in the background?
- First pass is to use [Cap’n Proto](https://capnproto.org/) for the file format, since it can perform zero-copy random access to structured data
- On top of that, we'll need a hash table implementation of some sort, probably a linear probing table for simplicity.
- Notably, this means we can't cache unserializable nodes between runs. This may require some careful graph design.

### Are you mad?

- Probably, but...
- Bazel has _lots_ of historical baggage related to native rules, which we aren't replicating
- Bazel has lots of code to watch files cross-platform, we will just use [watchman](https://facebook.github.io/watchman/)
- Bazel has complexity to support local sandboxing, we'll just use the Remote Execution API
- We _hopefully_ don't need 100% compatability to make something useful
- _maybe_ the cache files will be simpler than a server process?

### Setup

Developed against go commit `c894b442d1e5e150ad33fa3ce13dbfab1c037b3a`
Later commits break `go_test()`, see https://github.com/bazelbuild/rules_go/issues/2964
https://github.com/golang/tools/blob/master/gopls/doc/advanced.md#working-with-generic-code

`sudo apt install capnproto`

Fish:

```
bazel build @com_zombiezen_go_capnproto2//capnpc-go
begin
  set -lx PATH $PATH .bazel-bin/external/com_zombiezen_go_capnproto2/capnpc-go/capnpc-go_/
  capnp compile -ogo internal/dcache/graph/graph.capnp -I .bazel-bagel/external/com_zombiezen_go_capnproto2/std/
end
```

### Reading

- https://paper.dropbox.com/doc/Salsa-In-More-Depth-2019.01-wYaiL8t72BdGGPFrnmt7h and https://www.youtube.com/watch?v=i_IhACacPRY
- https://blog.nelhage.com/post/why-sorbet-is-fast
- http://neugierig.org/software/blog/2020/05/ninja.html

### Sorbet-inspired fast design

- All hashable values are 32 bit refs indexing into tables of immutable values
  - LabelRef, used for input filepaths and unconfigured rules
    - i.e. argument to label or label_list attr immediately converted to LabelRef
    - Can cheaply get from LabelRef -> LabelRef with "" as last section
  - ConfigRef, representing a build configuration
  - ActionRef, representing an action produced by a rule implementation
- All memo calls take at most two refs as their arguments
  - memotable = (funckind, arg1, arg2) -> (metadata, result?)
  -

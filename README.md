## What if I reimplemented [Bazel](https://bazel.build/) from scratch? How hard can it be?

### Design

Try to find the simple core of Bazel by supporting _only_ starlark rules.

Build one big task graph lazily, using a recursive-function-call strategy like what rust-analyzer does.

Instead of an explicit graph representation, just stick with memoized function call results.

- This is what rust-analyzer does, and it seems effective there
- This _should_ simplify implementation a lot
- Targets, Rules, and Actions are all memoized function calls
- Maybe later have a graph too for garbage collection and stuff

### Milestones

- Start without incremental builds (but with diamond deduplication)
- Start without a bunch of other stuff, I guess?

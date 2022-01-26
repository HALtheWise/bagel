load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "starlark_tasks",
    srcs = [
        "actions.go",
        "builtin_struct.go",
        "ctx.go",
        "depset.go",
        "evaluation.go",
        "files.go",
        "provider.go",
        "read_file.go",
        "rule.go",
    ],
    importpath = "github.com/HALtheWise/bagel/lib/starlark_tasks",
    visibility = ["//:__subpackages__"],
    deps = [
        "//lib/core",
        "//lib/refs",
        "@net_starlark_go//starlark",
    ],
)

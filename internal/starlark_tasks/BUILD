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
    importpath = "github.com/HALtheWise/bagel/internal/starlark_tasks",
    visibility = ["//:__subpackages__"],
    deps = [
        "//internal/labels",
        "//internal/task",
        "@net_starlark_go//starlark",
    ],
)

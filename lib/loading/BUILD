load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "loading",
    srcs = [
        "bzlDepset.go",
        "bzlProvider.go",
        "bzlRule.go",
        "bzlStruct.go",
        "tasks.go",
    ],
    importpath = "github.com/HALtheWise/bagel/lib/loading",
    visibility = ["//visibility:public"],
    deps = [
        "//lib/core",
        "//lib/refs",
        "@net_starlark_go//starlark",
    ],
)

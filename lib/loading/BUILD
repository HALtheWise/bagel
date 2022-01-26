load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "loading",
    srcs = [
        "bzlDepset.go",
        "bzlProvider.go",
        "bzlStruct.go",
        "read_file.go",
        "rule.go",
    ],
    importpath = "github.com/HALtheWise/bagel/lib/loading",
    visibility = ["//visibility:public"],
    deps = [
        "//lib/core",
        "//lib/refs",
        "@net_starlark_go//starlark",
    ],
)

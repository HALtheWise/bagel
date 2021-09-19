load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "internal_lib",
    srcs = ["main.go"],
    importpath = "github.com/HALtheWise/balez/internal",
    visibility = ["//visibility:private"],
    deps = [
        "//internal/labels",
        "//internal/starlark_tasks",
        "//internal/task",
        "@com_github_jessevdk_go_flags//:go_default_library",
    ],
)

go_binary(
    name = "internal",
    embed = [":internal_lib"],
    visibility = ["//:__subpackages__"],
)

print({
    go_binary: 1,
    # "ac".startswith: 2,
})

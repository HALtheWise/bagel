load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "lib_lib",
    srcs = ["main.go"],
    importpath = "github.com/HALtheWise/bagel/lib",
    visibility = ["//visibility:private"],
    deps = [
        "//lib/starlark_tasks",
        "@com_github_halthewise_bagel//lib/task:go_default_library",
    ],
)

go_binary(
    name = "lib",
    embed = [":lib_lib"],
    visibility = ["//:__subpackages__"],
)

print({
    go_binary: 1,
    # "ac".startswith: 2,
})

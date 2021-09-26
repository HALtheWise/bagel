load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "task",
    srcs = [
        "context.go",
        "id.go",
        "task.go",
    ],
    importpath = "github.com/HALtheWise/bagel/internal/task",
    visibility = ["//:__subpackages__"],
    deps = [
        "//internal/cache",
        "//internal/cache/graph",
        "@com_zombiezen_go_capnproto2//:go_default_library",
    ],
)

go_test(
    name = "task_test",
    srcs = ["task_test.go"],
    deps = [":task"],
)

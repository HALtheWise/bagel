load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "task",
    srcs = [
        "context.go",
        "id.go",
        "task.go",
    ],
    importpath = "github.com/HALtheWise/bales/internal/task",
    visibility = ["//:__subpackages__"],
)

go_test(
    name = "task_test",
    srcs = ["task_test.go"],
    deps = [":task"],
)

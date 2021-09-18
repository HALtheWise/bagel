load("@bazel_gazelle//:def.bzl", "gazelle")

sh_binary(
    name = "go",
    srcs = ["go_wrapper.bash"],
    data = ["@go_sdk//:bin/go"],
    deps = ["@bazel_tools//tools/bash/runfiles"],
)

gazelle(name = "gazelle")

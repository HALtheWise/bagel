load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "983e495889ae0ca210f9a3f99647c57c9c2ea9f90bbaaca024956a17fff0e30b",
    strip_prefix = "rules_go-9f77676a87f23c7d5963bb9aef91f07398706809",
    # sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        # "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/archive/9f77676a87f23c7d5963bb9aef91f07398706809.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_local_sdk", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

# TODO(eric): Package this into a .tar and check it in
go_local_sdk(
    name = "go_sdk",
    path = "/home/eric/git/goroot",
)

go_register_toolchains()

# Gazelle

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_deps")

# gazelle:repository_macro deps.bzl%go_deps
go_deps()

gazelle_dependencies()

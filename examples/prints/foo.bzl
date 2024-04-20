def _foo_binary_impl(ctx):
    print("analyzing", ctx.label)
    return [DefaultInfo(files = depset([]))]

foo_binary = rule(
    implementation = _foo_binary_impl,
)

print("bzl file evaluation")

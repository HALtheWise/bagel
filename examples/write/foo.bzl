def _foo_binary_impl(ctx):
    out = ctx.actions.declare_file(ctx.label.name)
    print("_foo_binary_impl evaluation")
    ctx.actions.write(
        output = out,
        content = "Hello!\n",
        # is_executable = False,
    )
    return [DefaultInfo(files = depset([out]))]

foo_binary = rule(
    implementation = _foo_binary_impl,
)

print("bzl file evaluation")

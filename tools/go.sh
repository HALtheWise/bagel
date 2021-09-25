#!/usr/bin/env bash
exec bazel run --tool_tag=go --ui_event_filters=-info,-stdout,-stderr --noshow_progress -- @go_sdk//:bin/go "${@}"

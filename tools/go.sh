#!/usr/bin/env bash
# cd "$(dirname "$0")"
exec bazel run --tool_tag=go --ui_event_filters=-info,-stdout,-stderr --noshow_progress -- @rules_go//go "${@}"

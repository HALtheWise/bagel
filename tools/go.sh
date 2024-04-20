#!/usr/bin/env bash
# cd "$(dirname "$0")"
exec bazel --output_base build/tools_data run --tool_tag=go --ui_event_filters=-info,-stdout,-stderr --noshow_progress -- @rules_go//go "${@}"

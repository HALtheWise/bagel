#!/usr/bin/env bash
exec bazel run --tool_tag=gopackagesdriver --ui_event_filters=-info,-stdout,-stderr --noshow_progress -- @io_bazel_rules_go//go/tools/gopackagesdriver "${@}"

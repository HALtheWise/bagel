#!/usr/bin/env bash
exec bazel --output_base build/tools_data run --tool_tag=gopackagesdriver --ui_event_filters=-info,-stdout,-stderr --noshow_progress -- @rules_go//go/tools/gopackagesdriver "${@}"

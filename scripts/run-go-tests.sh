#!/bin/bash

# Exit script if you try to use an uninitialized variable.
set -o nounset

# Exit script if a statement returns a non-true return value.
set -o errexit

# Use the error status of the first failure, rather than that of the last item in a pipeline.
set -o pipefail

go get github.com/axw/gocov/gocov
go get github.com/matm/gocov-html

test_reports_dir=/tmp/test-reports/go
junit_dir=$test_reports_dir/junit
cover_dir=$test_reports_dir/cover
cover_file=$cover_dir/cover.out
html_dir=$test_reports_dir/html
mkdir -p $junit_dir
mkdir -p $cover_dir
mkdir -p $html_dir
gotestsum \
            --junitfile $junit_dir/junit.xml \
            --format standard-quiet \
            -- -coverprofile=$cover_file \
            --covermode=count \
            -tags integration \
            -short \
            ./...
gocov convert $cover_file | gocov-html > $html_dir/coverage.html

#!/bin/bash

curl -SsL https://downloads.gauge.org/stable | sh
gauge install html-report
gauge install screenshot
gauge install xml-report

# test/verify installation
if gauge --version; then
    GAUGE_VERSION=$(echo "$GAUGE" | grep -Po '(?<=Gauge version: )\d.\d.\d' )
    if $GAUGE_VERSION != "1.0.8"; then
        echo "New version of Gauge: $GAUGE_VERSION - you must update the Docker tag to get the build to pass."
        exit 1
    fi
    echo "$GAUGE_VERSION has been installed to $(which gauge)"
else
    echo "Something went wrong; Gauge could not be installed"
    exit 1
fi

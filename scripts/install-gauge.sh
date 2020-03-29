#!/bin/bash

curl -SsL https://downloads.gauge.org/stable | sh
gauge install html-report
gauge install screenshot
gauge install xml-report

# test/verify installation
if gauge --version; then
    echo "Gauge $(gauge --version) has been installed to $(which gauge)"
else
    echo "Something went wrong; Gauge could not be installed"
    exit 1
fi

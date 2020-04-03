#!/bin/bash

# To use this globally installed taiko, consumers should navigate to the directory
# where the local version would be installed, and run `npm link taiko`.
# This speeds up the build by avoiding having to install Taiko each time, and means
# there are fewer moving parts during the build.
# See:
# https://github.com/getgauge/taiko/issues/326#issuecomment-452584249
# https://github.com/getgauge/gauge-js/issues/174#issuecomment-435271044

npm config set prefix ~/.npm-global
npm install -g taiko --unsafe-perm --allow-root

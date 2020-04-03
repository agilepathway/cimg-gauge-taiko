#!/bin/bash
# shellcheck disable=SC2002

# taken from https://github.com/CircleCI-Public/browser-tools-orb/blob/master/src/commands/install-chrome.yml

if [[ $EUID == 0 ]]; then export SUDO=""; else export SUDO="sudo"; fi

# install chrome
if uname -a | grep Darwin; then
    brew update && \
    HOMEBREW_NO_AUTO_UPDATE=1 brew cask install google-chrome
    echo -e "#\!/bin/bash\n" > google-chrome
    perl -i -pe "s|#\\\|#|g" google-chrome
    echo -e "/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \"\$@\"" >> google-chrome
    $SUDO mv google-chrome /usr/local/bin
    $SUDO chmod +x /usr/local/bin/google-chrome
    # test/verify installation
    if google-chrome --version; then
    echo "$(google-chrome --version)has been installed in the /Applications directory"
    echo "A shortcut has also been created at $(which google-chrome)"
    exit 0
    else
    echo "Something went wrong; Google Chrome could not be installed"
    exit 1
    fi
elif cat /etc/issue | grep Alpine; then
    # https://github.com/Zenika/alpine-chrome/blob/master/Dockerfile
    echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories
    echo @edge http://nl.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories
    apk add --no-cache \
    chromium@edge \
    harfbuzz@edge \
    nss@edge \
    freetype@edge \
    ttf-freefont@edge
    rm -rf /var/cache/*
    mkdir /var/cache/apk
    # test/verify installation
    if chromium-browser --version; then
    echo "$(chromium-browser --version)has been installed to $(which chromium-browser)"
    exit 0
    else
    echo "Something went wrong; Chromium could not be installed"
    exit 1
    fi
elif command -v yum; then
    # download chrome
    curl --silent --show-error --location --fail --retry 3 \
    --output google-chrome-stable_current_x86_64.rpm \
    https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
    curl --silent --show-error --location --fail --retry 3 \
    --output liberation-fonts.rpm \
    http://mirror.centos.org/centos/7/os/x86_64/Packages/liberation-fonts-1.07.2-16.el7.noarch.rpm
    $SUDO yum localinstall -y liberation-fonts.rpm
    $SUDO yum localinstall -y google-chrome-stable_current_x86_64.rpm
    rm -rf google-chrome-stable_current_x86_64.rpm liberation-fonts.rpm
else
    # download chrome
    curl --silent --show-error --location --fail --retry 3 \
    --output google-chrome-stable_current_amd64.deb \
    https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
    $SUDO apt-get update && \
    $SUDO apt-get install -y \
    apt-utils \
    fonts-liberation \
    libappindicator3-1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libatk1.0-0 \
    libatspi2.0-0 \
    libcairo2 \
    libcups2 \
    libdbus-1-3 \
    libgdk-pixbuf2.0-0 \
    libglib2.0-0 \
    libgtk-3-0 \
    libnspr4 \
    libnss3 \
    libpango-1.0-0 \
    libpangocairo-1.0-0 \
    libxcomposite1 \
    libxcursor1 \
    libxi6 \
    libxrandr2 \
    libxrender1 \
    libxss1 \
    libxtst6 \
    lsb-release \
    xdg-utils
    # setup chrome installation
    $SUDO dpkg -i \
    google-chrome-stable_current_amd64.deb
    rm -rf google-chrome-stable_current_amd64.deb
fi
# test/verify installation
if google-chrome --version; then
    echo "$(google-chrome --version) has been installed to $(which google-chrome)"
else
    echo "Something went wrong; Google Chrome could not be installed"
    exit 1
fi

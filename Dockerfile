FROM cimg/go:1.14-node

LABEL maintainer="John Boyes <john@agilepathway.co.uk>"

COPY scripts /usr/local/bin/

USER root
RUN    chmod 755 /usr/local/bin/install-chrome.sh \
    && chmod 755 /usr/local/bin/install-gauge.sh \
    && chmod 755 /usr/local/bin/install-gauge-plugins.sh \
    && chmod 755 /usr/local/bin/install-taiko.sh \
    && install-chrome.sh \
    && install-gauge.sh
USER circleci
RUN    install-gauge-plugins.sh \
    && install-taiko.sh
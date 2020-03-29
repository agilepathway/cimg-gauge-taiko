FROM cimg/node:13.11

LABEL maintainer="John Boyes <john@agilepathway.co.uk>"

COPY scripts /usr/local/bin/

USER root
RUN    chmod 700 /usr/local/bin/install-chrome.sh \
    && chmod 700 /usr/local/bin/install-gauge.sh \
    && install-chrome.sh \
    && install-gauge.sh
USER circleci

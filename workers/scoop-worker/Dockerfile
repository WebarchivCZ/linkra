FROM node:22.20 AS base

# Add Tini (tiny init that will clean zombie processes)
ENV TINI_VERSION=v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
# -----------------------------------------------------------
# Args

ARG TARGET_UID=1000
ARG TARGET_GID=1000

# -----------------------------------------------------------
# Install app

# Where should playwright look for browsers
ENV PLAYWRIGHT_BROWSERS_PATH=/ms-playwright

RUN mkdir -p /ms-playwright && chmod 755 /ms-playwright

WORKDIR /app

COPY package.json ./

RUN npm install

RUN npx playwright install-deps chromium

COPY main.js ./

# Please mount appropriate config file in docker compose or on command line
#COPY config.json ./
# -----------------------------------------------------------
# Create non-root user

RUN groupadd -g ${TARGET_GID} scoop && useradd -u ${TARGET_UID} -g scoop -m -d /home/scoop scoop

RUN chown -R scoop:scoop /app /home/scoop /ms-playwright

RUN chmod 555 /ms-playwright

# -----------------------------------------------------------
# Switch user and start app

USER scoop

WORKDIR /app

ENTRYPOINT ["/tini", "--" ,"node", "main.js", "run"]
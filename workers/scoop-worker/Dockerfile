FROM node:22.20 AS base

# Add Tini (tiny init that will clean zombie processes)
ENV TINI_VERSION=v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
# -----------------------------------------------------------

WORKDIR /app

COPY package.json ./

RUN npm install

RUN npx playwright install-deps chromium

COPY main.js ./

# Please mount apropriate config file in docker compose or on command line
#COPY config.json ./

ENTRYPOINT ["/tini", "--" ,"node", "main.js", "run"]
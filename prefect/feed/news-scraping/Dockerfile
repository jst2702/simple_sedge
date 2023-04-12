FROM python:3.8-alpine

# Set Alpine version
ENV ALPINE_VERSION="3.16"

# Add Alpine Community Repository
RUN echo "@community https://dl-cdn.alpinelinux.org/alpine/v${ALPINE_VERSION}/community" >> /etc/apk/repositories

# Update apk index
RUN apk update

# # Install Firefox
# # RUN apk add firefox-esr
# RUN apk add chromium
# RUN apk add chromium-chromedriver

# Install dev dependencies
RUN apt-get update && apt-get install make
RUN apk add python3-dev \
    gcc \
    libc-dev \
    libffi-dev \
    g++

# Install Poetry
RUN pip install poetry

# Create a directory for the app
RUN mkdir /app

# Set the working directory
WORKDIR /app

# Copy files
COPY . .

# Disable virtualenv
RUN poetry config virtualenvs.create false

# Install dependencies
RUN poetry install

CMD [ "make deploy" ]
ENTRYPOINT [ "/app/start.sh" ]
FROM mysql:8.0

ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_DATABASE

# install migrate
RUN apt-get update && apt-get install -y curl
RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get install -y migrate

# run init script
COPY ./files/migration.sh.tpl /docker-entrypoint-initdb.d/migration.sh.tpl
RUN apt-get install -y gettext-base
RUN /usr/bin/envsubst '${MYSQL_USER},${MYSQL_PASSWORD},${MYSQL_DATABASE}' < '/docker-entrypoint-initdb.d/migration.sh.tpl' > '/docker-entrypoint-initdb.d/migration.sh'

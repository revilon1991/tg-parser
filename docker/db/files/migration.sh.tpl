#!/bin/bash

/usr/local/bin/migrate -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@unix(/var/run/mysqld/mysqld.sock)/${MYSQL_DATABASE}" -path /migration up

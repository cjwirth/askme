#! /bin/bash

DB_PATH="postgres://askme@localhost/askme_dev"
MIGRATIONS_PATH="."

if [[ $2 == "production" ]] ; then
    DB_PATH="postgres://askme@localhost/askme"
fi

if [[ $1 == "up" ]] ; then
    migrate -url $DB_PATH -path $MIGRATIONS_PATH up
elif [[ $1 == "down" ]] ; then
    migrate -url $DB_PATH -path $MIGRATIONS_PATH down
elif [[ $1 == "reset" ]] ; then
    migrate -url $DB_PATH -path $MIGRATIONS_PATH reset
elif ! [[ $1 =~ '^[0-9]+$' ]] ; then
    migrate -url $DB_PATH -path $MIGRATIONS_PATH goto $1
else 
    echo "Unknown command $1"
fi



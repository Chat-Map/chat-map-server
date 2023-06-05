#!/bin/bash

function init() {
  sleep 2 # wait for the database to be initialized
  cockroach sql --insecure --execute "CREATE DATABASE IF NOT EXISTS chatmap;"
}

init &
cockroach start-single-node --insecure

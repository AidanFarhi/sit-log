#!/bin/bash

sqlite3 sitlog.db < sql/tables.sql
sqlite3 sitlog.db < sql/seed.sql
sqlite3 sitlog.db < sql/validate.sql

#!/bin/bash
# Clear a development database.
sql-migrate down && sql-migrate status
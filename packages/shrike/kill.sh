#!/bin/bash
# Clear a development database and remove models.
sql-migrate down && rm -Rf models && sql-migrate status
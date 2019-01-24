#!/bin/bash
# Bootstrap a development database and create models.
sql-migrate up && sql-migrate status
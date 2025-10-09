#!/bin/bash
set -e

# Configure pg_hba.conf untuk allow replication dari Docker network
echo "host replication replicator 0.0.0.0/0 md5" >> "$PGDATA/pg_hba.conf"
echo "host all all 0.0.0.0/0 md5" >> "$PGDATA/pg_hba.conf"

echo "pg_hba.conf configured for replication"

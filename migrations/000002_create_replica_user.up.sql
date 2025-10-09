-- buat user untuk replication
CREATE ROLE replicator WITH REPLICATION LOGIN ENCRYPTED PASSWORD 'replsecret';

-- buat physical replication slot (opsional, bisa dibuat manual)
SELECT * FROM pg_create_physical_replication_slot('replica_slot');

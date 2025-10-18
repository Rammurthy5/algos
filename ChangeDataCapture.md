## what
a design pattern and set of technologies for detecting and capturing changes (inserts, updates, deletes) in a source data system — usually a database — and then delivering those changes in real time to other systems -- usually a data warehouse for analytical purposes. this is the modern ETL 

## ways to implement
1. Query-based polling (least efficient)
Periodically query the database for rows with updated_at > last_sync.
✅ Simple to set up
❌ High latency, can miss changes, puts load on DB
2. Database triggers
Use triggers to log changes to a separate table or push to a queue.
✅ Real-time
❌ Can impact transaction latency, harder to maintain at scale
3. Transaction log / WAL-based CDC (Log-based) 🏆 (most common modern approach)
Tap directly into the database’s write-ahead log (WAL) or equivalent (binlog in MySQL, redo log in Oracle, oplog in MongoDB, etc.).
✅ Real-time and low overhead
✅ No need to modify application code or schema
❌ Requires log parsing and careful checkpointing

## popular CDC tools
Debezium
 → open-source CDC platform for Postgres, MySQL, SQL Server, Oracle, MongoDB

Kafka Connect → CDC connectors for various sources

AWS DMS (Database Migration Service)

Oracle GoldenGate

StreamSets, Fivetran, Airbyte, Hevo, Confluent Cloud

## why can't we handle this in the app code and why need a CDC tool?
You’ll add retry queues

You’ll have to ensure ordering and deduplication

You’ll deal with failure recovery

You’ll handle schema evolution

You’ll build monitoring, alerting, etc.

👉 At that point, you’ve just created your own brittle, hard-to-maintain CDC system inside your app.
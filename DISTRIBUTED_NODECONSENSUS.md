### what 
general mechanism for agreeing on a single value (or sequence of values) among multiple distributed nodes, even in the presence of failures. Databases are just one of the most common applications.

there are diff algorithms available: Paxos, RAFT, zab etc.

## example
Google Spanner uses Paxos for replication; CockroachDB, Consul, and k8s etcd use Raft. ZooKeeper uses Zab

## why
SDEs don’t have to implement Paxos (almost nobody does in industry).
But understanding them makes you much better at diagnosing, designing, and scaling these systems.

## How Paxos work?
it works on principles with 2-phase commit and node majority. Majority is custom set by the users. 
Use Cases: Paxos is utilized in real-world systems, such as Google Chubby and Apache Zookeeper. These systems use Paxos for distributed locking by employing a virtual lock represented as a file in a file system.
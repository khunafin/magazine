
### Description

1. Client produce N packets 4Kb of size and store each packet in local memory.
2. Server read packets and addpen them to wal file, after send ack.
3. Client read ACK and delete packet from local memory.
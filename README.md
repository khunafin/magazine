
### Description

1. Client produces N packets 4Kb of size and stores each packet in local memory.
2. Server read packets and add them to wal file, after sending ack.
3. Client read ACK and delete the packet from local memory.
Your Go code calls Query(...)
Delegates the call to the registered driver.
The driver serializes the SQL string and parameters into the wire protocol(e.g., MySQL client-server binary protocol)

### Driver
Piece of software(go package) that knows how to:
- establish a network/local connection to a specific type of database
- translate Go's `database/sql` interface calls into the actual database [wire protocol](#wire-protocol) and queries
- marshal/unmarshal Go values into/from database-specific types

#### Wire Protocol
- refers to a way of getting data from point to point
- needed if more than one application has to interoperate
- the description of the data going across the network connection is the "wire-level protocol"

ToDo:
- maintain posts order across sessions




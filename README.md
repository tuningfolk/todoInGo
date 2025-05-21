### Driver
Piece of software(go package) that knows how to:
- establish a network/local connection to a specific type of database
- translate Go's `database/sql` interface calls into the actual database (wire protocol)[####Wire Protocol] and queries
- marshal/unmarshal Go values into/from database-specific types

#### Wire Protocol
- refers to a way of getting data from point to point
- needed if more than one application has to interoperate
- the description of the data going across the network connection is the "wire-level protocol"




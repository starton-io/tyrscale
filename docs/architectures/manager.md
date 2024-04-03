# Manager API

The manager API is reponsible for storing configurations and managing different part of the system :
- Gateway (oss)
- Distributed Monitor Service (enterprise)
- Distributed Recommender Service (enterprise)


The data configurations are stored as protobuf messags in the redis database (compress and divide the size of the payload by two). In the future, we want to use grpc protocol to communicate with the manager API for reduce the bandwith between internal tyrscale components.
Be careful, Redis database is not persistent database, so you need to be sure to backup your data before shutting. Starton recommend to have redis cluster to be able to recover data in case of failure.

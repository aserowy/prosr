# Prosr compiler

The compiler builds clients and hubs for SignalR, which are specified in prosr. The language definition for prosr1 can be found [here](./doc/prosr1_specification.md).

```sh
prosr -p path/to/SearchHub.prosr -l csharp
```

This will put together the hub definition of SearchHub for csharp classes. A namespace must be specified in the csharp template. The compilation is written to the directory of SearchHub.prosr.

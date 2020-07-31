# Introduction
prosr is a descriptive language for defining SignalR hubs and clients. The language is mainly based on [proto](https://developers.google.com/protocol-buffers) from GRPC and takes over a large part of the feature set.

prosr1 is based on [proto3](https://developers.google.com/protocol-buffers/docs/proto3) and takes over only rudimentary features in the first step to build a compiler for e.g. C# clients.

> #### Important
> If something is not listed in this documentation, although it is available as a feature in proto3, it is still not supported.

# Message types
As in [proto3](https://developers.google.com/protocol-buffers/docs/proto3#simple), messages form the basis of a class that can be used in a hub or hub client.

```
syntax = "prosr1";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```

# [Scalar Value Types](https://developers.google.com/protocol-buffers/docs/proto3#scalar)
|.prosr type | C# type |
|---|---|
| int32 | int |
| bool | bool |
| string | string |

# Using other message types
References to complex objects must be resolved [exactly as in proto3](https://developers.google.com/protocol-buffers/docs/proto3#other).

```
message SearchResponse {
  Result result = 1;
}

message Result {
  string url = 1;
  string title = 2;
  string snippets = 3;
}
```

# Defining hubs and hub clients
## Hub
A hub consists of at least one action and thus describes which interface the hub provides to the outside.

```
hub SearchHub {
  action Search(SearchRequest) returns (SearchResponse) to caller;

  returns (SearchRequest) to all;
}
```

## Actions and returns
An action consists of a method name and an optional transfer parameter. The transfer parameter must be of type message.

You can also define a return with a type for an action. Like the transfer parameter, the return type must be of type message. For the return, the target must also be described using "to".
The following values are currently supported as possible targets: caller, all

Unlike the proto definitions of grpc services, the concept of stream does not exist in SignalR. Each definition of an action and corresponding responses under hub can be executed multiple times. For example, multiple requests can be sent to the hub without each request getting its own response.

```
action ActionWithParameterAndReturnToCaller(OptionalParameter) returns (MandatoryParameter) to caller;

action ActionWithParameterWithoutReturn(OptionalParameter);

action ActionWithReturnToAll() returns (MandatoryParameter) to all;

action ActionWithParameterAndReturnToAll(OptionalParameter) returns (MandatoryParameter) to all;
```

Beyond actions, pure returns can also be defined in a hub. These are not available as end points in the hub, but can only be used for clients. Since the caller is unclear in such a scenario, only the target "all" can be used (as long as group is not yet available).

```
returns (MandatoryParameter) to all;
```
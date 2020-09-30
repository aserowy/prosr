# Prosr1 language guide

## Introduction

prosr is a descriptive language for defining SignalR hubs and clients. The language is mainly based on [proto3](https://developers.google.com/protocol-buffers/docs/proto3) and takes over only rudimentary features in the first step to build a compiler for e.g. C# clients.

> ### Important
>
> If something is not listed in this documentation, although it is available as a feature in proto3, it is still not supported.

## Syntax

Syntax specifies the given version used to describe hubs and clients.

```prosr1
syntax = "prosr1";
```

## Packages

Packages define blocks used for name resolution and e.g. namespaces in C# and other language templates. To define a package you can add multiple `package` specifier in your prosr1 definition. Each specifier defines a standalone namespace, thus sub namespacing is not supported.

```prosr1
package Search.SignalR;
```

## Messages

### Types

As in [proto3](https://developers.google.com/protocol-buffers/docs/proto3#simple), messages form the basis of a class that can be used in a hub or hub client.

```prosr1
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```

### Scalar Value Types

Message fields can have the following types. Like in [proto3](https://developers.google.com/protocol-buffers/docs/proto3#scalar) type conversersions are given in the following table.

|.prosr type | C# type |
|---|---|
| int32 | int |
| bool | bool |
| string | string |

### Enumeration with repeated

With repeated every type except maps can be enumerated. The listing retains the sequence.

```prosr1
message SearchResponse {
    repeated Result result = 1;
}
```

|.prosr | C# type |
|---|---|
| repeated | IEnumerable<> |

### Maps/Dictionaries

To create a map, use the keyword map and specify the types of the map. The first type must be int32 or string.

```prosr1
message SearchRequest {
    ...
    map<string, string> filter = 4;
}
```

|.prosr | C# type |
|---|---|
| map<,> | IDictionary<,> |

### Using other message types

References to complex objects must be resolved [exactly as in proto3](https://developers.google.com/protocol-buffers/docs/proto3#other).

```prosr1
message SearchResponse {
  repeated Result result = 1;
}

message Result {
  string url = 1;
  string title = 2;
  string snippets = 3;
}
```

## Defining hubs and hub clients

### Hub

A hub consists of at least one action and thus describes which interface the hub provides to the outside. Besides actions, calls can be defined independently of actions. These describe pure endpoints of the client.

```prosr1
hub SearchHub {
  action Search(SearchRequest) calls Search(SearchResponse) on caller;

  calls SearchCalled(SearchRequest) on all;
}
```

### Actions and calls

An action consists of a method name and an optional transfer parameter. The transfer parameter must be of type message.

You can also define a "calls" with an identifier and a type for an action. Like the transfer parameter, the return type must be of type message. For the call, the target must also be described using "on".
The following values are currently supported as possible targets: caller, all

Each definition of an action and corresponding responses under hub can be executed multiple times. For example, multiple requests can be sent to the hub without each request getting its own response.

Streams are not supported yet.

```prosr1
action ActionWithParameterAndReturnToCaller(OptionalParameter) calls CallIdent(OptionalParameter) on caller;

action ActionWithParameterWithoutReturn(OptionalParameter);

action ActionWithReturnToAll() calls CallIdent(OptionalParameter) on all;

action ActionWithParameterAndReturnToAll(OptionalParameter) calls CallIdent(OptionalParameter) on all;
```

Beyond actions, pure "calls" can also be defined in a hub. These are not available as end points in the hub, but can only be used for clients. Since the caller is unclear in such a scenario, only the target "all" can be used (as long as group is not yet available).

```prosr1
calls CallIdent(OptionalParameter) on all;
```

# Squzy - opensource monitoring, incident and alerting system

[![version](https://img.shields.io/github/v/release/squzy/squzy.svg)](https://github.com/squzy/squzy)
[![codecov](https://codecov.io/gh/squzy/squzy/branch/develop/graph/badge.svg)](https://codecov.io/gh/squzy/squzy)
[![GolangCI](https://golangci.com/badges/github.com/squzy/golangci-lint.svg)](https://golangci.com)
[![Join the chat at https://gitter.im/squzyio/community](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/squzyio/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## About

Squzy - is a high-performance open-source monitoring and alerting system written in Golang with [Bazel](https://bazel.build/) and love.

## Contains:

### [Squzy Monitoring](https://github.com/squzy/squzy/tree/develop/apps/squzy_monitoring)

High perfomance health check system

**System Health Checks Capabilities**
1) HTTP/HTTPS
2) TCP
3) GRPC - https://github.com/grpc/grpc/blob/master/doc/health-checking.md
4) SiteMap.xml - https://www.sitemaps.org/protocol.html
5) Value from http response by selectors(https://github.com/tidwall/gjson)

### [Squzy Agents](https://github.com/squzy/squzy/tree/develop/apps/agent_client)

Small application for get information from Host(server)

Which information you can get:
1. CPU load per each
2. Memory usage (used/free/total/shared)
3. Disk (used/free/total) per each disk
4. Net (bytes sent/get, package sent/get , err stat)

### [Squzy Agents Server](https://github.com/squzy/squzy/tree/develop/apps/squzy_agent_server)

Basic implementation for [agent server](https://github.com/squzy/squzy_proto/blob/master/proto/v1/squzy_agent_server.proto#L10)

### [Squzy Storage](https://github.com/squzy/squzy/tree/develop/apps/squzy_storage)

[Storage implementation](https://github.com/squzy/squzy_proto/blob/master/proto/v1/squzy_storage.proto#L13) will get all information from every product and store that

### [Squzy Dashboard](https://github.com/squzy/squzy-dashboard)

Squzy Web GUI - helps interact with squzy.

### [Squzy API](https://github.com/squzy/squzy/tree/develop/apps/squzy_api)

Squzy API server for works GUI + applications

# Authors
- [Iurii Panarin](https://github.com/PxyUp)
- [Nikita Kharitonov](https://github.com/DreamAndDrum)

# Want to help?
Want to file a bug, contribute some code, or improve documentation? Excellent!

Add merge request with description.

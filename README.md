# Thrift Binary Protocol Dump for Cassandra

This program can read the output of tcpdump and decode any cassandra thrift binary protocol requests into a human
readable format. This is primarily useful for visualizing inbound traffic in real time or verifying that an application
is sending data correctly (consistency levels etc.)

### Usage

 Generate some data on a cassandra node:

    tcpdump -enx -w cassandra-dump port 9160

Run it though the parser:

    go run main.go -f cassandra-dump

If no infile in supplied the program will attempt to read binary data from stdin:

    tcpdump -enx -w - port 9160 | go run main.go

If you want to see the discarded packets also you can supply the `-v` flag.

### Limitations

Currently thrift result messages are not parsed, this is not for any good reason other than the list of structures
is already ridiculously long.

Pointers are not de-referenced in output so some properties are not output.

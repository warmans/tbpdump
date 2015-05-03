# Cassandra Thrift Parser for tcpdump

Sometimes you want to see what thrift traffic is hitting a cassandra cluster but the clients are using the
binary format. This tool will take the output of a tcpdump session and parse any thift requests into a somewhat
readable output.

### Usage

 Generate some data on a cassandra node:

    tcpdump -enx -w cassandra-dump port 9160

Run it though the parser

    go run main.go -f cassandra-dump

If you want to see the discarded packets also you can supply the `-v` flag.

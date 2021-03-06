package main

import (
	"flag"
	"fmt"
	"github.com/akrennmair/gopcap"
	"github.com/araddon/cass/cassandra"
	"github.com/fatih/color"
	"github.com/pomack/thrift4go/lib/go/src/thrift"
	"os"
)

const VERSION string = "0.0.2"

func main() {

	inFile := flag.String("f", "", "the file to parse (generated by tcpdump e.g. tcpdump -enx -w cassandra-dump port 9160)")
	verbose := flag.Bool("v", false, "print information about skipped packets")
	version := flag.Bool("V", false, "print the version and exit")
	flag.Parse()

	if (*version == true) {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	var f *os.File
	if *inFile != "" {
		f, err := os.Open(*inFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	} else {
		f = os.Stdin
	}

	reader, err := pcap.NewReader(f)
	if err != nil {
		panic(err)
	}

	for pkt := reader.Next(); pkt != nil; pkt = reader.Next() {

		pkt.Decode()

		//thrify bits
		transport := thrift.NewTMemoryBuffer()
		transport.Write(pkt.Payload)
		protocol := thrift.NewTBinaryProtocol(transport, true, true)

		name, _, _, err := protocol.ReadMessageBegin()
		if err != nil {
			if *verbose == true {
				fmt.Println(pkt.String())
				fmt.Println(color.YellowString("not a thrift message"))
			}
			continue
		}

		tStruct := nameToTStruct(name)
		if tStruct != nil {
			tStruct.Read(protocol)
			fmt.Println(pkt.String())
			fmt.Println(color.GreenString(tStruct.String()))
		} else {
			fmt.Println(color.RedString("Unknown struct: ", name))
		}
	}
}

func nameToTStruct(name string) (tstruct ThriftStruct) {
	switch name {
	case "login":
		return cassandra.NewLoginArgs()
	case "set_keyspace":
		return cassandra.NewSetKeyspaceArgs()
	case "get":
		return cassandra.NewGetArgs()
	case "get_slice":
		return cassandra.NewGetSliceArgs()
	case "get_count":
		return cassandra.NewGetCountArgs()
	case "multiget_slice":
		return cassandra.NewMultigetSliceArgs()
	case "multiget_count":
		return cassandra.NewMultigetCountArgs()
	case "get_range_slices":
		return cassandra.NewGetRangeSlicesArgs()
	case "get_paged_slice":
		return cassandra.NewGetPagedSliceArgs()
	case "get_indexed_slices":
		return cassandra.NewGetIndexedSlicesArgs()
	case "insert":
		return cassandra.NewInsertArgs()
	case "add":
		return cassandra.NewAddArgs()
	case "remove":
		return cassandra.NewRemoveArgs()
	case "remove_counter":
		return cassandra.NewRemoveCounterArgs()
	case "batch_mutate":
		return cassandra.NewBatchMutateArgs()
	case "atomic_batch_mutate":
		return cassandra.NewAtomicBatchMutateArgs()
	case "truncate":
		return cassandra.NewTruncateArgs()
	case "describe_schema_versions":
		return cassandra.NewDescribeSchemaVersionsArgs()
	case "describe_keyspaces":
		return cassandra.NewDescribeKeyspacesArgs()
	case "describe_cluster_name":
		return cassandra.NewDescribeClusterNameArgs()
	case "describe_version":
		return cassandra.NewDescribeVersionArgs()
	case "describe_ring":
		return cassandra.NewDescribeRingArgs()
	case "describe_token_map":
		return cassandra.NewDescribeTokenMapArgs()
	case "describe_partitioner":
		return cassandra.NewDescribePartitionerArgs()
	case "describe_snitch":
		return cassandra.NewDescribeSnitchArgs()
	case "describe_keyspace":
		return cassandra.NewDescribeKeyspaceArgs()
	case "describe_splits":
		return cassandra.NewDescribeSplitsArgs()
	case "trace_next_query":
		return cassandra.NewTraceNextQueryArgs()
	case "describe_splits_ex":
		return cassandra.NewDescribeSplitsExArgs()
	case "system_add_column_family":
		return cassandra.NewSystemAddColumnFamilyArgs()
	case "system_drop_column_family":
		return cassandra.NewSystemDropColumnFamilyArgs()
	case "system_add_keyspace":
		return cassandra.NewSystemAddKeyspaceArgs()
	case "system_drop_keyspace":
		return cassandra.NewSystemDropKeyspaceArgs()
	case "system_update_keyspace":
		return cassandra.NewSystemUpdateKeyspaceArgs()
	case "system_update_column_family":
		return cassandra.NewSystemUpdateColumnFamilyArgs()
	case "execute_cql_query":
		return cassandra.NewExecuteCqlQueryArgs()
	case "execute_cql3_query":
		return cassandra.NewExecuteCql3QueryArgs()
	case "prepare_cql_query":
		return cassandra.NewPrepareCqlQueryArgs()
	case "prepare_cql3_query":
		return cassandra.NewPrepareCql3QueryArgs()
	case "execute_prepared_cql_query":
		return cassandra.NewExecutePreparedCqlQueryArgs()
	case "execute_prepared_cql3_query":
		return cassandra.NewExecutePreparedCql3QueryArgs()
	case "set_cql_version":
		return cassandra.NewSetCqlVersionArgs()
	default:
		return
	}
}

type ThriftStruct interface {
	Read(iprot thrift.TProtocol) (err thrift.TProtocolException)
	String() string
}

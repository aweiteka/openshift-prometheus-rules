# Testing Rules locally

The data directory has text files that simulate loading of metrics. This uses the promql package as a golang test. See [example files](https://github.com/prometheus/prometheus/tree/master/promql/testdata) that Prometheus uses. For a thorough explaination see the [measurementlab repo](https://github.com/m-lab/prometheus-support/tree/master/cmd/query_tester#test-format).

1. Get the test script

        go get github.com/aweiteka/openshift-prometheus-rules/test
1. Run test from this directory

        go test


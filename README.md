# openshift-prometheus-rules
A set of Prometheus recording and alerting rules for OpenShift

## Development Environment

1. Stand up OpenShift

         oc cluster up --public-hostname=127.0.0.1.nip.io
1. Login

         oc login -u system:admin
1. Run playbook

        ansible-playbook playbooks/adhoc/prometheus/prometheus.yml 
1. Run this script

        ./bootstrap-prometheus.sh
1. View prometheus here

        https://prometheus-openshift-metrics.127.0.0.1.nip.io/graph

### Issues

Cannot auth with this playbook since certs don't match subdomain. We open up prometheus as a workaround.

## Developing Alerts

An alert gets a human involved. It may be in the middle of the night. Humans can only react to one thing at a time. It must be used sparingly. A well-understood alert has these components:

- At least one metric that can measure it
- A query that can be used to view alerting state
- A way to reproduce alerting state
- An alert rule
- A documented debug path that answers, "What do I do when I see this alert?" This may include related metrics or logging queries, pre-defined graphs, and/or standard operating procedure (SOP) documentation.

Ideally, reproducing the alerting state can be automated to aid development and automated testing. This may involve using a script to push a metric to the Pushgateway.


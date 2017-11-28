# openshift-prometheus-rules
A set of Prometheus recording and alerting rules for OpenShift

## Developing Alerts

An alert involves a human. It may be in the middle of the night. Humans can only react to one thing at a time. It must be used sparingly. A well-understood alert has these components:

- At least one metric that can measure it
- A query that can be used to view alerting state
- A way to reproduce alerting state
- An alert rule
- A documented debug path that answers, "What do I do when I see this alert?" This may include related metrics or logging queries, pre-defined graphs, and/or standard operating procedure (SOP) documentation.

Ideally, reproducing the alerting state can be automated to aid development and automated testing. This may involve using a script to push a metric to the Pushgateway.

## Development Environment

See https://github.com/openshift/openshift-ansible/blob/master/README_CONTAINER_IMAGE.md

1. Stand up OpenShift

         oc cluster up --public-hostname=127.0.0.1 --routing-suffix=127.0.0.1.nip.io
1. Login

         oc login -u system:admin
1. Enable scheduling (not performed with openshift-ansible?)

        oc label node localhost region=infra app=prometheus
1. Set hostname (required?)

        sudo hostnamectl set-hostname localhost
1. Build custom image with iproute pkg (required?)

        docker build -t custom-openshift-ansible hack/.
1. Run installer

        docker run --rm -u root \
               --net=host \
               -v `pwd`/hack/inventory:/tmp/inventory:z \
               -v $HOME/.kube/config:/etc/origin/master/admin.kubeconfig:z \
               -e KUBECONFIG=/etc/origin/master/admin.kubeconfig \
               -e INVENTORY_FILE=/tmp/inventory \
               -e PLAYBOOK_FILE=playbooks/byo/openshift-cluster/openshift-prometheus.yml \
               -e OPTS="-v" -it   custom-openshift-ansible
1. Add developer user to openshift-metrics project

        oc policy add-role-to-user admin developer -n openshift-metrics
1. Deploy node exporter template

        curl https://raw.githubusercontent.com/openshift/origin/master/examples/prometheus/node-exporter.yaml | oc create -f - -n kube-system
1. Add hostaccess SCC so node exporter can get system metrics

        oc adm policy add-scc-to-user -z prometheus-node-exporter -n kube-system hostaccess
1. View prometheus service at https://prometheus-openshift-metrics.127.0.0.1.nip.io/graph

### Issues

1. Cannot auth with this playbook since certs don't match subdomain. We open up prometheus auth as a workaround.

        oc policy add-role-to-user view system:anonymous -n openshift-metrics
1. If rebuilding sometimes the openshift configuration needs to be wiped, then redeploy from step 1.

        sudo rm -rf /var/lib/origin/openshift.local.config


# Developing Rules

## Syntax

- use the yaml format
- name the files `<component>.rules`
- use descriptive and unique names, CamelCase
- include summary and description that is helpful to the operator

## Tools

Rules cannot be dynamically loaded into OpenShift. For now we redeploy prometheus--downtime is ~10s. NOTE: the deployment will go into **crash loop backoff** on a rule file syntax error. Therefore, a syntax check is highly recommended before redeploying edited rules files. The `promtool` maybe downloaded from the [Prometheus web site](https://prometheus.io/download/).

## Loading rules files from this repo

With this deployment method all files in the rules directory are mounted into the pod as a configmap.

1. Create a configmap of the rules directory

        oc create configmap base-rules --from-file=rules/ -n openshift-metrics
1. Edit the prometheus configuration.

	oc edit cm prometheus -n openshift-metrics
1. Add the name of the rules directory

	  prometheus.yml: |
	    rule_files:
	      - 'prometheus.rules'
	      - 'rules/*.rules'
1. Attach the configmap to the prometheus statefulset as a volume

        oc volume statefulset/prometheus --add \
           --configmap-name=base-rules --name=base-rules -t configmap \
           --mount-path=/etc/prometheus/rules -n openshift-metrics
1. Reload the rules using the procedure below.

## Updating rules files

1. Edit or add a local rules file
1. Validate the rules directory

        promtool check rules rules/*
1. Update the configmap

        oc delete cm base-rules -n openshift-metrics
        oc create configmap base-rules --from-file=rules/ -n openshift-metrics
1. Reload rules using the procedure below

## Reloading Alerts

NOTE: It can take over 60 seconds for changes to a configmap to appear in a pod. It is more reliable to simply delete the pod so it creates a new one with the new configmap. This has the cost of ~10s downtime but ensures you've got the updated config.

1. Delete pod so it restarts with new configuration

        oc delete $(oc get pods -o name --selector='app=prometheus' -n openshift-metrics) -n openshift-metrics
1. The pod immediately recreates. NOTE: there are 5 containers in the pod. To watch

        oc get pods -w -n openshift-metrics

## Debugging

Working with prometheus is unique.

- Prometheus runs as a stateful set, not a deploymentconfig.

        oc describe statefulset prometheus -n openshift-metrics
- There are 5 containers in the pod.

        $ oc get pod prometheus-0 -n openshift-metrics
        NAME           READY     STATUS    RESTARTS   AGE
        prometheus-0   5/5       Running   0          4m

  - prom-proxy
  - prometheus
  - alerts-proxy
  - alert-buffer
  - alertmanager

- To enter the pod you need to specify the container you want. For example

        oc rsh -c prometheus prometheus-0 -n openshift-metrics


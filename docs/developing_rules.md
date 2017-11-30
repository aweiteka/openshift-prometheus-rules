# Rules Development Process

## Tools

Rules cannot be dynamically loaded into OpenShift. A syntax check is highly recommended before uploading edited rules files. The `promtool` maybe downloaded from the [Prometheus web site](https://prometheus.io/download/).

## Bootstrapping rules files from this repo or adding new rules files

1. Create a new rules file in rules directory
1. Validate the rules files in the rules directory

        promtool check rules rules/*
1. Create a configmap of the file

        oc create configmap base-rules --from-file=rules/ -n openshift-metrics
1. Edit the prometheus configuration.

	oc edit cm prometheus -n openshift-metrics
1. Add the name of the new alerts file

	  prometheus.yml: |
	    rule_files:
	      - 'prometheus.rules'
	      - 'rules/os.rules'
1. Attach the configmap to the prometheus statefulset as a volume

        oc volume statefulset/prometheus --add \
           --configmap-name=base-rules --name=base-rules -t configmap \
           --mount-path=/etc/prometheus/rules -n openshift-metrics
1. Reload the alerts using the procedure below.

## Editing or Updating Existing Alerts files

1. Edit a local rules file
1. Validate the file

        promtool check rules rules/*
1. Update the configmap

        oc delete cm base-rules -n openshift-metrics
        oc create configmap base-rules --from-file=rules/ -n openshift-metrics
1. Reload alerts using the procedure below

## Reloading Alerts

NOTE: It can take over 60 seconds for changes to a configmap to appear in a pod. It is more reliable to simply delete the pod so it creates a new one with the new configmap.

1. Delete pod so it restarts with new configuration

        oc delete $(oc get pods -o name --selector='app=prometheus' -n openshift-metrics) -n openshift-metrics


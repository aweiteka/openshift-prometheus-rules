# Rules Development Process

## Adding Alerts files

1. Edit the prometheus configuration.

	oc edit cm prometheus
1. Add the name of the new alerts file

	  prometheus.yml: |
	    rule_files:
	      - 'prometheus.rules'
	      - 'rules/system.rules'
1. Create a configmap of the file

        oc create configmap custom-rules --from-file=rules/
1. Attach the configmap to the prometheus statefulset as a volume

        oc volume statefulset/prometheus --add \
           --configmap-name=custom-rules --name=custom-rules -t configmap \
           --mount-path=/etc/prometheus/rules
1. Reload the alerts using the procedure below.

## Editing or Updating Existing Alerts files

1. Edit a local rules file
1. Update the configmap

        oc delete cm custom-rules
        oc create configmap custom-rules --from-file=rules/
1. Reload alerts using the procedure below

## Reloading Alerts

NOTE: It can take over 60 seconds for changes to a configmap to appear in a pod. It is more reliable to simply delete the pod so it creates a new one with the new configmap.

1. Delete pod so it restarts with new configuration

        oc delete $(oc get pods -o name --selector='app=prometheus')


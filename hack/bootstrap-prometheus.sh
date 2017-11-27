#!/bin/bash

if ! oc status ; then
  echo "First authenticate the oc client as cluster administrator"
  exit 1
fi

oc policy add-role-to-user admin developer -n openshift-metrics
oc label node localhost app=prometheus region=infra
oc project openshift-metrics
oc policy add-role-to-user view system:anonymous -n openshift-metrics

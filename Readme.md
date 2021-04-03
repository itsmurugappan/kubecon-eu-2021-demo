# Health Device and Apps Date integration with Cloud Events and Knative Eventing

This repo was developed for the demo at kubecon Europe 2021

### Pre Req

To run this app you need the below

* Knative Serving and Eventing
* KO

### Install Mysql and Grafana

```
kubectl create ns kubecon

kubectl apply -f config/mysql.yaml

# see the steps in config/mysql.md to configure mysql

kubectl apply -f config/grafana.yaml
```

Create the datasource in grafana and connect to mysql

Import the dashboard from config/grafana-dashboard.yaml

### Deploy Knative Artifacts

```
export KO_DOCKER_REPO=<your git repo>

kubectl apply -f config/rb.yaml

ko apply -RBf serving/

ko apply -RBf eventing/
```

### Triggering the flow

```
export KN_DOMAIN = <your kn domain>

curl "http://health-data-ingest-kubecon.${KN_DOMAIN}/?labels=app=health-data-ingest&DEVICE_TYPE=STRAVA&history=0"

curl "http://health-data-ingest-kubecon.${KN_DOMAIN}/?labels=app=health-data-ingest&ACT_DATE=2021-04-03&DEVICE_TYPE=FITBIT&history=0"

curl "http://health-data-ingest-kubecon.${KN_DOMAIN}/?labels=app=health-data-ingest&DEVICE_TYPE=RUNKEEPER&history=0"
```

### Dashboard

![](./images/dashboard.jpg)



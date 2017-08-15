###############################################################################
# Google Cloud Platform
###############################################################################
gcp_init:
	#Installing: App Engine Go
	gcloud components install app-engine-go
	${HOME}/work/google-cloud-sdk/bin/dev_appserver.py

gcp_sdk_reset:
	gcloud init


###############################################################################
# Projects
###############################################################################
proj_init:
	gcloud config set project go-goa
	gcloud config set compute/zone europe-west1-c

proj_check:
	gcloud config list
	#https://cloud.google.com/sdk/gcloud/reference/config/set


###############################################################################
# Container Engine
###############################################################################
#https://console.cloud.google.com/kubernetes/list?project=go-goa&hl=ja

container_engine_init:
	#https://cloud.google.com/container-engine/docs/tutorials/persistent-disk?hl=ja
	gcloud components update kubectl


create_cluster:
	gcloud container clusters create go-goa --num-nodes=3

	#NAME    ZONE            MASTER_VERSION  MASTER_IP     MACHINE_TYPE   NODE_VERSION  NUM_NODES  STATUS
	#go-goa  europe-west1-c  1.6.7           35.195.86.72  n1-standard-1  1.6.7         3          RUNNING


###############################################################################
# Create
###############################################################################
create_disk:
	gcloud compute disks create --size 1GB mysql-disk
	#gcloud compute disks create --size 1GB web-disk

	#NAME        ZONE            SIZE_GB  TYPE         STATUS
	#mysql-disk  europe-west1-c  1        pd-standard  READY

create_key:
	kubectl create secret generic mysql --from-literal=rootpass=root
	kubectl create secret generic mysql2 --from-literal=pass=12345678

deploy_mysql:
	kubectl create -f mysql.yaml
	kubectl create -f mysql-service.yaml

deploy_web:
	kubectl create -f web.yaml
	kubectl create -f web-service.yaml

deploy_all: deploy_mysql deploy_web


###############################################################################
# Delete
###############################################################################
delete_mysql:
	kubectl delete pod -l app=mysql
	kubectl delete service mysql

	#OR

	kubectl delete -f mysql.yaml
	kubectl delete -f mysql-service.yaml


delete_web:
	kubectl delete pod -l app=web
	kubectl delete service web

	#OR

	kubectl delete -f web.yaml
	kubectl delete -f web-service.yaml


delete_all: delete_web delete_mysql

cleanup: delete_all
	gcloud compute forwarding-rules list
	gcloud container clusters delete go-goa
	gcloud compute disks delete mysql-disk


###############################################################################
# Update # execute after deleting
###############################################################################
update_mysql:
	kubectl apply -f mysql.yaml

update_web:
	kubectl apply -f web.yaml

update_all: update_mysql update_web


###############################################################################
# Reset
###############################################################################
resetm:
	kubectl delete pod -l app=mysql
	kubectl apply -f mysql.yaml


###############################################################################
# Check
###############################################################################
check_mysql:
	kubectl get pod -l app=mysql
	kubectl get service mysql

	#kubectl describe pod mysql-3804075166-8fzbp
	#kubectl logs mysql-3804075166-8fzbp

	#NAME      CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
	#mysql     10.59.245.61   <none>        3306/TCP   25s

check_web:
	kubectl get pod -l app=web
	kubectl get svc -l app=web

	#NAME      CLUSTER-IP      EXTERNAL-IP      PORT(S)        AGE
	#web       10.59.253.208   35.195.138.190   80:31205/TCP   2m

	# http://35.195.135.47/

#getm:
#	$(eval ID := $(shell kubectl get pod -l app=mysql | tail -n 1 | awk '{print $1}'))
#	@echo $(ID)
#	#$(ID)

mys:
	kubectl get pod -l app=mysql

logm:
	$(SHELL) -c './mysqlog.sh'


###############################################################################
# New
###############################################################################
create_new:
	gcloud container clusters create go-goa --num-nodes=3
	#gcloud compute disks create --size 1GB mysql-disk

	kubectl create -f pv.yaml

	kubectl create secret generic mysql-secret --from-literal=mysql-root-password=root --from-literal=mysql-user=hiromaily --from-literal=mysql-password=12345678
	kubectl create configmap db --from-literal=mysql-database=hiromaily

	kubectl create -f mysql.yaml

	kubectl get pod

create_new_after_running:
	kubectl create -f mysql-service.yaml

	kubectl create -f web.yaml
	kubectl create -f web-service.yaml

	kubectl get svc -l app=web


check_new:
	#kubectl get pod

	#kubectl get pod -l mysql-server
	#kubectl get service mysql-server
	#kubectl get svc -l app=web

	kubectl logs mysql-3449492693-6dg8w
	kubectl describe pod mysql-3449492693-6dg8w

	kubectl logs web-899485354-b3v9z

delete_new:
	kubectl delete -f web.yaml
	kubectl delete -f mysql.yaml
	kubectl delete -f pv.yaml

	#gcloud compute disks delete mysql-disk

	gcloud container clusters delete go-goa


mysql:
	kubectl port-forward mysql-4197980765-57rmp 3307:3306
	mysql -u root -P 3307



mysql_fix:
    #brew services start mysql

	#mysql -u root -D hiromaily -p
	#ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/tmp/mysql.sock'

	#mysqld stop

	#sudo chown -R _mysql:_mysql /usr/local/var/mysql
	#sudo chmod -R o+rwx /usr/local/var/mysql

    #mysql.server start
    #sudo mysql.server start




get_all_pod:
	kubectl get pod
	kubectl get rc
	kubectl get svc

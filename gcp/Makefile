###############################################################################
# Google Cloud Platform
###############################################################################
gcp_init:
	#Installing: App Engine Go
	gcloud components install app-engine-go
	${HOME}/work/google-cloud-sdk/bin/dev_appserver.py

gcp_sdk_reset:
	gcloud init

container_engine_init:
	#https://cloud.google.com/container-engine/docs/tutorials/persistent-disk?hl=ja
	gcloud components update kubectl


###############################################################################
# minikube
#  https://kubernetes.io/docs/tutorials/stateless-application/hello-minikube/
###############################################################################
minikube_init:
	brew install docker-machine-driver-xhyve

	# docker-machine-driver-xhyve need root owner and uid
	sudo chown root:wheel $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve
	sudo chmod u+s $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve

	brew cask install minikube

mini:
	minikube start --vm-driver=xhyve
	#minikube start

	#kubectl cluster-info

	#kubectl proxy

ministop:
	minikube stop


#minikube service mysql-server --url
# => http://192.168.99.100:32741
#    $ mysql -uroot -p -h 192.168.99.100 --port 32741

#minikube service web --url
# => http://192.168.99.100:31983


#minikube Commands:
#  addons           Modify minikube's kubernetes addons
#  completion       Outputs minikube shell completion for the given shell (bash)
#  config           Modify minikube config
#  dashboard        Opens/displays the kubernetes dashboard URL for your local cluster
#  delete           Deletes a local kubernetes cluster
#  docker-env       Sets up docker env variables; similar to '$(docker-machine env)'
#  get-k8s-versions Gets the list of available kubernetes versions available for minikube
#  ip               Retrieves the IP address of the running cluster
#  logs             Gets the logs of the running localkube instance, used for debugging minikube, not user code
#  mount            Mounts the specified directory into minikube
#  profile          Profile sets the current minikube profile
#  service          Gets the kubernetes URL(s) for the specified service in your local cluster
#  ssh              Log into or run a command on a machine with SSH; similar to 'docker-machine ssh'
#  ssh-key          Retrieve the ssh identity key path of the specified cluster
#  start            Starts a local kubernetes cluster
#  status           Gets the status of a local kubernetes cluster
#  stop             Stops a running local kubernetes cluster
#  update-context   Verify the IP address of the running cluster in kubeconfig.
#  version          Print the version of minikube


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
	kubectl create secret generic mysql-secret --from-literal=mysql-root-password=root --from-literal=mysql-user=hiromaily --from-literal=mysql-password=12345678
	kubectl create configmap db --from-literal=mysql-database=hiromaily

create_map:
	#kubectl create configmap my-config --from-file=key1=/path/to/bar/file1.txt --from-file=key2=/path/to/bar/file2.txt
	kubectl create configmap mysql-initdb-config \
	--from-file=user.sql=/Users/hy/work/go/src/github.com/hiromaily/go-goa/docker/mysql/init.d/user.sql \
	--from-file=dump.sql=/Users/hy/work/go/src/github.com/hiromaily/go-goa/docker/mysql/init.d/dump.sql

deploy_mysql:
	kubectl create -f mysql.yaml
	kubectl create -f mysql-service.yaml

deploy_web:
	kubectl create -f web.yaml
	kubectl create -f web-service.yaml


deploy:
	#gcloud container clusters create go-goa --num-nodes=3
	#gcloud compute disks create --size 1GB mysql-disk

	kubectl create -f pv.yaml

	kubectl create secret generic mysql-secret --from-literal=mysql-root-password=root --from-literal=mysql-user=hiromaily --from-literal=mysql-password=12345678
	kubectl create configmap db --from-literal=mysql-database=hiromaily
	kubectl create configmap mysql-initdb-config \
	--from-file=user.sql=/Users/hy/work/go/src/github.com/hiromaily/go-goa/docker/mysql/init.d/user.sql \
	--from-file=dump.sql=/Users/hy/work/go/src/github.com/hiromaily/go-goa/docker/mysql/init.d/dump.sql

	kubectl create -f mysql.yaml

	kubectl get pod
	#kubectl get pod -l app=mysql

deploy2:
	kubectl create -f mysql-service.yaml

	kubectl create -f web.yaml
	kubectl create -f web-service.yaml

	kubectl get svc -l app=web
	#kubectl get pod -l app=web

	#minikube service web --url

restart:
	kubectl delete -f web-service.yaml
	kubectl delete -f web.yaml
	kubectl delete -f mysql-service.yaml
	kubectl delete -f mysql.yaml
	kubectl delete -f pv.yaml


###############################################################################
# Delete
###############################################################################
delete_mysql:
	kubectl delete -f mysql-service.yaml
	kubectl delete -f mysql.yaml
	#kubectl delete pod -l app=mysql
	#kubectl delete service mysql

delete_web:
	kubectl delete -f web-service.yaml
	kubectl delete -f web.yaml
	#kubectl delete pod -l app=web
	#kubectl delete service web

cleanup:
	#gcloud compute forwarding-rules list
	kubectl delete -f pv.yaml

	gcloud container clusters delete go-goa
	#gcloud compute disks delete mysql-disk


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
check:
	kubectl get pod
	#kubectl get rc
	kubectl get svc
	kubectl get events

	kubectl config view

check_mysql:
	kubectl get pod -l app=mysql
	kubectl get service mysql

	#kubectl describe pod mysql-3804075166-8fzbp
	#kubectl logs mysql-3804075166-8fzbp

check_web:
	kubectl get pod -l app=web
	kubectl get svc -l app=web

logm:
	$(SHELL) -c './mysqlog.sh'

logw:
	$(SHELL) -c './weblog.sh'

#getid:
#	$(eval ID := $(shell kubectl get pod -l app=mysql | tail -n 1 | awk '{print $1}'))
#	@echo $(ID)
#	#$(ID)


###############################################################################
# MySQL
###############################################################################
local_settings:
	#brew services start mysql

	#mysql -u root -D hiromaily -p
	#ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/tmp/mysql.sock'

	#mysqld stop

	#sudo chown -R _mysql:_mysql /usr/local/var/mysql
	#sudo chmod -R o+rwx /usr/local/var/mysql

	#mysql.server start
	#sudo mysql.server start

mysql:
	kubectl port-forward mysql-4197980765-57rmp 3307:3306
	mysql -u root -P 3307


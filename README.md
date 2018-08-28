Hello World

The presentation is [here](https://docs.google.com/presentation/d/1LyRuNm0cv-yk0UWnYsxALLSq2LXFG7zFnXX9-uT9cRI/edit?usp=sharing)


# demo note

## dev side

```shell
cd ./dev
# build docker image
./docker/builder/build
# start app inside docker
./docker/app/app start
# test the app
curl localhost:8080/counter

# push image to registry
# NOTE: 10.12.34.101 must accessible by adding routing to network namespace (**1)
#       and registry must be running first (**2)
#       see below
docker tag counter-app:latest 10.12.34.101:8000/counter-app:latest
docker push 10.12.34.101:8000/counter-app:latest
```

## ops side
```shell
cd ./ops

# download all files
./download-files

# add routing to the network (**1)
sudo ip route add 10.12.34.0/24 via `./docker/net ip`

# open new terminal: run vm 1
./docker/kvm run_vm 1
# open new terminal: run vm 2
./docker/kvm run_vm 2
# open new terminal: run vm 3
./docker/kvm run_vm 3
# open new terminal: run vm 4
./docker/kvm run_vm 4
# open new terminal: run vm 5
./docker/kvm run_vm 5

# ansible
cd ./ansible-playbook
# run bootstrap playbook
./run-playbook playbooks/coreos_bootstrap/main.yml
# deploy registry container (**2)
./run-playbook playbooks/registry/main.yml
# deploy database container
./run-playbook playbooks/database/main.yml
# deploy counter_app
./run-playbook playbooks/counter_app/main.yml
# deploy lb
./run-playbook playbooks/lb/main.yml

# test the app
curl 10.12.34.101/counter
```

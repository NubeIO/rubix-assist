## cmd cli

these docs are for using the linux or windows command line

- cd into the cmd dir
- once in the main dir there is access to different protocols, or you can for example add/remove a host

```
cd cmd
```

## apps

install a rubix app

- download from asset from GitHub
- generate the systemd file
- install as a linux service

```
go build main.go && sudo ./main apps --owner=NubeIO  --repo=rubix-bios --dest=/data --target=rubix-bios-app  --arch=amd64 --tag=v1.5.2 --token=<token>

## main app

### hosts

add a new a host

````

go run main.go hosts --new=true --name=test --ip=192.178.12.1

````

update a host ip (will up by the host name)

````

go run main.go hosts --update-ip=true --name=RC --ip=192.168.15.11

````

update a host ssh port (will up by the host name)

````

go run main.go hosts --update-port=true --name=RC --port=2022

````

delete all hosts

````

go run main.go hosts --drop=true

````

## flow network

### ip scan

look for all devices on a local network with ports like 1616, 1313, 502, 22

````

go run main.go flow --scan=true --ip-range=192.168.15.1-12

````

## modbus

### example

````

(cd modbus && go run main.go reg --ip=192.168.15.202 --type=writeCoil --value=0)

````

### docs

https://github.com/NubeDev/bacnet#examples

## bacnet

### example

````

(cd bacnet && go run main.go whois --interface=wlp3s0)

````

### docs

https://github.com/NubeDev/bacnet#examples

## rubix-io

### example

````

(cd rubixio && go run main.go read --ip=192.168.15.10 --port=5001)

````

### docs

## systemctl

### example

`this will run as sudo`
````

(cd systemctl && go build ctl.go && sudo ./ctl service --status=true --service=myservice)

````

### docs

see https://github.com/NubeIO/nubeio-rubix-lib-rest-go/tree/master/pkg/nube/rubixio/cmd#rubix-io-rest-client

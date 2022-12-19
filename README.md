# SSHAWS

`sshaws` is a helper to simplify the *AWS instances ssh access*.

> It will show a list of instances depending on some specified tags and help you to ssh on it.

## Getting Started

```bash
usage: sshaws [--silent true] [--region xx_yy_j] [-l user ] [-ssh true] [-k true] [instance_name]
```

### Examples

```bash
# List all the instances from eu-west-1
> sshaws

Name: *   Region: eu-west-1
---------------------------------------------------------

ID           NAME                IP              INSTANCE ID            SIZE          LAUNCHTIME
 0      front-prod-1        172.16.31.12      i-0978df6a92b39d434      t2.nano    2020-12-18 08:08:08
 1      front-prod-2        172.16.24.226     i-0b3914e83516392378     t2.nano    2020-12-18 08:08:08
 2      front-staging       172.16.39.121     i-0b3914e89237829ad1     t2.micro   2020-12-18 08:08:08
 3      back-prod           172.16.33.21      i-0b3914421237829ad1     t3.large   2020-12-18 08:08:08
 4      back-staging        172.16.19.93      i-0b391351e237829ad1     t3.micro   2020-12-18 08:08:08

Which one do you want to ssh in?
# WAITING INPUT FROM USER
3
>> Starting a new ssh session to 172.16.33.21
[...] Stablishing SSH connection with the desired server
```

```bash
# List all the instances with front on name, from eu-west-1
> sshaws front

Name: front   Region: eu-west-1
---------------------------------------------------------

ID           NAME                IP              INSTANCE ID            SIZE          LAUNCHTIME
 0      front-prod-1        172.16.31.12      i-0978df6a92b39d434      t2.nano    2020-12-18 08:08:08
 1      front-prod-2        172.16.24.226     i-0b3914e83516392378     t2.nano    2020-12-18 08:08:08
 2      front-staging       172.16.39.121     i-0b3914e89237829ad1     t2.micro   2020-12-18 08:08:08

Which one do you want to ssh in?
# WAITING INPUT FROM USER
0
>> Starting a new ssh session to 172.16.31.12
[...] Stablishing SSH connection with the desired server
```

### Usage in depth

Download the binary in this repository and execute it, or compile by yourself following the next steps.

```bash
sshaws --help
Usage of sshaws:
  -k	Push temporal public key to instance [short mode]
  -l string
    	SSH login name (default "ubuntu")
  -n string
    	Instance Name [short mode] (default "*")
  -name string
    	Instance Name (default "*")
  -region string
    	AWS Region (default "eu-west-1")
  -silent
    	Show only IP
  -ssh
    	Use SSH instead of SSM [short mode]
  -version
    	Display app version
```

### Upgrade sshaws version

To generate a new version of the sshaws tool, you can make your changes and add a new tag.

Please follow the vX format (e.g. v1, v2, v3, etc).

To add a tag and push it, you can execute

```bash

git tag vX && git push --tags

```

The download URLs for the binary will be the same, as they point to latest.

### Diagram

![The flow to connect EC2:](images/ssm-sessionmanager.png#center)

### Prerequisites

#### Client requirements

There are some basic prerequisites :

- [**AWS CLI**](https://docs.aws.amazon.com/es_es/cli/latest/userguide/install-cliv2.html) tool
- [**Session Manager Plugin**](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html#plugin-version-history). If you want to connect using ssm.
- **SSH**: You will need ssh, and, if necessary, configure the ssh connection (The configuration in `~/.ssh/configuration` will be applied to sshaws).
- **[OPTIONAL] Golang**: If you want to compile `sshaws` you will need the `golang` executable.

Configure your **credentials** ~/.aws/ with the correct keys.
- ~/.aws/config
```bash
[default]
region = eu-west-1
```

- ~/.aws/credentials
```bash
[default]
aws_access_key_id = AKIAXXXXXXXXXXXXX
aws_secret_access_key = xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

#### Server requirements

- [OPTIONAL] `ec2-instance-connect`: If you want to copy the keys to the server, you will need to install `ec2-instance-connect` package in the server

### Installing

Currently the CI build on github generates the binary files for Linux, OSX and Windows.

### From binary

#### For Linux

```bash
# Download latest version from https://github.com/holaluz/sshaws/releases/latest
sudo wget -O /usr/local/bin/sshaws https://github.com/holaluz/sshaws/releases/latest/download/sshaws

# Give it execution permissions
sudo chmod +x /usr/local/bin/sshaws
```

#### For MAC

```bash
# Download latest version from https://github.com/holaluz/sshaws/releases/latest
sudo wget -O /usr/local/bin/sshaws https://github.com/holaluz/sshaws/releases/latest/download/sshaws.mac

# Give it execution permissions
sudo chmod +x /usr/local/bin/sshaws
```

#### For Windows

Download the windows executable from https://github.com/holaluz/sshaws/releases/latest/download/sshaws.win

Right click on the file, ensure it is not Locked and that it has execute permissions.


### From source code

After downloading this repository:

```bash
# Enter repository folder
cd sshaws

# Build the binary
go build -o sshaws cmd/sshaws/main.go

# Move to a desired folder
sudo mv sshaws /usr/local/bin/sshaws
```

## Running the tests

```bash
# Run the test recursively
go test ./...
```

## Built With

- [Golang](https://golang.org/)

## Contributing

Please create a fork and create a PR!
It would be awesome to have new proposals!

## Versioning

Travis is configured to automatically build and update the version. (You can change the current version on [version.go](pkg/cmd/version.go)).

## Authors

- [uritau](https://github.com/uritau), [Oriol Tauleria](mailto:oriol.tauleria@gmail.com)
- [xaf](https://github.com/xafardero)
- [jiba21](https://github.com/jiba21)
- [viscat](https://github.com/viscat)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

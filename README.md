# SSHAWS

`sshaws` is a helper to simplify the *AWS instances ssh access*.

> It will show a list of instances depending on some specified tags and help you to ssh on it.


## Getting Started
> Limitations: sshaws will only display instances with Environment and Application tags defined.
```bash
usage: sshaws [-e Environment_tag] [-n instance_name] [-a Application_tag] [--silent true] [--region xx_yy_j] [-l user ][instance_name]
```
### Examples
```bash
# List all the instances from eu-west-1
> sshaws

Application: *   Environment: *   Name: *   Region: eu-west-1
---------------------------------------------------------

[0] front-prod-1 - 172.16.31.12 (i-0978df6a92b39d434, t2.nano)
[1] front-prod-2 - 172.16.24.226 (i-0b3914e83516392378, t2.nano)
[2] front-staging - 172.16.39.121 (i-0b3914e89237829ad1, t2.micro)
[3] back-prod - 172.16.33.21 (i-0b3914421237829ad1, t3.large)
[4] back-staging - 172.16.19.93 (i-0b391351e237829ad1, t3.micro)

Which one do you want to ssh in?
# WAITING INPUT FROM USER
3
>> Starting a new ssh session to 172.16.33.21
[...] Stablishing SSH connection with the desired server
 ```

```bash
# List all the instances with front on name, from eu-west-1
> sshaws front

Application: *   Environment: *   Name: front   Region: eu-west-1
---------------------------------------------------------

[0] front-prod-1 - 172.16.31.12 (i-0978df6a92b39d434, t2.nano)
[1] front-prod-2 - 172.16.24.226 (i-0b3914e83516392378, t2.nano)

Which one do you want to ssh in?
# WAITING INPUT FROM USER
0
>> Starting a new ssh session to 172.16.31.12
[...] Stablishing SSH connection with the desired server
 ```

```bash
# List all the instances from eu-west-1 with TAG Environment=staging
> sshaws -e staging

Application: *   Environment: staging   Name: *   Region: eu-west-1
---------------------------------------------------------

[0] front-staging - 172.16.39.121 (i-0b3914e89237829ad1, t2.micro)
[1] back-staging - 172.16.19.93 (i-0b391351e237829ad1, t3.micro)

Which one do you want to ssh in?
# WAITING INPUT FROM USER
1
>> Starting a new ssh session to 172.16.19.93
[...] Stablishing SSH connection with the desired server
 ```

```bash
# List all the instances from eu-west-1 with TAG Environment=staging and back in the name
> sshaws -e staging back

Application: *   Environment: staging   Name: back   Region: eu-west-1
---------------------------------------------------------

[0] back-staging - 172.16.19.93 (i-0b391351e237829ad1, t3.micro)

>> Starting a new ssh session to 172.31.2.93
# If  there is only one instance, it will launch ssh automatically without waiting user input
[...] Stablishing SSH connection with the desired server
```

```bash
# List all the instances from eu-west-1 with TAG Environment=staging and launching ssh with user "test"
> sshaws -e staging -l test

Application: *   Environment: staging   Name: *   Region: eu-west-1
---------------------------------------------------------

[0] front-staging - 172.16.39.121 (i-0b3914e89237829ad1, t2.micro)
[1] back-staging - 172.16.19.93 (i-0b391351e237829ad1, t3.micro)

Which one do you want to ssh in?
# WAITING INPUT FROM USER
0
>> Starting a new ssh session to test@172.31.36.121
[...] Stablishing SSH connection with the desired server
 ```


### Usage in depth

Download the binary in this repository and execute it, or compile by yourself following the next steps.

```bash
sshaws
  --app | -a
        Tag Application of the instance (default "*")
  --env | -e
        Tag Environment of the instance (default "*")
  --name | -n
        Instance Name (default "*")
  -l string
        SSH login name (default "")
  --region
        AWS Region (default "eu-west-1")
  --silent | -s
        Show only matching instances IPs
  --version | -v
        Show sshaws version
  --help | -h
        Show this text
```

### Diagram

![The flow to connect EC2:](images/ssm-sessionmanager.png#center)

### Prerequisites
There are some basic prerequisites :

- [**AWS CLI**](https://docs.aws.amazon.com/es_es/cli/latest/userguide/install-cliv2.html) tool
- Configure your **credentials** ~/.aws/ with the correct keys.
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
- [**Session Manager Plugin**](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html#plugin-version-history)
- **SSH**: You will need ssh, and, if necessary, configure the ssh connection (The configuration in `~/.ssh/configuration` will be applied to awssh).
- **[OPTIONAL] Golang**: If you want to compile `sshaws` you will need the `golang` executable.

##### Installing

#### From binary

#### For Linux
```bash
# Download latest version from https://github.com/uritau/sshaws/releases/latest
sudo wget -O /usr/local/bin/sshaws https://github.com/uritau/sshaws/releases/latest/download/sshaws

# Give it execution permissions
sudo chmod +x /usr/local/bin/sshaws
```
#### For MAC
```bash
# Download latest version from https://github.com/uritau/sshaws/releases/latest
sudo wget -O /usr/local/bin/sshaws https://github.com/uritau/sshaws/releases/latest/download/sshaws.mac

# Give it execution permissions
sudo chmod +x /usr/local/bin/sshaws
```
#### From source code
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

* [Golang](https://golang.org/)

## Contributing

Please create a fork and create a PR!
It would be awesome to have new proposals!

## Versioning

Travis is configured to automatically build and update the version. (You can change the current version on [version.go](pkg/cmd/version.go)).

## Authors

* [uritau](https://github.com/uritau), formerly [Oriol Tauleria](mailto:oriol.tauleria@gmail.com)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

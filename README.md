# SSHAWS

`sshaws` is a helper to simplify the *AWS instances ssh access*.

> It will show a list of instances depending on some specified tags and help you to ssh on it.

## Getting Started

### Usage

Download the binary in this repository and execute it, or compile by yourself following the next steps.

```
sshaws
  --app | -a
        Tag Application of the instance (default "*")
  --env | -e
        Tag Environment of the instance (default "*")
  --name | -n
        Instance Name (default "*")
  --region
        AWS Region (default "eu-west-1")
  --version | -v
        Show sshaws version
  --help | -h
        Show this text
```

### Prerequisites

 * **AWS CLI**: You need to have your ~/.aws/credentials in place and with the correct keys.
* **SSH**: You will need ssh, and, if necessary, configure the ssh connection (The configuration in `~/.ssh/configuration` will be applied to awssh).
* **[OPTIONAL] Golang**: If you want to compile `sshaws` you will need the `golang` executable.

### Installing

#### From binary

```bash
# Download latest version from https://github.com/uritau/sshaws/releases/latest
sudo wget -O /usr/local/bin/sshaws https://github.com/uritau/sshaws/releases/latest/download/sshaws

# Give it execution permissions
sudo chmod +x /usr/local/bin/sshaws
```

#### From source code
After downloading this repository:

```bash
# Enter repository folder
cd sshaws

# Install golang required packages:
go get .

# Build the binary
go build -o sshaws *.go

# Move to a desired folder
sudo mv sshaws /usr/local/bin/sshaws
```


## Running the tests

TO DO (No test available at this point)

## Built With

* [Golang](https://golang.org/)

## Contributing

Please create a fork and create a PR!
It would be awesome to have new proposals!

## Versioning

Travis is configured to automatically build and update the version. (You can change the current version on [version.go](version.go)).

## Authors

* [uritau](https://github.com/uritau), formerly [Oriol Tauleria](mailto:oriol.tauleria@gmail.com)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

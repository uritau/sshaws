# SSHAWS

`sshaws` is a helper to simplify the AWS instances dynamic ssh access.

> It will show a list of instances depending on some specified tags and help you to ssh on it.
Even being a go project it will use your ssh configuration files (so if you need to specify some key or username do it in there).

## Getting Started

### Usage

Download the binary in this repository and execute it, or compile by yourself following the next steps.

```
sshaws
  -app | -a
        Tag Application of the instance (default "*")
  -env | -e
        Tag Environment of the instance (default "*")
  -name | -n
        Instance Name (default "*")
  -region
        AWS Region (default "eu-west-1")
```

### Prerequisites

You will need ssh, and, if necessary, configure the ssh connection.
If you want to compile and build by yourself you will need go up and running.

### Installing

After downloading this repository, install golang required packages:

```
go get .
```

Build the binary

```
go build *.go
```

You can move it to somehwere in order to be able to use it from any point of the system.

```
sudo mv sshaws /usr/local/bin/sshaws
```

## Running the tests

TO DO (No test available at this point)

## Built With

* [Golang](https://golang.org/)

## Contributing

Please create a fork and create a PR! It will be awesome to have new proposals!

## Versioning

TO BE DEFINED

## Authors

* **Oriol Tauleria** - [uritau](https://github.com/uritau)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


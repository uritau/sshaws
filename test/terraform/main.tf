terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

locals {
  name = "sshaws-test-${random_pet.instance_name.id}"
}

data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"]

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
}

resource "random_pet" "instance_name" {}

resource "aws_instance" "instance" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t2.micro"

  tags = {
    Name        = local.name
    Application = "sshaws-test"
    Environment = "test"
  }
}

output "instance_name" {
  value = local.name
}

output "instance_id" {
  value = aws_instance.instance.id
}

output "instance_ip_address" {
  value = aws_instance.instance.private_ip
}

output "instance_type" {
  value = aws_instance.instance.instance_type
}

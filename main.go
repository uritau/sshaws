package main

import (
  "flag"
  "fmt"
  "log"
  "os/exec"

  "os"
  "strconv"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
)

type configuration struct {
  Region string
  Env    string
  App    string
}

func NewConfiguration(region, env, app string) *configuration {
  config := new(configuration)
  config.Region = region
  config.Env = env
  config.App = app
  return config
}

type instance struct {
  Name string
  IP   string
  ID   string
  Size string
}

func NewInstance(name, ip, id, size string) *instance {
  new_instance := new(instance)
  new_instance.Name = name
  new_instance.IP = ip
  new_instance.ID = id
  new_instance.Size = size
  return new_instance
}


func read_params() configuration {
  region := flag.String("region", "eu-west-1", "Default AWS region")
  env := flag.String("env", "prod", "Environment of the application")
  app := flag.String("app", "*", "Servers of the application")
  flag.Parse()
  return *NewConfiguration(*region, *env, *app)
}

func get_instances(region, env, app string) *ec2.DescribeInstancesOutput {

  fmt.Println("flag app:", app)
  fmt.Println("flag env:", env)
  fmt.Println("flag region:", region)
  sess, err := session.NewSession(&aws.Config{
    Region: aws.String(region)},
  )
  ec2svc := ec2.New(sess)
  params := &ec2.DescribeInstancesInput{
    Filters: []*ec2.Filter{
      {
        Name:   aws.String("tag:Environment"),
        Values: []*string{aws.String(env)},
      },
      {
        Name:   aws.String("tag:Application"),
        Values: []*string{aws.String(app)},
      }, {
        Name:   aws.String("instance-state-name"),
        Values: []*string{aws.String("running"), aws.String("pending")},
      },
    },
  }
  resp, err := ec2svc.DescribeInstances(params)
  if err != nil {
    fmt.Println("there was an error listing instances in", err.Error())
    log.Fatal(err.Error())
  }
  return resp
}

func launch_ssh(dest_ip string) {
  fmt.Printf("ssh %s \n", dest_ip)
  cmd := exec.Command("ssh", dest_ip)
  cmd.Stderr = os.Stderr
  stdout, err := cmd.Output()
  if err != nil {
    println(err.Error())
    return
  }
  print(string(stdout))
}

func select_instance() int{
  var input string
  fmt.Println("Which one do you want to shh in?")
  fmt.Scanln(&input)
  fmt.Println("PERFECT: ", input)
  index, err := strconv.Atoi(input)
  if err != nil {
    fmt.Println("What you entered is not a valid number.")
    select_instance()

  }
  return index
}

func show_instance_list(instance_list []instance) {
  for idx, inst := range instance_list {
    fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
  }
}

func main() {
  var instance_list []instance
  config := read_params()
  region := config.Region
  env := config.Env
  app := config.App
  resp := get_instances(region, env, app)



  for idx, _ := range resp.Reservations {
    for _, inst := range resp.Reservations[idx].Instances {
      name := ""
      for _, tag := range inst.Tags {
        if *tag.Key == "Name" {
          name = *tag.Value
        }
      }
      new_instance := NewInstance(name, *inst.PrivateIpAddress, *inst.InstanceId, *inst.InstanceType)
      instance_list = append(instance_list, *new_instance)


    }
  }
  show_instance_list(instance_list)
  instance_number := select_instance()
  launch_ssh(instance_list[instance_number].IP)
}

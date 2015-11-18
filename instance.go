package main

import (
    "fmt"

  "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
//    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/ec2"

)

func main() {
    svc := ec2.New(session.New())

    params := &ec2.DescribeInstancesInput{
     //   DryRun: aws.Bool(true),i-65e27fce  9f2dea3d
        
        InstanceIds: []*string{
           aws.String("i-65e27fce"), // Required
            // More values...
        },
       // MaxResults: aws.Int64(1),
       // NextToken:  aws.String("String"),
        }
    
    resp, err := svc.DescribeInstances(params)

    if err != nil {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("%T\n",resp.Reservations[0].Instances[0])

//fmt.Println(resp)
    // Pretty-print the response data.
   fmt.Printf("%v\n",*resp.Reservations[0].Instances[0].PrivateIpAddress)
       if resp.Reservations[0].Instances[0].PublicIpAddress != nil{
           fmt.Printf("%v\n",*resp.Reservations[0].Instances[0].PublicIpAddress)
       }
                fmt.Printf("%v\n",*resp.Reservations[0].Instances[0].State.Name)



}
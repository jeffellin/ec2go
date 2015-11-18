package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
//    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/cloudformation"

)

func main() {

    stackId := "arn:aws:cloudformation:us-east-1:926164720730:stack/StackName3/66cb5360-8d41-11e5-b3a4-500150b34c18"
     fmt.Println(stackId)
    // Create an EC2 service object in the "us-west-2" region
    // Note that you can also configure your region globally by
    // exporting the AWS_REGION environment variable
    svc := cloudformation.New(session.New())

    params := &cloudformation.DescribeStacksInput{
        StackName: aws.String("sdf"),
        
    }
    resp, err := svc.DescribeStacks(params)

    if err != nil {
        fmt.Println("Houston we have a problem")
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
    }else{
        for _,element := range resp.Stacks[0].Outputs {
  // element is the element from someSlice for where we are
            outputV := *element.OutputValue
            if(*element.OutputKey=="PrivateIp"){
               fmt.Printf("%v\n", outputV)
            }
            if(*element.OutputKey=="InstanceID"){
               fmt.Printf("%v\n", outputV)
            }
            
  // fmt.Println("---")
}   
    }



    // Pretty-print the response data.
    //fmt.Println(resp)


}
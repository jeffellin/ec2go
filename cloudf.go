package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"time"
	//    "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func main() {
	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := cloudformation.New(session.New())

	params := &cloudformation.CreateStackInput{
		StackName:   aws.String("StackName5"),
		TemplateURL: aws.String("https://s3.amazonaws.com/com.tamr.fe.users/jellin/docker.json"),
		Parameters: []*cloudformation.Parameter{
			{ // Required
				ParameterKey:   aws.String("KeyName"),
				ParameterValue: aws.String("FE-CI"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateStack(params)

	if err != nil {
		fmt.Println("Houston we have a problem")
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	} else {
		// Pretty-print the response data.
		fmt.Println(resp)

		stackId := *resp.StackId

		fmt.Println(stackId)

		WaitFor(stackAvailable)

	}

}

func stackAvailable() bool {

	svc := cloudformation.New(session.New())

	params := &cloudformation.DescribeStacksInput{
		StackName: aws.String("StackName5"),
	}
	resp, err := svc.DescribeStacks(params)

	if err != nil {
		fmt.Println("Houston we have a problem")
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return false
	}
	if *resp.Stacks[0].StackStatus == cloudformation.ResourceStatusCreateComplete {
		return true
	} else {
		fmt.Println("...Not Yet")
		return false
	}
}

//should be import from mcnutils
func WaitFor(f func() bool) error {
	return WaitForSpecific(f, 60, 3*time.Second)
}

func WaitForSpecific(f func() bool, maxAttempts int, waitInterval time.Duration) error {
	return WaitForSpecificOrError(func() (bool, error) {
		return f(), nil
	}, maxAttempts, waitInterval)
}

func WaitForSpecificOrError(f func() (bool, error), maxAttempts int, waitInterval time.Duration) error {
	for i := 0; i < maxAttempts; i++ {
		stop, err := f()
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
		time.Sleep(waitInterval)
	}
	return fmt.Errorf("Maximum number of retries (%d) exceeded", maxAttempts)
}
//end imports we will use from mcnutils
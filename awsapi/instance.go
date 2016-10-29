package awsapi

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/ec2"
)

type client interface {
    DescribeInstanceStatus(input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error)
}

func getInstanceStatus(c client, instanceID string) (string, error) {
    input := &ec2.DescribeInstanceStatusInput{
        InstanceIds: []*string{
            aws.String(instanceID),
        },
    }
    output, err := c.DescribeInstanceStatus(input)
    status := output.InstanceStatuses[0]
    return status.InstanceState.String(), err
}

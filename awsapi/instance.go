package awsapi

import (
    "github.com/aws/aws-sdk-go/service/ec2"
)

func getInstanceStatus(ec2client *ec2.EC2) {
    ec2client.DescribeInstanceStatus(input)
}

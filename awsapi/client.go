package awsapi

import (
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func getEC2Client() (ec2client *ec2.EC2) {
    sess := session.New()
    ec2client = ec2.New(sess)
    return ec2client
}

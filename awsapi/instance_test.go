package awsapi

import (
    "github.com/aws/aws-sdk-go/aws"
    ec2 "github.com/aws/aws-sdk-go/service/ec2"
    "testing"
)

type testclient struct {
}

func (t testclient) DescribeInstanceStatus(input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
    instancesIDs := input.InstanceIds
    state := &ec2.InstanceState{
        Code: aws.Int64(16),
        Name: aws.String("running"),
    }

    status := &ec2.InstanceStatus{
        AvailabilityZone: aws.String("Test Zone"),
        InstanceId:       instancesIDs[0],
        InstanceState:    state,
    }

    ouput := &ec2.DescribeInstanceStatusOutput{
        InstanceStatuses: []*ec2.InstanceStatus{status},
    }
    return ouput, nil
}
func TestGetInstanceStatus(t *testing.T) {
    tc := testclient{}
    _, err := getInstanceStatus(tc, "i-testinstance")

    if err != nil {
        t.Error("Error calling getInstanceStatus")
    }
}

func TestStartInstance(t *testing.T) {

}

func TestStopInstance(t *testing.T) {

}

func TestTerminateInstance(t *testing.T) {

}

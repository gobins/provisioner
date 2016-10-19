package awsapi

import "testing"

func TestgetEC2Client(t *testing.T) {
    c := getEC2Client()

    if c == nil {
        t.Error("Error calling getEC2Client")
    }

}

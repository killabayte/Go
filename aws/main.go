package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Read the AWS region from the environment variable
	region := os.Getenv("AWS_REGION")
	if region == "" {
		fmt.Println("AWS_REGION environment variable is not set.")
		return
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	svc := ec2.New(sess)

	input := &ec2.DescribeSecurityGroupsInput{}
	result, err := svc.DescribeSecurityGroups(input)

	if err != nil {
		fmt.Println("Error describing security groups:", err)
		return
	}

	file, err := os.OpenFile("security_groups.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	// Iterate through security groups and write each to the JSON file
	for _, sg := range result.SecurityGroups {
		// Create a struct to represent the security group details
		securityGroup := struct {
			ID            string
			Name          string
			InboundRules  []ec2.IpPermission
			OutboundRules []ec2.IpPermission
		}{
			ID:            *sg.GroupId,
			Name:          *sg.GroupName,
			InboundRules:  sg.IpPermissions,
			OutboundRules: sg.IpPermissionsEgress,
		}

		// Encode the security group details as JSON and write to the file
		err := encoder.Encode(securityGroup)
		if err != nil {
			fmt.Println("Error writing to JSON file:", err)
			return
		}
	}
}

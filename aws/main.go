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

	// Create an AWS session using the specified region
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an EC2 service client
	svc := ec2.New(sess)

	// Describe all security groups
	input := &ec2.DescribeSecurityGroupsInput{}
	result, err := svc.DescribeSecurityGroups(input)

	if err != nil {
		fmt.Println("Error describing security groups:", err)
		return
	}

	// Open a JSON file in append mode for writing
	file, err := os.OpenFile("security_groups.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	// Iterate through security groups and write each to the JSON file with a newline separator
	for _, sg := range result.SecurityGroups {
		// Convert sg.IpPermissions from []*ec2.IpPermission to []ec2.IpPermission
		inboundRules := make([]ec2.IpPermission, len(sg.IpPermissions))
		for i, rule := range sg.IpPermissions {
			inboundRules[i] = *rule
		}

		// Convert sg.IpPermissionsEgress from []*ec2.IpPermission to []ec2.IpPermission
		outboundRules := make([]ec2.IpPermission, len(sg.IpPermissionsEgress))
		for i, rule := range sg.IpPermissionsEgress {
			outboundRules[i] = *rule
		}

		// Create a struct to represent the security group details
		securityGroup := struct {
			ID            string
			Name          string
			InboundRules  []ec2.IpPermission
			OutboundRules []ec2.IpPermission
		}{
			ID:            *sg.GroupId,
			Name:          *sg.GroupName,
			InboundRules:  inboundRules,  // Use the converted slice
			OutboundRules: outboundRules, // Use the converted slice
		}

		// Encode the security group details as JSON and write to the file
		err := encoder.Encode(securityGroup)
		if err != nil {
			fmt.Println("Error writing to JSON file:", err)
			return
		}

		fmt.Println("Wrote security group details to JSON file:", *sg.GroupName)
	}
}

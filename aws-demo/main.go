package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	var (
		instanceId string
		err        error
	)

	if instanceId, err = createEC2(context.Background(), "eu-central-1"); err != nil {
		fmt.Printf("createEC2 error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Instance ID: %s", instanceId)
}

func createEC2(ctx context.Context, region string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("Unable to load SDK config, %s", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)

	keyPairs, err := ec2Client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("key-name"),
				Values: []string{"go-aws-sdk"},
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("Unable to DescribeKeyPair, %s", err)
	}

	if len(keyPairs.KeyPairs) == 0 {
		keyPair, err := ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
			KeyName: aws.String("go-aws-sdk"),
		})
		if err != nil {
			return "", fmt.Errorf("Unable to crete KeyPair, %s", err)
		}
		err = os.WriteFile("go-aws-sdk.pem", []byte(*keyPair.KeyMaterial), 0600)
		if err != nil {
			return "", fmt.Errorf("Wrire keyFile error: %s", err)
		}
	}

	imageOutput, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})
	if err != nil {
		return "", fmt.Errorf("Describe image error: %s", err)
	}
	if len(imageOutput.Images) == 0 {
		return "", fmt.Errorf("imageOutput.Images is of 0 length")
	}

	instance, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId:      imageOutput.Images[0].ImageId,
		KeyName:      aws.String("go-aws-sdk"),
		InstanceType: types.InstanceTypeT3Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	})
	if err != nil {
		return "", fmt.Errorf("Run instance error: %s", err)
	}

	if len(instance.Instances) == 0 {
		return "", fmt.Errorf("Instances is of 0 length")
	}

	return *instance.Instances[0].InstanceId, nil
}

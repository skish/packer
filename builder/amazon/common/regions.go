package common

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

func getValidationSession() *ec2.EC2 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2conn := ec2.New(sess)
	return ec2conn
}

func listEC2Regions(ec2conn ec2iface.EC2API) []string {
	var regions []string
	resultRegions, err := ec2conn.DescribeRegions(nil)
	if err != nil {
		log.Printf("DescribeRegions: %v", err)
	}
	for _, region := range resultRegions.Regions {
		regions = append(regions, *region.RegionName)
	}

	return regions
}

// ValidateRegion returns true if the supplied region is a valid AWS
// region and false if it's not.
func ValidateRegion(region string, ec2conn ec2iface.EC2API) error {
	regions := listEC2Regions(ec2conn)
	for _, valid := range regions {
		if region == valid {
			return nil
		}
	}
	return fmt.Errorf("Invalid region %s, available regions: %v", region, regions)
}

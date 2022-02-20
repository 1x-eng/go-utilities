package rdsdataapi

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
)

func getNewSession(awsRegion *string) (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(*awsRegion),
	})

	if err != nil {
		panic(err)
	}

	return session, nil
}

func GetNewClient(awsRegion *string) *rdsdataservice.RDSDataService {
	session, _ := getNewSession(awsRegion)
	rdsDataServiceClient := rdsdataservice.New(session)

	return rdsDataServiceClient
}

func Heartbeat() string {
	return "Alive"
}

package rdsdataapi

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
)

type AuroraRDSConfig struct {
	database             *string
	resourceArn          *string
	secretArn            *string
	rdsDataServiceClient *rdsdataservice.RDSDataService
}

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

func ExecuteSQL(rdsConfig *AuroraRDSConfig, sql *string) {
	req, resp := rdsConfig.rdsDataServiceClient.ExecuteStatementRequest(&rdsdataservice.ExecuteStatementInput{
		Database:    aws.String(*rdsConfig.database),
		ResourceArn: aws.String(*rdsConfig.resourceArn),
		SecretArn:   aws.String(*rdsConfig.secretArn),
		Sql:         aws.String(*sql),
	})

	err1 := req.Send()
	if err1 == nil { // resp is now filled
		fmt.Println("Response:", resp)
	} else {
		fmt.Println("error:", err1)
	}
}

func Heartbeat() string {
	return "Alive"
}

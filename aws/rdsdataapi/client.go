package rdsdataapi

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
)

type AuroraRDSConfig struct {
	AwsRegion            *string
	Database             *string
	ResourceArn          *string
	SecretArn            *string
	RdsDataServiceClient *rdsdataservice.RDSDataService
}

type ExecuteSQLResponse struct {
    GeneratedFields        []string
    NumberOfRecorcdUpdated  int
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

func GenerateAuroraRDSConfig(awsRegion *string, dbName *string, dbResourceArn *string, dbSecretArn *string) *AuroraRDSConfig {
	rdsConfig := AuroraRDSConfig{
		AwsRegion:            awsRegion,
		Database:             dbName,
		ResourceArn:          dbResourceArn,
		SecretArn:            dbSecretArn,
		RdsDataServiceClient: GetNewClient(awsRegion),
	}

	return &rdsConfig
}

func GetNewClient(awsRegion *string) *rdsdataservice.RDSDataService {
	session, _ := getNewSession(awsRegion)
	rdsDataServiceClient := rdsdataservice.New(session)

	return rdsDataServiceClient
}

func ExecuteSQL(rdsConfig *AuroraRDSConfig, sql *string) *ExecuteSQLResponse {
	req, resp := rdsConfig.RdsDataServiceClient.ExecuteStatementRequest(&rdsdataservice.ExecuteStatementInput{
		Database:    aws.String(*rdsConfig.Database),
		ResourceArn: aws.String(*rdsConfig.ResourceArn),
		SecretArn:   aws.String(*rdsConfig.SecretArn),
		Sql:         aws.String(*sql),
	})

	err := req.Send()

	if err != nil {
		panic(err)
	}

	fmt.Println("Response after executing sql = ", resp)
    
    return &ExecuteSQLResponse{
        GeneratedFields: resp.GeneratedFields,
        NumberOfRecorcdUpdated: resp.NumberOfRecorcdUpdated
    }
}

func Heartbeat() string {
	return "Alive"
}

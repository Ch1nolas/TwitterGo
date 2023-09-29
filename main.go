package main

import (
	"context"

	"os"

	"Github/Ch1nolas/TwitterGo/awsgo"
	"Github/Ch1nolas/TwitterGo/secretmanager"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.IniciarAWS()

	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. Deben incluit 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

}

func ValidoParametros() bool {
	_, traeParametros := os.LookupEnv("SecretName")
	if !traeParametros {
		return traeParametros
	}

	_, traeParametros = os.LookupEnv("BucketName")
	if !traeParametros {
		return traeParametros
	}

	_, traeParametros = os.LookupEnv("UrlPrefix")
	if !traeParametros {
		return traeParametros
	}

	return traeParametros
}

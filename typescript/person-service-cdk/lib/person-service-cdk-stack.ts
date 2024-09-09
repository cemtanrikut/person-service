import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as apigateway from 'aws-cdk-lib/aws-apigateway';

export class PersonServiceStack extends Stack {
    constructor(scope: Construct, id: string, props?: StackProps) {
        super(scope, id, props);

        const table = new dynamodb.Table(this, 'Persons', {
            partitionKey: { name: 'id', type: dynamodb.AttributeType.STRING },
            billingMode: dynamodb.BillingMode.PAY_PER_REQUEST
        });

        const personLambda = new lambda.Function(this, 'PersonHandler', {
            runtime: lambda.Runtime.GO_1_X,
            handler: 'main',
            code: lambda.Code.fromAsset('/Users/cemtanrikut/Documents/GitHub/person-service/typescript/person-service-cdk/main'),
            environment: {
                TABLE_NAME: table.tableName,
            }
        });

        const api = new apigateway.RestApi(this, 'PersonApi', {
            restApiName: 'Person Service API',
            description: 'This API services person data.'
        });

        const persons = api.root.addResource('persons');
        const postIntegration = new apigateway.LambdaIntegration(personLambda, {
            requestTemplates: { 'application/json': '{ "statusCode": "200" }' }
        });

        persons.addMethod('POST', postIntegration);
        persons.addMethod('GET', postIntegration);  
        
        table.grantReadWriteData(personLambda);
    }
}

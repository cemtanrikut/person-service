import { Template } from '@aws-cdk/assertions';
import { PersonServiceStack } from './person-service-cdk-stack';
import { App } from 'aws-cdk-lib';

describe('PersonServiceStack', () => {
  const app = new App();
  const stack = new PersonServiceStack(app, 'PersonServiceStackTest');

  // Örnek bir template oluşturun
  const template = Template.fromStack(stack);

  it('has a DynamoDB table', () => {
    template.resourceCountIs('AWS::DynamoDB::Table', 1);
  });

  it('has a Lambda function', () => {
    template.resourceCountIs('AWS::Lambda::Function', 1);
  });

  it('has an API Gateway', () => {
    template.resourceCountIs('AWS::ApiGateway::RestApi', 1);
  });

  // Lambda fonksiyonunun doğru ortam değişkenlerine sahip olduğunu kontrol edin
  it('Lambda has correct environment variables', () => {
    template.hasResourceProperties('AWS::Lambda::Function', {
      Environment: {
        Variables: {
          TABLE_NAME: {
            Ref: expect.anything() // DynamoDB tablo isminin referansını kontrol eder
          }
        }
      }
    });
  });
});

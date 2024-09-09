# Person Service

Person Service is a serverless microservice implemented using AWS CDK and TypeScript. It manages person data, allowing users to create and list person entries through a REST API.

## Prerequisites

Before you begin, ensure you have the following installed:
- Node.js (version 12.x or later)
- npm (which comes with Node.js)
- AWS CLI
- AWS account and configured AWS credentials

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Setup

Clone the repository to your local machine:

```bash
git clone https://github.com/cemtanrikut/person-service.git
cd person-service
```

Install the necessary dependencies:

```bash
npm install
```

### Build

Compile the TS to JS
```bash
npm run build
```

### Deploy

Deploy the CDK stack to your AWS account:

```bash
cdk deploy
```

Ensure you have the AWS CLI configured with your account and the correct region

### Testing

Run unit tests to ensure everything is set up correctly:

```bash
npm test
```
### Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any 
contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
3. Commit your Changes (git commit -m 'Add some AmazingFeature')
4. Push to the Branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

### License

Distributed under the MIT License. See LICENSE for more information.

### Useful commands

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run test`    perform the jest unit tests
* `npx cdk deploy`  deploy this stack to your default AWS account/region
* `npx cdk diff`    compare deployed stack with current state
* `npx cdk synth`   emits the synthesized CloudFormation template

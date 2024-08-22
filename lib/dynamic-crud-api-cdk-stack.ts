import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";
import * as cdk from "aws-cdk-lib";
import { EndpointType, LambdaIntegration } from "aws-cdk-lib/aws-apigateway";
import { Construct } from "constructs";

export class DynamicCrudApiCdkStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const lambda = new GoFunction(this, "dynamic-crud-api-handler", {
      functionName: "dynamic-crud-api",
      entry: "src/cmd/api",
      memorySize: 768,
      timeout: cdk.Duration.seconds(60),
      environment: {},
    });

    const api = new cdk.aws_apigateway.RestApi(this, "dynamic-crud-api", {
      restApiName: "dynamic-crud-api",
      endpointTypes: [EndpointType.REGIONAL],
    });

    api.root.addResource("{proxy+}").addMethod(
      "ANY",
      new LambdaIntegration(lambda, {
        proxy: true,
      }),
    );
  }
}

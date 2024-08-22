#!/usr/bin/env node
import * as cdk from "aws-cdk-lib";
import * as dotenv from "dotenv";
import "source-map-support/register";
import env from "../env";
import { DynamicCrudApiCdkStack } from "../lib/dynamic-crud-api-cdk-stack";

dotenv.config();

const app = new cdk.App();
new DynamicCrudApiCdkStack(app, "DynamicCrudApiCdkStack", {
  env: {
    region: env.AWS.REGION,
  },
});

cdk.Tags.of(app).add("Project", env.TAG.PROJECT);
cdk.Tags.of(app).add("Owner", env.TAG.OWNER);

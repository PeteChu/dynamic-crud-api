import * as env from "env-var";
import * as dotenv from "dotenv";

dotenv.config();

export default {
  NODE_ENV: env
    .get("NODE_ENV")
    .required()
    .asEnum(["development", "production"]),
  AWS: {
    REGION: env.get("AWS_REGION").required().asString(),
  },
  TAG: {
    PROJECT: env.get("TAG_PROJECT").required().asString(),
    OWNER: env.get("TAG_OWNER").required().asString(),
  },
};

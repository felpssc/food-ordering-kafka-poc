require("dotenv/config");
const aws = require('aws-sdk');

class S3Storage {
  constructor() {
    this.s3 = new aws.S3({
      region: process.env.AWS_DEFAULT_REGION,
    });
  }

  async upload(buffer, fileName) {

    const Key = `${fileName}${new Date().getTime()}`;

    await this.s3.putObject({
      Bucket: process.env.AWS_BUCKET_NAME,
      Key,
      Body: buffer,
      ContentType: 'application/pdf',
      ACL: 'public-read',
    }).promise();

    return `https://${process.env.AWS_BUCKET_NAME}.s3.us-east-1.amazonaws.com/${Key}`;
  }
}

module.exports = { S3Storage };
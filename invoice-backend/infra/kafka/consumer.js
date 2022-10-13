const kafka = require('kafkajs');

const kafkaConfig = {
  clientId: 'invoice-backend',
  brokers: ['host.docker.internal:9094'],
  retry: {
    initialRetryTime: 300,
    retries: 10,
  },
};

const kafkaConsumer = new kafka.Kafka({
  clientId: kafkaConfig.clientId,
  brokers: kafkaConfig.brokers,
  retry: kafkaConfig.retry,
});

module.exports = {
  kafkaConsumer,
}
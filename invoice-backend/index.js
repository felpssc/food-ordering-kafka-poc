const { Invoice } = require('./helpers/Invoice');
const { HtmlToPDF } = require('./helpers/HtmlToPDF');

const { kafkaConsumer } = require('./infra/kafka/consumer');
const { S3Storage } = require('./helpers/S3storage');

const consumer = kafkaConsumer.consumer({ groupId: 'invoice-backend' });

const run = async () => {
  await consumer.connect();
  await consumer.subscribe({ topic: 'order_confirmed' });

  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      const order = JSON.parse(message.value.toString());
      order.products = JSON.parse(order.products);
      order.isWatermark = true;
      
      const invoice = new Invoice(order);
      const html = await invoice.generate();

      const pdf = await new HtmlToPDF(html).generate();
      
      const s3 = new S3Storage();
      const url = await s3.upload(pdf, order.document_id);

      console.log(`Order ${order.document_id} invoice: ${url}`);
    },
  });
}

(async () => { await run(); })();
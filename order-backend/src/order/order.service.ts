import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientKafka } from '@nestjs/microservices';
import { InjectModel } from '@nestjs/mongoose';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { Model } from 'mongoose';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order, OrderDocument } from './entities/order.entity';

@Injectable()
export class OrderService implements OnModuleInit {
  public kafkaProducer: Producer;

  constructor(
    @InjectModel(Order.name) private orderModel: Model<OrderDocument>,
    @Inject('KAFKA_SERVICE') public kafkaClient: ClientKafka,
  ) {}

  async onModuleInit() {
    this.kafkaProducer = await this.kafkaClient.connect();
  }

  async create({ client_name, client_email, products }: CreateOrderDto) {
    const createdAt = new Date();

    const total_price = products.reduce(
      (total, product) => total + product.price,
      0,
    );

    const createdOrder = new this.orderModel({
      client_name,
      client_email,
      products,
      total_price,
      status: 'pending',
      created_at: createdAt,
    });

    await this.kafkaProducer.send({
      topic: 'order_details',
      messages: [
        {
          key: 'order_details',
          value: JSON.stringify({
            document_id: String(createdOrder._id),
            client_name,
            client_email,
            products: JSON.stringify(products),
            total_price: parseFloat(total_price.toFixed(2)),
            status: 'pending',
            created_at: String(createdAt),
          }),
        },
      ],
      acks: -1,
    });

    return await createdOrder.save();
  }

  findAll() {
    return `This action returns all order`;
  }

  findOne(id: number) {
    return `This action returns a #${id} order`;
  }

  update(id: number, updateOrderDto: UpdateOrderDto) {
    return `This action updates a #${id} order`;
  }

  remove(id: number) {
    return `This action removes a #${id} order`;
  }
}

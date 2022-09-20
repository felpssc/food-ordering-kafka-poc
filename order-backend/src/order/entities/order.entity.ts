import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Product } from '../../products/entities/product.entity';

import { Document } from 'mongoose';

type Status = 'pending' | 'confirmed' | 'cancelled';

export type OrderDocument = Order & Document;

@Schema()
export class Order {
  @Prop()
  client_name: string;

  @Prop()
  client_email: string;

  @Prop()
  total_price: number;

  @Prop()
  products: Product[];

  @Prop()
  status: Status;

  @Prop()
  created_at: Date;
}

export const OrderSchema = SchemaFactory.createForClass(Order);

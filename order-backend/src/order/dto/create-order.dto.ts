import { Product } from '../../products/entities/product.entity';
export class CreateOrderDto {
  client_name: string;
  client_email: string;
  products: Product[];
}

import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { getDataSourceToken } from '@nestjs/typeorm';
import { DataSource } from 'typeorm';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await app.init();

  const dataSource = app.get<DataSource>(getDataSourceToken());
  await dataSource.synchronize(true);

  const productRepository = dataSource.getRepository('Product');

  await productRepository.insert([
    {
      id: 'e2cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 1',
      description: 'product 1 description',
      image_url: 'https://via.placeholder.com/150',
      price: 180,
    },
    {
      id: 'e2cc6d2b-45e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 2',
      description: 'product 2 description',
      image_url: 'https://via.placeholder.com/150',
      price: 280,
    },
    {
      id: 'f4cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 3',
      description: 'product 3 description',
      image_url: 'https://via.placeholder.com/150',
      price: 140,
    },
    {
      id: 'g9cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 4',
      description: 'product 4 description',
      image_url: 'https://via.placeholder.com/150',
      price: 150,
    },
    {
      id: 'k2cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 5',
      description: 'product 5 description',
      image_url: 'https://via.placeholder.com/150',
      price: 110,
    },
    {
      id: 'a2cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 6',
      description: 'product 6 description',
      image_url: 'https://via.placeholder.com/150',
      price: 135,
    },
    {
      id: 'm2cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 7',
      description: 'product 7 description',
      image_url: 'https://via.placeholder.com/150',
      price: 100,
    },
    {
      id: 'b6cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 8',
      description: 'product 8 description',
      image_url: 'https://via.placeholder.com/150',
      price: 80,
    },
    {
      id: 'y6cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 9',
      description: 'product 9 description',
      image_url: 'https://via.placeholder.com/150',
      price: 155,
    },
    {
      id: 'a3cc6d2b-54e2-4654-ba47-f5ccf92bd3b8',
      name: 'Product 10',
      description: 'product 10 description',
      image_url: 'https://via.placeholder.com/150',
      price: 189,
    },
  ]);

  await app.close();
}
bootstrap();

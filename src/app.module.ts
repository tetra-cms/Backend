import { Module } from '@nestjs/common';
import { AuthModule } from './auth/auth.module';
import { PrismaModule } from './prisma/prisma.module';
import { UserController } from './user/user.controller';
import { UserService } from './user/user.service';
import { ClientController } from './client/client.controller';
import { ClientService } from './client/client.service';
import { ProductController } from './product/product.controller';
import { ProductService } from './product/product.service';
import { CategoryController } from './category/category.controller';
import { CategoryService } from './category/category.service';

@Module({
  imports: [AuthModule, PrismaModule],
  controllers: [UserController, ClientController, ProductController, CategoryController],
  providers: [UserService, ClientService, ProductService, CategoryService],
})
export class AppModule {}

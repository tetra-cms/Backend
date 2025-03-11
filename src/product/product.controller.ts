import { Controller, Get, Param } from '@nestjs/common';
import { ProductService } from './product.service';

@Controller('product')
export class ProductController {
    constructor(
            private productService: ProductService) {}

    @Get('list')
    async getProductList()
    {
        return await this.productService.getProductList();
    }

    @Get(':id')
    async getProductInfo(@Param() id)
    {
        return await this.productService.getProductById(Number(id));
    }
}
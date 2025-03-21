import { Controller, Get, Param, Query, Req } from '@nestjs/common';
import { ProductService } from './product.service';

@Controller('product')
export class ProductController {
    constructor(
            private productService: ProductService) {}

    @Get('list')
    async getProductList(@Query('category') category: string)
    {
        return await category ? this.productService.getProductListByCategory(Number(category)) : await this.productService.getProductList();
    }

    @Get(':id')
    async getProductInfo(@Param() params: any)
    {
        return await this.productService.getProductById(Number(params.id));
    }
}
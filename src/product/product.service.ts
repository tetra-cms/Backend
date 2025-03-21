import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class ProductService {
    
    constructor(
        private prisma: PrismaService) {}


    /**
     * @description Returns list of selling items in platform
     * @author fet1sov
     */
    async getProductList()
    {
        return this.prisma.product.findMany({});
    }

    /**
     * @param {number} categoryId - category unique identifier
     * @description Returns list of products by selected category
     * @author fet1sov
     */
    async getProductListByCategory(categoryId: number)
    {
        return this.prisma.product.findMany({
            where: {
                categoryId: categoryId
            }
        });
    }

    /**
     * @param {number} productId - unique product identifier
     * @description Returns list of selling items in platform
     * @author fet1sov
     */
    async getProductById(productId: number)
    {
        const productInfo = await this.prisma.product.findUnique({
            where: {
                id: productId
            }
        });

        const categoryInfo = await this.prisma.category.findUnique({
            where: {
                id: productInfo?.categoryId
            }
        });

        return {
            ...productInfo,
            categoryInfo: categoryInfo
        }
    }   
}

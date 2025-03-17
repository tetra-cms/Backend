import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class CategoryService {
    constructor(
        private prisma: PrismaService) {}

    /**
     * @description Returns list item categories in platform
     * @author fet1sov
     */
    async getCategoryList()
    {
        return this.prisma.category.findMany({});
    }

    /**
     * @param {number} categoryId - unique category identifier
     * @description Returns list of selling items in platform
     * @author fet1sov
     */
    async getCategoryById(categoryId: number)
    {
        return this.prisma.category.findUnique({
            where: {
                id: categoryId
            }
        });
    }   
}

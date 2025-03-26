import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class ContentService {
    constructor(
        private prisma: PrismaService,
    ) {

    }

    /**
     * @description Returns all pages list
     * @author fet1sov
     */
    async getPagesList()
    {
        return await this.prisma.pages.findMany({
            select: {
                id: true,
                route: true
            }
        });
    }

    /**
     * @param {string} route - page route
     * @description Give the page markdown by using the route
     * @author fet1sov
     */
    async getPageContent(route: string)
    {
        return await this.prisma.pages.findFirst({
            where: {
                route: route
            }
        });
    }
}

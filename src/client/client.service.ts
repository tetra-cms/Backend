import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class ClientService {

    constructor(
            private prisma: PrismaService,
    ) {
    
    }

    async getClientList(ownerid: number)
    {
        return this.prisma.client.findMany({
            where: {
                userId: ownerid
            }
        });
    }

    async getClientById(clientId: number)
    {
        return this.prisma.client.findFirst({
            where: {
                userId: clientId
            }
        });
    }
}

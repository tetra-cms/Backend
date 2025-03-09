import { Controller, ForbiddenException, Get, Param, Req, UseGuards } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';
import { ClientService } from './client.service';

@Controller('client')
export class ClientController {
    constructor(
            private clientService: ClientService) {}

    @UseGuards(AuthGuard('jwt'))
    @Get('list')
    async getClientList(@Req() request) {
        return await this.clientService.getClientList(request.user.id);
    }

    @UseGuards(AuthGuard('jwt'))
    @Get('info/:id')
    async getClientInfo(@Req() request, @Param() id) {
        const client = await this.clientService.getClientById(id);

        if (client?.userId === request.user.id)
        {
            return client;
        } else {
            return new ForbiddenException("Don't belong this user");
        }
    }
}

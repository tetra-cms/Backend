import { Controller, Get, Req, UseGuards } from '@nestjs/common';
import { UserService } from './user.service';
import { AuthGuard } from '@nestjs/passport';

@Controller('user')
export class UserController {
    constructor(
        private userService: UserService) {}

    @UseGuards(AuthGuard('jwt'))
    @Get('profile')
    async getProfile(@Req() request) {
        const userInfo = request.user;
        delete userInfo["password"];

        return userInfo;
    }
}

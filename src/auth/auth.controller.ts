import { Body, Controller, ForbiddenException, Ip, Post } from '@nestjs/common';
import { AuthService } from './auth.service';
import { RegisterUserDto } from './dto/register-user.dto';

@Controller('auth')
export class AuthController {
  constructor(private authService: AuthService) {}

  @Post('login')
  async login(@Body() body: { username: string; password: string }) {
    const user = await this.authService.validateUser(body.username, body.password);
    if (!user) throw new ForbiddenException();
    return this.authService.login(user);
  }

  @Post('register')
  async register(@Body() body: RegisterUserDto, @Ip() ip) {
    return this.authService.register(body, ip);
  }
}
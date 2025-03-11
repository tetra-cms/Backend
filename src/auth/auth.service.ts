import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';
import { PrismaService } from '../prisma/prisma.service';
import { RegisterUserDto } from './dto/register-user.dto';
import { $Enums } from '@prisma/client';

@Injectable()
export class AuthService {
  constructor(
    private prisma: PrismaService,
    private jwtService: JwtService,
  ) {}

  async validateUser(username: string, pass: string): Promise<any> {
    const user = await this.prisma.user.findUnique({ where: { username } });
    if (user && await bcrypt.compare(pass, user.password)) {
      const { password, ...result } = user;
      return result;
    }
    return null;
  }

  async login(user: any) {
    const payload = { username: user.email, sub: user.id };
    return {
      access_token: this.jwtService.sign(payload),
    };
  }

  async register(userData: RegisterUserDto, ipAddress: string) {
    const hashedPassword = await bcrypt.hash(userData.password, 10);

    const user = await this.prisma.user.create({
      data: {
        username: userData.username,
        password: hashedPassword,
        email: userData.email,
        role: $Enums.Role.USER,
        loginIp: ipAddress,
        regIp: ipAddress
      },
    });

    return {
      createdAt: user.createdAt
    }
  }
}
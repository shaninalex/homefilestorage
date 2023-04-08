import { Controller, Post, Req } from '@nestjs/common';
import { Request } from 'express';
import { JwtService } from '@nestjs/jwt';


interface LoginRequest {
    email: string,
    password: string
}

interface LoginResponse {
    access: string
    refresh: string
}


@Controller('api/v1/auth')
export class AuthController {

    constructor(private jwtService: JwtService) {}

    @Post('login')
    getAllNotes(@Req() request: Request): LoginResponse {
        const body: LoginRequest = <LoginRequest>request.body;
        const payload = { username: body.email, sub: 1 };
        return {
            access: this.jwtService.sign(payload),
            refresh: this.jwtService.sign(payload),
        };
    }
}

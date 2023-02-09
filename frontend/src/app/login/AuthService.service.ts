import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { LoginRequest, LoginResponse, RegisterRequest } from './typedefs';


@Injectable()
export class AuthService {
    constructor(private http: HttpClient) { }

    configUrl: string = '/api/v1/auth/login';
    url_register: string = '/api/v1/auth/create/';
    url_refresh: string = '/api/v1/auth/refresh/';

    login(login_data: LoginRequest) {
        return this.http.post<LoginResponse>(this.configUrl, login_data);
    }

    register(register_data: RegisterRequest) {
        return this.http.post<RegisterRequest>(this.url_register, register_data);
    }
}
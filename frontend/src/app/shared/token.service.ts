import { Injectable } from '@angular/core';
import { LoginResponse } from '../login/typedefs';

@Injectable({
    providedIn: 'root',
})
export class TokenService {
    private issuer = {
        login: '/api/v1/account/login',
        register: '/api/v1/account/create',
    };

    constructor() { }

    handleData(token: LoginResponse) {
        localStorage.setItem('access', token.access);
        localStorage.setItem('refresh', token.refresh);
    }

    getToken() {
        return localStorage.getItem('access');
    }

    // Verify the token
    isValidToken() {
        const token = this.getToken();
        if (token) {
            const payload = this.payload(token);
            if (payload) {
                return true
                // return Object.values(this.issuer).indexOf(payload.iss) > -1
                //     ? true
                //     : false;
            }
            return false;

        } else {
            return false;
        }
    }

    payload(token: any) {
        const jwtPayload = token.split('.')[1];
        return JSON.parse(atob(jwtPayload));
    }

    // User state based on valid token
    isLoggedIn() {
        return this.isValidToken();
    }

    // Remove token
    removeToken() {
        localStorage.removeItem('access');
        localStorage.removeItem('refresh');
    }
}
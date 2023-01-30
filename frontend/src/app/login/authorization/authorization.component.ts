import { Component } from '@angular/core'
import { FormBuilder, FormControl, Validators } from '@angular/forms';
import { AuthService } from '../AuthService.service';
import { LoginRequest, LoginResponse } from '../typedefs';
import { Router } from "@angular/router";
import { TokenService } from 'src/app/shared/token.service';


@Component({
    selector: 'app-authorization',
    templateUrl: './authorization.component.html'
})
export class AuthorizationComponent {
    error: string = ''
    loginForm = this.fb.group({
        email: new FormControl('test@test.com', [Validators.required, Validators.email]),
        password: new FormControl('Alex123_', [Validators.required, Validators.minLength(8)])
    });

    constructor(private fb: FormBuilder,
        private backend: AuthService,
        private router: Router,
        private token: TokenService) {
    }

    submitForm(): void {
        this.backend.login(<LoginRequest>this.loginForm.value)
            .subscribe({
                next: (data: LoginResponse) => {
                    this.token.handleData(data);
                    this.router.navigate(['']);
                },
                error: (error) => {
                    this.error = error.error.detail
                }
            });
    }
}

import { Component } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, Validators } from '@angular/forms';
import { AuthService } from '../AuthService.service';
import { RegisterRequest } from '../typedefs';


function createCompareValidator(controlOne: AbstractControl, controlTwo: AbstractControl) {
    return () => {
        if (controlOne.value !== controlTwo.value)
            return { match_error: 'Value does not match' };
        return null;
    };
}

@Component({
    selector: 'app-registration',
    templateUrl: 'registration.component.html'
})
export class RegistrationComponent {

    global_form_error: string = '';
    response: any = null;
    registerForm = this.fb.group({
        email: new FormControl('test@test.com', [Validators.required, Validators.email, Validators.maxLength(64)]),
        password: new FormControl('Alex123_', [Validators.required, Validators.minLength(8)]),
        password_confirm: new FormControl('Alex123_', [Validators.required, Validators.minLength(8)]),
    });

    constructor(private fb: FormBuilder,
        private authService: AuthService) {
        this.registerForm.addValidators(
            createCompareValidator(
                this.registerForm.controls.password,
                this.registerForm.controls.password_confirm
            )
        );
    }

    submit() {
        this.authService.register(<RegisterRequest>this.registerForm.value)
            .subscribe({
                next: (data: any) => {
                    if (data.status) this.registerForm.reset();
                    this.response = data;
                },
                error: (error) => {
                    console.log(error);
                    this.response = { status: false };
                }
            });
    }
}

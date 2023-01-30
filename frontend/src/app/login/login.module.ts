import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { LoginComponent } from './login.component';
import { LoginRoutingModule } from './login-routing.module';
import { RestorePasswordComponent } from './restore-password/restore-password.component';
import { AuthorizationComponent } from './authorization/authorization.component';
import { RegistrationComponent } from './registration/registration.component';

import { AuthService } from './AuthService.service';


@NgModule({
  declarations: [
    LoginComponent,
    RestorePasswordComponent,
    AuthorizationComponent,
    RegistrationComponent
  ],
  imports: [
    CommonModule,
    LoginRoutingModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
  ],
  providers: [AuthService]
})
export class LoginModule { }

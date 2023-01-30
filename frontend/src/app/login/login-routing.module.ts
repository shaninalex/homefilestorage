import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthorizationComponent } from './authorization/authorization.component';
import { LoginComponent } from './login.component';
import { RegistrationComponent } from './registration/registration.component';
import { RestorePasswordComponent } from './restore-password/restore-password.component';

const routes: Routes = [
  {
    path: '', component: LoginComponent, children: [
      { path: 'restore', component: RestorePasswordComponent },
      // { path: 'registration', component: RegistrationComponent },
      { path: 'login', component: AuthorizationComponent },
      { path: '**', redirectTo: 'login' }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class LoginRoutingModule { }

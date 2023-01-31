import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { AuthInterceptor } from '../shared/auth.interceptor';
import { DashboardComponent } from './dashboard.component';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { SidebarComponent } from './layout/sidebar/sidebar.component';
import { HeaderComponent } from './layout/header/header.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { FileIconComponent } from './globals/components/file-icon.component';


@NgModule({
    declarations: [
        DashboardComponent,
        SidebarComponent, 
        HeaderComponent,
    ],
    imports: [
        CommonModule,
        DashboardRoutingModule,
        FontAwesomeModule
    ],
    providers: [
        {
            provide: HTTP_INTERCEPTORS,
            useClass: AuthInterceptor,
            multi: true
        }
    ]
})
export class DashboardModule { }

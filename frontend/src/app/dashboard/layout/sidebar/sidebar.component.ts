import { Component, OnInit } from '@angular/core';
import { TokenService } from 'src/app/shared/token.service';


@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent implements OnInit {

  constructor(private token: TokenService) { }

  ngOnInit() {}

  sign_out(): void {
    this.token.removeToken();
    // empty Redux Store
    window.location.href = '/auth/login/';
  }

}

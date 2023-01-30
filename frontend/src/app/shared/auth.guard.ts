import { Injectable } from "@angular/core";
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, UrlTree, Router } from "@angular/router";
import { Observable } from "rxjs";
import { TokenService } from "./token.service";

@Injectable()
export class AuthGuard implements CanActivate {
    constructor(private tokenService: TokenService, private router: Router) { }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
        if (!this.tokenService.isValidToken()) {
            this.router.navigate(['/auth/']);
            return false;
        }
        return true;
    }
}
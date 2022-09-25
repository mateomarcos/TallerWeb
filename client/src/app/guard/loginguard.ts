import { LocalStorageService } from "../services/local-storage-service.service";

import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot } from '@angular/router';

@Injectable()
export class LoginGuard implements CanActivate {
    constructor(public storage: LocalStorageService, public router:Router) { }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
        if (this.storage.IsLoggedIn()) {
            this.router.navigate(['user/projects'])
            //PENDIENTE agregar popup de q ya estas logeado wacho
            return false;
        }
        return true;
    }
}

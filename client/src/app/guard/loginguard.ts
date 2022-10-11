import { LocalStorageService } from "../services/local-storage-service.service";

import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot } from '@angular/router';
import { ToastrService } from "ngx-toastr";
  /*Guards (from my understanding) are pre-requisite functions the client has to achieve in order to access a specific route. Used in the app routing module
  with a "canActivate" method, each guard does some kind of verification and then shows a pop-up message using the Toastr service.
  */
@Injectable()
export class LoginGuard implements CanActivate {
    constructor(public storage: LocalStorageService, public router:Router, public toastr: ToastrService ) { }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
        if (this.storage.IsLoggedIn()) {
            this.router.navigate(['user/projects'])
            this.toastr.error(`You are already authenticated!`, `Try loging out`)
            return false;
        }
        return true;
    }
}

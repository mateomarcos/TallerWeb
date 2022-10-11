import { LocalStorageService } from "../services/local-storage-service.service";

import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot } from '@angular/router';
import { ToastrService } from "ngx-toastr";
  /*Guards (from my understanding) are pre-requisite functions the client has to achieve in order to access a specific route. Used in the app routing module
  with a "canActivate" method, each guard does some kind of verification and then shows a pop-up message using the Toastr service.
  */
@Injectable()
export class RouteGuard implements CanActivate {
    constructor(public storage: LocalStorageService, public router:Router, public toastr: ToastrService ) { }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
        if (this.storage.IsLoggedIn()) {
            this.router.navigate(['/user/projects']);
        } else this.router.navigate(['/login']);
        this.toastr.info("This route doesn't exist!", "Invalid Access");
        return false;
    }
}

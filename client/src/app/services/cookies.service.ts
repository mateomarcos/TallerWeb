import { Injectable } from '@angular/core';
import {CookieService} from 'ngx-cookie-service';
import { JwtHelperService } from '@auth0/angular-jwt';
/*Interface to implement cookie support provided by ngx-cookie-service. Also adding some jwt decodification techniques to check for 
authentication and expiring tokens.*/

@Injectable({
  providedIn: 'root'
})
export class CookiesService {

  constructor(private cookieService: CookieService) { }

  Set(key: string, value: string) {
    this.cookieService.set(key, value);
}

Get(key: string) {
    return this.cookieService.get(key);
}

Remove(key: string) {
  this.cookieService.delete(key)
}

IsLoggedIn() : boolean {
    var token = this.cookieService.get("token");
    if (token == null) {
        return false;
    }
    var helper = new JwtHelperService;

    if (helper.isTokenExpired(token)) {
      this.cookieService.delete("token")
        return false;
    }

    return true;
}
}


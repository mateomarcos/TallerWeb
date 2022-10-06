import { Injectable } from '@angular/core';
import {CookieService} from 'ngx-cookie-service';
import { JwtHelperService } from '@auth0/angular-jwt';

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
    //console.log(helper.isTokenExpired(token))
    //console.log(helper.getTokenExpirationDate(token))
    if (helper.isTokenExpired(token)) {
      this.cookieService.delete("token")
        return false;
    }

    return true;
}
}


import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable()
export class LocalStorageService {

    Set(key: string, value: string) {
        localStorage.setItem(key, value);
    }

    Get(key: string) {
        return localStorage.getItem(key);
    }

    Remove(key: string) {
        localStorage.removeItem(key);
    }

    IsLoggedIn() : boolean {
        var token = localStorage.getItem("token")
        if (token == null) {
            return false;
        }
        var helper = new JwtHelperService;
        //console.log(helper.isTokenExpired(token))
        //console.log(helper.getTokenExpirationDate(token))
        if (helper.isTokenExpired(token)) {
            localStorage.removeItem("token");
            return false;
        }

        return true;
    }
}
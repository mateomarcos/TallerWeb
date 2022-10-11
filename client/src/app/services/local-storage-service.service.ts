import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
/*Interface to implement Local Storage support. Also adding some jwt decodification techniques to check for 
authentication and expiring tokens. */

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

        if (helper.isTokenExpired(token)) {
            localStorage.removeItem("token");
            return false;
        }

        return true;
    }
}
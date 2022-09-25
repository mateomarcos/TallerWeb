import { Injectable } from '@angular/core';

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
        if (localStorage.getItem("token") == null) {
            return false;
        }
        return true;
    }
}
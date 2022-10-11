import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import { LocalStorageService } from '../services/local-storage-service.service';
import { Emitters } from '../emitters/emitter';
import { CookiesService } from '../services/cookies.service';
  /* The nav bar is always on top of any content displayed in any route, giving the user quick access to some pages and functions depending on whether he is
  authenticated or not checking the status of the emitter. If he isn't, then he can get to Login and Signup pages from the nav bar and if he is then he can go to his own page or logout.
  Logout is a simple function that deletes the token stored in both cookies and localStorage.s
  */
@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.css']
})
export class NavComponent implements OnInit {
  authenticated = false;

  constructor(private http: HttpClient, private cookie:CookiesService) {
    var storage = new LocalStorageService;
    if (storage.IsLoggedIn()) {
      this.authenticated=true;
    }
  }

  ngOnInit(): void {
    Emitters.authEmitter.subscribe(
      (auth: boolean) => {
        this.authenticated = auth;
      }
    );
  }

  logout(): void {
    var storage = new LocalStorageService;
    storage.Remove("token");
    this.cookie.Remove("token");
    this.authenticated = false;
  }

}
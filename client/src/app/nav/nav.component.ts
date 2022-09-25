import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import { LocalStorageService } from '../services/local-storage-service.service';
import { Emitters } from '../emitters/emitter';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.css']
})
export class NavComponent implements OnInit {
  authenticated = false;

  constructor(private http: HttpClient) {
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
    /*var storage = new LocalStorageService;
    if (storage.IsLoggedIn()) {
      this.authenticated=true;
    }*/
  }

  logout(): void {
    var storage = new LocalStorageService;
    storage.Remove("token");
    this.authenticated = false;
  }

}
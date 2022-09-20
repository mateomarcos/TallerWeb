import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs/Observable";
import { CookieService } from "ngx-cookie-service";

@Injectable({
  providedIn: "root"
})
export class UserService {
  constructor(private http: HttpClient, private cookies: CookieService) {}

  login(user: Any): Observable<any> {
    return this.http.post("localhost:9000/login", user);
  }
  register(user: Any): Observable<any> {
    return this.http.post("localhost:9000/signup", user);
  }
  setToken(token: String) {
    this.cookies.set("token", token);
  }
  getToken() {
    return this.cookies.get("token");
  }
  //getUser() {
  //  return this.http.get("https://reqres.in/api/users/2");
  //}
  getUserLogged() {
    const token = this.getToken();
    // Aquí iría el endpoint para devolver el usuario para un token
  }
}
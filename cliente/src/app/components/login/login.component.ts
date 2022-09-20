import { Component } from '@angular/core';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  username: string;
  password: string;

  constructor(public userService: UserService) {}


  login() {
    const user = {username: this.username, password: this.password};
    this.userService.login(user).subscribe( data => {
      console.log(data);
    });
  }
}
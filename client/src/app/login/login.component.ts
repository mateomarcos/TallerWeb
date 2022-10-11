import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { LocalStorageService } from '../services/local-storage-service.service';
import { Emitters } from '../emitters/emitter';
import { ToastrService } from 'ngx-toastr';
import { CookiesService } from '../services/cookies.service';
  /*The login component renders a form which users can use to authenticate themselves and have access to the rest of the application's functionality.
  Each input type in the .html file is binded to a form in the same FormGroup that then is sent to the backend for final authentication, where a 
  token is received and stored in both localStorage and a cookie. Why both? Starting this project, I had opted for localStorage since you can store 
  5MB of Data vs 4KB given by cookies. However, they are not as safe as cookies so I decided to implement both since it is not a real life product.
  The submit function also emmits a true statement for the nav bar.
  */
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form: FormGroup;

  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router, private toastr: ToastrService, private cookie:CookiesService) {
    this.form = this.formBuilder.group({
      username:'',
      password:''
    });
  }

  ngOnInit(): void {}

  submit(): void  {
      this.http.post('http://localhost:9000/login', this.form.getRawValue(), {withCredentials:true}).subscribe( res => {
        this.router.navigateByUrl('user/projects')

        var storage = new LocalStorageService;
        var stringRes = JSON.stringify(res)
        var jsonRes = JSON.parse(stringRes)
        
        storage.Set("token",jsonRes["token"])
        this.cookie.Set("token",jsonRes["token"])

        Emitters.authEmitter.emit(true);
      })

  }
}
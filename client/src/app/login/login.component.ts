import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { LocalStorageService } from '../services/local-storage-service.service';
import { Emitters } from '../emitters/emitter';
import { ToastrService } from 'ngx-toastr';
import { CookiesService } from '../services/cookies.service';

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



        //console.log(jsonRes["token"])
        Emitters.authEmitter.emit(true);
      })

  }
}
import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { LocalStorageService } from '../services/local-storage-service.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  form: FormGroup;

  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router) {
    this.form = this.formBuilder.group({
      username:'', 
      password:''
    });
  }

  ngOnInit(): void {}

  //FUNCIONA PERO A VECES DEVUELVE 400 BAD REQUEST PORQUE LOS CAMPOS TIENEN LONGITUD MINIMA Y NO SE HACE CHEQUEO EN EL FRONTEND, hay que AGREGAR VALIDATORS
  submit(): void  {
      this.http.post('http://localhost:9000/signup', this.form.getRawValue()).subscribe(res => {
        //mensaje de login successfull
        this.router.navigateByUrl('/login')
      })

  }

}

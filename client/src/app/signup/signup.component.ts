import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { LocalStorageService } from '../services/local-storage-service.service';
  /* Since authentication is basically the main part of this project, clientes use this page to create an account.
  Each input type in the .html file is binded to a form in the same FormGroup that then is sent to the backend for database insertion.
  Since the backend does some validations to those inputs, sometimes an error is returned specifying the fields have failed in "min" tags.
  */
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

  submit(): void  {
      this.http.post('http://localhost:9000/signup', this.form.getRawValue()).subscribe(() => {
        this.router.navigateByUrl('/login')
      })

  }

}

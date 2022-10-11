import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

  /*This footer exposes information about the application and offers the user some references to other projects people have submitted recently.
    On Init does an http request to retrieve such projects from the backend.
  */
 
@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css']
})
export class FooterComponent implements OnInit {
  projects: any = "";

  constructor(private http: HttpClient, private router: Router) {
  }


  ngOnInit(): void {
    this.http.get('http://localhost:9000/activeUsers').subscribe(data => {
      var stringRes = JSON.stringify(data)
      this.projects = JSON.parse(stringRes)
    })
  }

}

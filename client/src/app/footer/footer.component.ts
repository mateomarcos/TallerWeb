import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

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
      //console.log(this.users)
    })
  }

}

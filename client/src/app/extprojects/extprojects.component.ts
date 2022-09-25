import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-extprojects',
  templateUrl: './extprojects.component.html',
  styleUrls: ['./extprojects.component.css']
})
export class ExtprojectsComponent implements OnInit {
  projects: any = "";
  user: string | null;

  constructor(private http: HttpClient, private router: Router, private aRouter: ActivatedRoute) { 
    this.user = this.aRouter.snapshot.paramMap.get('username');
    console.log(this.aRouter.snapshot.paramMap.get('username'))
  }

  ngOnInit(): void {
    this.getExtProjects();
  }

  getExtProjects() {
    console.log(this.user)
    var route = 'http://localhost:9000/user/' +this.user+ '/projects'
    this.http.get(route).subscribe(data => {
      var stringRes = JSON.stringify(data)
      this.projects = JSON.parse(stringRes)
      console.log(this.projects)
      //console.log(this.projects[0]["author"])
    }, error => {
      console.log(error);
    })
  }


}

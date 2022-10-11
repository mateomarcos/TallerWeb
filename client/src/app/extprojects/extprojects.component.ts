import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute, Router } from '@angular/router';
  /*The exterior projects component works as a possible route within the application to show other user's pages. Retrieves data using get http calls to the server.
  */
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
  }

  ngOnInit(): void {
    this.getExtProjects();
  }

  getExtProjects() {
    var route = 'http://localhost:9000/user/' +this.user+ '/projects'
    this.http.get(route).subscribe(data => {
      var stringRes = JSON.stringify(data)
      this.projects = JSON.parse(stringRes)
      console.log(this.projects)
    }, error => {
      this.router.navigateByUrl('/user/projects')
    })
  }


}

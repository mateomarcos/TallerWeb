import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
  /*The projects component is used to display the authenticated users' projects, retrieveng the information from the backend when the page loads and allowing him to
  read them, delete them and add a new one using the same form style input used in Login, binding each input to a Form that is then sent to the backend.
  When deleting a project, the backend returns a deletion value created by the mongodb driver. However, since it has no use for us in the frontend I do not manage
  the response in the callback function.
  */
@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.css']
})
export class ProjectsComponent implements OnInit {
  form: FormGroup;
  projects: any = "";

  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router) {
    this.form = this.formBuilder.group({
      name:'', 
      description:'', 
      repository:'', 
    }
    );
  }

  ngOnInit(): void {
    this.getProjects();
  }

  getProjects() {
    this.http.get('http://localhost:9000/user/projects').subscribe(data => {
      var stringRes = JSON.stringify(data)
      this.projects = JSON.parse(stringRes)
    })
  }

  submit(): void  {
      console.log(this.form.getRawValue())
      this.http.post('http://localhost:9000/user/projects', this.form.getRawValue()).subscribe( () => {
        window.location.reload()
      })

  }

  delete(name: string): void{
    console.log(name);
    var route = 'http://localhost:9000/user/projects/' +name;
    this.http.delete(route).subscribe(() => {
      window.location.reload();

    });
  }
}

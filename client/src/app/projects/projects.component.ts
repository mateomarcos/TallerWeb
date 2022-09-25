import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { Project } from '../models/project';

@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.css']
})
export class ProjectsComponent implements OnInit {
  form: FormGroup;
  projectsList: Project[] = [];
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
      console.log(this.projects)
      //console.log(this.projects[0]["author"])
    }, error => {
      console.log(error);
    })
  }

  submit(): void  {
      console.log(this.form.getRawValue())
      this.http.post('http://localhost:9000/user/projects', this.form.getRawValue()).subscribe( () => {
        this.router.navigateByUrl('/login') //cambiar por la misma pagina de los proyectos para que se recargue
      })

  }

}

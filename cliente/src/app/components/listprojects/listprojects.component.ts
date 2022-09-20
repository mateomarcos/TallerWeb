import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Project } from 'src/app/models/project';

@Component({
  selector: 'app-listprojects',
  templateUrl: './listprojects.component.html',
  styleUrls: ['./listprojects.component.css']
})
export class ListprojectsComponent implements OnInit {
  projectForm: FormGroup

  constructor(private fb: FormBuilder,
              private router: Router) {
    this.projectForm = this.fb.group({
      name:['', Validators.required],
      description:['', Validators.required],
      repository:['', Validators.required]
    })
   }

  ngOnInit(): void {
  }

  addProject() {    
    const PROJECT: Project = {
      Name: this.projectForm.get('name')?.value,
      Description: this.projectForm.get('description')?.value,
      Repository: this.projectForm.get('repository')?.value
    }

    this.router.navigate(['user/projects'])
  }

}

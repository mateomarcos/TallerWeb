import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ListextprojectsComponent } from './components/listextprojects/listextprojects.component';
import { ListprojectsComponent } from './components/listprojects/listprojects.component';

//Components
import { LoginComponent } from './components/login/login.component';
import { SignupComponent } from './components/signup/signup.component';

const routes: Routes = [
  {path: '', component: LoginComponent},
  {path: 'login', component: LoginComponent},
  {path: 'signup', component: SignupComponent},
  {path: 'user/projects', component: ListprojectsComponent},
  {path: 'user/:username/projects', component: ListextprojectsComponent},



  {path: '**', redirectTo: '', pathMatch: 'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

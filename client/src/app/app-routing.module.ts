import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import { ExtprojectsComponent } from './extprojects/extprojects.component';
import { AuthGuard } from './guard/authguard';
import { LoginGuard } from './guard/loginguard';
import { RouteGuard } from './guard/routeguard';
import {LoginComponent} from './login/login.component';
import { ProjectsComponent } from './projects/projects.component';
import { SignupComponent } from './signup/signup.component';


const routes: Routes = [
  {path: '', component: LoginComponent, canActivate: [LoginGuard]},
  {path: 'login', component: LoginComponent, canActivate: [LoginGuard]},
  {path: 'signup', component: SignupComponent, canActivate: [LoginGuard]},
  {path: 'user/projects', component: ProjectsComponent, canActivate: [AuthGuard]},
  {path: 'user/:username/projects', component: ExtprojectsComponent, canActivate: [AuthGuard]},
  {path: '**', redirectTo:'' ,canActivate:[RouteGuard]}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import { ExtprojectsComponent } from './extprojects/extprojects.component';
import { AuthGuard } from './guard/authguard';
import { LoginGuard } from './guard/loginguard';
import { RouteGuard } from './guard/routeguard';
import {LoginComponent} from './login/login.component';
import { ProjectsComponent } from './projects/projects.component';
import { SignupComponent } from './signup/signup.component';
/* Possible routes of the application are specified in this module. Each component has a Guard to manage traffic within the application and a wildcard and unkownroute
path to avoid unexisting routes.
*/

const routes: Routes = [
  {path: '', component: LoginComponent, canActivate: [LoginGuard]},
  {path: 'login', component: LoginComponent, canActivate: [LoginGuard]},
  {path: 'signup', component: SignupComponent, canActivate: [LoginGuard]},
  {path: 'user/projects', component: ProjectsComponent, canActivate: [AuthGuard]},
  {path: 'user/:username/projects', component: ExtprojectsComponent, canActivate: [AuthGuard]},
  {path: 'unknownRoute', component:LoginComponent ,canActivate:[RouteGuard]},
  {path: '**', redirectTo:'unknownRoute'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
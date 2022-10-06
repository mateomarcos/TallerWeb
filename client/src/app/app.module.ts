import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SignupComponent } from './signup/signup.component';
import { LoginComponent } from './login/login.component';
import { NavComponent } from './nav/nav.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { UniversalAppInterceptor } from './services/http-interceptor.service';
import { ProjectsComponent } from './projects/projects.component';
import { ExtprojectsComponent } from './extprojects/extprojects.component';
import { LoginGuard } from './guard/loginguard';
import { LocalStorageService } from './services/local-storage-service.service';
import { AuthGuard } from './guard/authguard';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ToastrModule } from 'ngx-toastr';
import { FooterComponent } from './footer/footer.component';
import { RouteGuard } from './guard/routeguard';

@NgModule({
  declarations: [
    AppComponent,
    SignupComponent,
    LoginComponent,
    NavComponent,
    ProjectsComponent,
    ExtprojectsComponent,
    FooterComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot()
  ],
  providers: [    {
    provide: HTTP_INTERCEPTORS,
    useClass: UniversalAppInterceptor,
    multi: true,
  },
  LoginGuard,
  LocalStorageService,
  AuthGuard,
  RouteGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }

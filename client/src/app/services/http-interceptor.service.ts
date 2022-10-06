import { Injectable, Inject, Optional } from '@angular/core';
import { 
  HttpEvent, HttpRequest, HttpHandler, 
  HttpInterceptor, HttpErrorResponse 
} from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';
import { LocalStorageService } from './local-storage-service.service';
import { ToastrService } from 'ngx-toastr';

@Injectable()
export class UniversalAppInterceptor implements HttpInterceptor {

  constructor(private toastr: ToastrService) { }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    var storage = new LocalStorageService;
    const token = storage.Get("token")
    if (token != null) {
      req = req.clone({
        url:  req.url,
        setHeaders: {
          Token: `${token}`
        }
      });
    }
    return next.handle(req).pipe(
      catchError(error => {
        let errorMessage = '';
        if (error instanceof ErrorEvent) {
          // client-side error
          errorMessage = `Client-side error: ${error.error.message}`;
        } else {
          // backend error
          errorMessage = `Server-side error: ${error.status} \n ${error.message} \n ${error.error.error}`;
        }
        //console.log(errorMessage);
        console.log(error)
        this.toastr.error(`${error.error.error}`, `Error ${error.status}: ${error.statusText}`)
        return throwError(() => errorMessage);
      })
    );
  }
}
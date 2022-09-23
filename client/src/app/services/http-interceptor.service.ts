import { Injectable, Inject, Optional } from '@angular/core';
import { HttpInterceptor, HttpHandler, HttpRequest } from '@angular/common/http';
import { LocalStorageService } from './local-storage-service.service';

@Injectable()
export class UniversalAppInterceptor implements HttpInterceptor {

  constructor() { }

  intercept(req: HttpRequest<any>, next: HttpHandler) {
    var storage = new LocalStorageService;
    const token = storage.Get("token")
    if (token != null) {
      //Si hay token, analizar si las url son login y signup, importar router y directamente redirect ya que ya esta autenticado.
      req = req.clone({
        url:  req.url,
        setHeaders: {
          Token: `${token}`
        }
      });
    }
    return next.handle(req);
  }
} 
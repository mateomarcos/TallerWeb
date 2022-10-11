import {EventEmitter} from '@angular/core';
  /*An EventEmitter just implements a method that helps data sharing from emitter to subscriber. Here it is used to emit 
  the authentication status of an user so some componentes can be rendered accordingly. (i.e NAVigation bar)
  */
export class Emitters {
  static authEmitter = new EventEmitter<boolean>();
}
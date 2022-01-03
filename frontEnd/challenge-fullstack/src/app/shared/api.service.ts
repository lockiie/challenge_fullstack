import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { MatSnackBar } from '@angular/material/snack-bar';

@Injectable({
  providedIn: 'root'
})

export class ApiService {

  endpoint: string = 'http://127.0.0.1:4000/api/v0';
  headers = new HttpHeaders().set('Content-Type', 'application/json');

  constructor(private http: HttpClient, private snackBar: MatSnackBar) { }

  Add(uri: string, data: any): Observable<any> {
    let API_URL = `${this.endpoint}/${uri}`;
    return this.http.post(API_URL, data)
      .pipe(
        catchError((error) => {
          this.errorMgmt(error)
          return throwError(() => error);
        }),
      )
  }

  Get(uri: string): Observable<any> {
    return this.http.get(`${this.endpoint}/${uri}`);
  }

  GetByID(uri: string, id: number): Observable<any> {
    let API_URL = `${this.endpoint}/${uri}/${String(id)}`;
    return this.http.get(API_URL, { headers: this.headers })
      .pipe(
        map((res: any) => {
          return res || {}
        }),
        catchError((error) => {
          this.errorMgmt(error)
          return throwError(() => error);
        }),
      )
  }

  Update(uri: string, id: number, data: any): Observable<any> {
    let API_URL = `${this.endpoint}/${uri}/${String(id)}`;
    return this.http.put(API_URL, data, { headers: this.headers })
      .pipe(
        catchError((error) => {
          this.errorMgmt(error)
          return throwError(() => error);
        }),
      )
  }

  Delete(uri: string, id : number): Observable<any> {
    var API_URL = `${this.endpoint}/${uri}/${String(id)}`;
    return this.http.delete(API_URL)
      .pipe(
        catchError((error) => {
          this.errorMgmt(error)
          return throwError(() => error);
        }),
      )
  }

  openSnackBar(message: string, action: string): void {
    this.snackBar.open(message, action, {
        duration: 2000,
        panelClass: ['snackbar'],
    });
}

  errorMgmt(error: HttpErrorResponse) {
    console.log(error.error.message)
    let errorMessage = '';
    if (error.error!.message! != '') {
      errorMessage = error.error.message;
    } else {
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    this.openSnackBar(errorMessage, 'ok');
  }

}
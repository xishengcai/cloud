import {Component, Injectable} from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { HttpErrorHandler, HandleError } from '../../http-error-handler.service';
import { Cluster } from './clusters'
import { Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import {HttpResponse} from "../job";

@Injectable()
export class ClustersService{
  handleError: HandleError;

  httpOptions = {
    headers: new HttpHeaders({
      'Context-Type': 'Application/json',
    })
  }
  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('ClustersService');
  }

  /** GET heroes from the server */
  getClusters(): Observable<Cluster[]> {
    return this.http.get<Cluster[]>("/api/v1/cluster/masters")
      .pipe(
        catchError(this.handleError('getClusters', []))
      );
  }


  addCluster(cluster: Cluster){
    this.http.post<HttpResponse>("/api/v1/cluster/masters",cluster,this.httpOptions).subscribe({
      next: data =>{
      if (data.eventCode != 0){
        alert(data.resMsg)
        }},
      error: error => {
      console.error('There was an error!', error.message);
    }})
  }
}

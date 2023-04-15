import {Component, Injectable} from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { HttpErrorHandler, HandleError } from '../../http-error-handler.service';
import {HttpResponse} from "../job";
import {Host} from "../clusters/clusters";

@Injectable()
export class SlaveService{
  handleError: HandleError;

  httpOptions = {
    headers: new HttpHeaders({
      'Context-Type': 'Application/json',
    })
  }
  constructor(
    private http: HttpClient,
    httpErrorHandler: HttpErrorHandler) {
    this.handleError = httpErrorHandler.createHandleError('Install slave nodes error');
  }

  addSlave(slave: Slave){
    this.http.post<HttpResponse>("/api/v1/cluster/slaves",slave,this.httpOptions).subscribe({
      next: data =>{
        if (data.eventCode != 0){
          alert(data.resMsg)
        }},
      error: error => {
        console.error('There was an error!', error.message);
      }})
  }
}
export class Slave {
  constructor(
    public master: Host,
    public nodes: Host[],
  ) {}
}


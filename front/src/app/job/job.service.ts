import {Component, Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {HttpResponse, Job} from './job'
import { Observable } from 'rxjs';


@Injectable()
export class JobService{
  constructor(private http: HttpClient) {
  }
  httpOptions = {
    headers: new HttpHeaders({
      'Context-Type': 'Application/json',
      responseType: 'json'
    })}

  public jobs$: Job[] = []


  getJobs(): Job[] {

     this.http.get<any>("/api/v1/cluster/jobs?namespace=install_k8s&jobType=master",this.httpOptions).subscribe(
      resp => {
        let response: HttpResponse = resp
        this.jobs$ = response.data
      }
    )
    this.jobs$.forEach((val, idx, array)=>{
        val.argsJson = JSON.stringify(val.args)
        val.date= new Date(val.t);
    })
    return this.jobs$
  }

}

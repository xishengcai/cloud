import { Component, OnInit } from '@angular/core';
import { FormBuilder, NgForm } from "@angular/forms"
import { Job } from "./job"
import { Observable } from "rxjs";
import { JobService } from "./job.service";
import { CommonModule } from "@angular/common";
import {HttpClient, HttpHeaders} from '@angular/common/http';

@Component({
  selector: 'app-job',
  templateUrl: './job.component.html',
  styleUrls: ["./job.component.css"],
})

export class JobComponent implements OnInit {
  jobs: Job[]=[];

  constructor(
    private jobService: JobService,
  ) {}

  ngOnInit(): void {
    this.getJobs()
  }

  getJobs(): void{
     this.jobs = this.jobService.getJobs()
  }
}

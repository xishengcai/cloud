import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { JobComponent } from './job/job.component';
import { ClustersComponent } from './job/clusters/clusters.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { HttpClientModule } from '@angular/common/http';
import { HttpErrorHandler } from './http-error-handler.service';
import { MessageService } from './message.service';
import { AppRoutingModule } from "./app-routing.module";
import { JobService } from "./job/job.service";
import { ResponseComponent } from './response/response.component';
import { SlaveComponent } from './job/slave/slave.component';


@NgModule({
  declarations: [
    AppComponent,
    JobComponent,
    ClustersComponent,
    PageNotFoundComponent,
    ResponseComponent,
    SlaveComponent,
  ],
  imports: [
    BrowserModule,
    ReactiveFormsModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [
    HttpErrorHandler,
    MessageService,
    JobService
    ],
  bootstrap: [AppComponent]
})
export class AppModule {

}

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {JobComponent} from "./job/job.component";
import {ClustersComponent} from "./job/clusters/clusters.component";
import {SlaveComponent} from "./job/slave/slave.component";


const routes: Routes = [
  {path: 'job', component: JobComponent},
  {path: 'cluster', component: ClustersComponent},
  {path: 'nodes', component: SlaveComponent},
]

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}

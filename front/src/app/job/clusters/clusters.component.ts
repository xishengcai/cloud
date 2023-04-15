import { Component, OnInit } from '@angular/core';
import { FormControl,FormGroup,FormBuilder} from '@angular/forms';
import { ClustersService } from "./clusters.service";

@Component({
  selector: 'app-clusters',
  providers: [ClustersService],
  templateUrl: './clusters.component.html',
  styleUrls: ["./clusters.component.css"],
})

export class ClustersComponent implements OnInit{
  netWorkPlugs = ['cilium','calico']
  registries = ['registry.aliyuncs.com/google_containers','k8s.gcr.io'];
  profileForm = new FormGroup({});

  ngOnInit(): void {
    this.profileForm = this.fb.group({
      name: new FormControl(''),
      controlPlaneEndpoint: new FormControl(''),
      netWorkPlug: new FormControl(''),
      podCidr: new FormControl('10.244.0.0/16'),
      serviceCidr: new FormControl('10.96.0.0/16'),
      version: new FormControl('1.17.11'),
      networkPlug: this.netWorkPlugs[0],
      registry: this.registries[0],
      primaryMaster: this.fb.group(
        {
          ip: new FormControl(''),
          port: new FormControl(22),
          user: new FormControl('root'),
          password: new FormControl(''),
        }
      ),
    });
  }

  constructor(
    private fb: FormBuilder,
    private http: ClustersService,
  ) {}

  onSubmit() {
    console.log(this.profileForm.value)
    this.http.addCluster(this.profileForm.value)
  }

}



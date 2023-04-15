import{Host} from "./clusters/clusters"
export class Job {
  constructor(
    public name: string,
    public id: string,
    public t: number,
    public date: Date,
    public args: InstallMasterArgs,
    public argsJson: string,
    public unique: boolean) {
  }
}

export class HttpResponse {
  constructor(
    public eventCode: number,
    public resMsg: string,
    public data: any
  ) {
  }
}


export class InstallMasterArgs {
  constructor(
    public name: string,
    public clusterName: string,
    public  controlPlaneEndpoint: string,
    public registry: string,
    public primaryMaster: Host,
    public netWorkPlug: string,
    public podCidr: string,
    public serviceCidr: string,
    public version: string,
  ) {
  }
}

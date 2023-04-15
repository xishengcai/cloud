export class Cluster {
  constructor(
    public name: string,
    public controlPlaneEndpoint: string,
    public registry: string,
    public netWorkPlug: string,
    public podCidr: string,
    public serviceCidr: string,
    public version: string,
    public port: number,
    public primaryMaster: Host[],
    ) {

  }

}

export class Host {
  constructor(
    public ip: string,
    public password: string,
    public port: string,
    public user: string,
  ) {
  }
}

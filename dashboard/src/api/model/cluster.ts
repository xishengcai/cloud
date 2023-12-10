
export interface Host {
    ip: string;
    port: number;
    password: string;
}

export interface Cluster {
    id: string;
    name: string;
    netWorkPlug: string;
    registry: string;
    version: string;
    controlPlaneEndpoint: string;
    podCidr: string;
    serviceCidr: string;
    master: Host[];
    slaveNode: Host[];
}


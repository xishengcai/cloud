import { Component, OnInit } from '@angular/core';
import {Host} from "../clusters/clusters";
import {FormBuilder, FormControl, FormGroup, FormArray} from "@angular/forms";
import {SlaveService} from "./slave.service";

@Component({
  selector: 'app-slave',
  providers:[SlaveService],
  templateUrl: './slave.component.html',
  styleUrls: ['./slave.component.css']
})

export class SlaveComponent implements OnInit {
  profileForm = new FormGroup({});
  public slaves = new FormArray([]);


  ngOnInit(): void {
    this.profileForm = this.fb.group({
      master: this.fb.group({
        ip: new FormControl(''),
        port: new FormControl(22),
        user: new FormControl('root'),
        password: new FormControl(''),
      }),
      slaves: this.slaves
    });
    this.addSlaveItem()

  }

  constructor(
    private fb: FormBuilder,
    private http: SlaveService,
  ) {}

  get nodeArray() {
    return <FormArray>this.profileForm.get('nodeArray');
  }

  addSlaveItem(){
    this.slaves.push(
      this.fb.group(
        {
          ip: new FormControl(''),
          port: new FormControl(22),
          user: new FormControl('root'),
          password: new FormControl(''),
        }
      )
    )
  }

  removeFeeItem() {
    this.slaves.removeAt(this.slaves.length - 1);
  }

  onSubmit() {
    console.log(this.profileForm.value)
    this.http.addSlave(this.profileForm.value)
  }
}

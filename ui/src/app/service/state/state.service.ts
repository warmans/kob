import { Injectable } from '@angular/core';
import { Observable,  } from 'rxjs/Observable';
import { BehaviorSubject  } from 'rxjs/BehaviorSubject';

@Injectable()
export class StateService {

  private state:  { [id: string]: any } = JSON.parse(localStorage.getItem('state')) || {};
  public onStateChange: BehaviorSubject<{ [id: string]: any }> = new BehaviorSubject(this.state);

  constructor() {
  }

  pushKey(key: string, value: any) {
    this.state[key] = value;
    localStorage.setItem('state', JSON.stringify(this.state));
    this.onStateChange.next(this.state);
  }
}

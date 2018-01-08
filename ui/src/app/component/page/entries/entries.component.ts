import { Component, OnInit } from '@angular/core';
import { StateService } from '../../../service/state/state.service';
import { NotificationService } from '../../../service/notification/notification.service';

@Component({
  selector: 'app-entries',
  templateUrl: './entries.component.html',
  styleUrls: ['./entries.component.scss']
})
export class EntriesComponent implements OnInit {

  public claims: any;

  constructor(private state: StateService, private notify: NotificationService) { }

  ngOnInit() {
    this.state.onStateChange.subscribe((newState) => {
      this.claims = parseJwt(newState['token']);
      console.log(this.claims);
    });
  }
}

function parseJwt (token) {
  const base64Url = token.split('.')[1];
  const base64 = base64Url.replace('-', '+').replace('_', '/');
  return JSON.parse(window.atob(base64));
}

import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { StateService } from '../../../service/state/state.service';

@Component({
  selector: 'app-login-complete',
  templateUrl: './login-complete.component.html',
  styleUrls: ['./login-complete.component.scss']
})
export class LoginCompleteComponent implements OnInit {

  public err: string;
  public token: string;

  constructor(private router: Router, private activatedRoute: ActivatedRoute, private state: StateService) { }

  ngOnInit() {
    this.activatedRoute.queryParams.subscribe(params => {
      this.err = params['err'];
      this.token = params['token'];
      if (this.err === '' && this.token !== '') {
        console.log(this.token);
        this.state.pushKey('token', this.token);
        this.router.navigate(['/entries']);
      }
    });
  }
}

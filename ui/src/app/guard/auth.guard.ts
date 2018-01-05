import { StateService } from '../service/state/state.service';
import { Router, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Injectable } from '@angular/core';

@Injectable()
export class AuthGuard {

    private curState: {[index: string]: any} = {};

    constructor(private state: StateService, private router: Router) {
        state.onStateChange.subscribe((newState) => {
            this.curState = newState;
        });
    }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
        const token = this.curState['token'];
        if (token !== undefined && token !== '' && token !== null) {
            return true;
        }
        this.router.navigate(['/login']);
    }
}

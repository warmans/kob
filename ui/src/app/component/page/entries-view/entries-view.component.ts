import {Component, OnInit} from '@angular/core';
import {ApiClientService} from '../../../service/api-client';
import {ActivatedRoute} from "@angular/router";
import {rpcEntry} from "../../../service/api-client/models/rpcentry.model";
import {Subscription} from "rxjs/Subscription";
import {rpcAuthor} from "../../../service/api-client/models/rpcauthor.model";

@Component({
  selector: 'app-entries-view',
  templateUrl: './entries-view.component.html',
  styleUrls: ['./entries-view.component.scss']
})
export class EntriesViewComponent implements OnInit {

  public entry: rpcEntry = <rpcEntry>{author: <rpcAuthor>{}};
  private sub: Subscription;

  constructor(private apiClient: ApiClientService, private route: ActivatedRoute) {
  }

  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.apiClient.GetEntry(params['id']).subscribe((data) => {
        this.entry = data.body;
      });
    });
  }
}

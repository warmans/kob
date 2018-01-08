import { Component, OnInit } from '@angular/core';
import { rpcCreateEntryRequest } from '../../../service/api-client/models/rpccreateentryrequest.model';
import { ApiClientService } from '../../../service/api-client';
import {Router} from "@angular/router";

@Component({
  selector: 'app-entries-edit',
  templateUrl: './entries-edit.component.html',
  styleUrls: ['./entries-edit.component.scss']
})
export class EntriesEditComponent implements OnInit {

  public entry: rpcCreateEntryRequest = <rpcCreateEntryRequest>{};

  constructor(private apiClient: ApiClientService, private router: Router) { }

  ngOnInit() {
  }

  public create() {
    console.log(this.entry);
    this.apiClient.CreateEntry(this.entry).subscribe((data) => {
      this.router.navigateByUrl("/entries/view/"+data.body.id)
    });
  }

}

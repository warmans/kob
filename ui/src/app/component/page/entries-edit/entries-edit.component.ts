import { Component, OnInit } from '@angular/core';
import { rpcCreateEntryRequest } from '../../../service/api-client/models/rpccreateentryrequest.model';
import { ApiClientService } from '../../../service/api-client/index';

@Component({
  selector: 'app-entries-edit',
  templateUrl: './entries-edit.component.html',
  styleUrls: ['./entries-edit.component.scss']
})
export class EntriesEditComponent implements OnInit {

  public entry: rpcCreateEntryRequest = <rpcCreateEntryRequest>{};

  constructor(private apiClient: ApiClientService) { }

  ngOnInit() {
  }

  public create() {
    console.log(this.entry);
    this.apiClient.CreateEntry(this.entry).subscribe((data) => {
      console.log(data);
    });
  }

}

import { Component, OnInit } from '@angular/core';
import { rpcEntryList } from '../../../service/api-client/models/rpcentrylist.model';
import { ApiClientService } from '../../../service/api-client/index';

@Component({
  selector: 'app-entries-index',
  templateUrl: './entries-index.component.html',
  styleUrls: ['./entries-index.component.scss']
})
export class EntriesIndexComponent implements OnInit {

  public entryList: rpcEntryList = <rpcEntryList>{};

  constructor(private apiClient: ApiClientService) { }

  ngOnInit() {
    this.apiClient.ListEntries('0', '25').subscribe(
      (data) => {
        this.entryList = data.body;
        console.log(data);
      },
      (err) => {
        console.log(err); //todo: message service
      }
    );
  }
}

import { rpcEntry } from './rpcentry.model';

export interface rpcEntryList {
  num_results: string;
  entries: rpcEntry[];
}

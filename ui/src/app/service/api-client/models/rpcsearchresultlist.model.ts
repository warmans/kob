import { rpcSearchResult } from './rpcsearchresult.model';

export interface rpcSearchResultList {
  num_results: string;
  max_score: number;
  results: rpcSearchResult[];
}

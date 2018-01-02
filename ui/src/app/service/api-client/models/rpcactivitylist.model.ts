import { rpcActivity } from './rpcactivity.model';

export interface rpcActivityList {
  num_results: string;
  activities: rpcActivity[];
}

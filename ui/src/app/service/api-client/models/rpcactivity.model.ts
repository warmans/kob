import { ActivityAction } from './activityaction.model';
import { rpcAuthor } from './rpcauthor.model';

export interface rpcActivity {
  id: string;
  description: string;
  URI: string;
  action: ActivityAction;
  author: rpcAuthor;
}

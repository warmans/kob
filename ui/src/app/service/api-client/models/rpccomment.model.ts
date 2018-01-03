import { rpcAuthor } from './rpcauthor.model';

export interface rpcComment {
  id: string;
  content: string;
  created_date: string;
  updated_date: string;
  author: rpcAuthor;
}

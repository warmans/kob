import { rpcAuthor } from './rpcauthor.model';

export interface rpcComment {
  id: string;
  created_date: string;
  content: string;
  author: rpcAuthor;
}

import { rpcAuthor } from './rpcauthor.model';
import { rpcTag } from './rpctag.model';

export interface rpcEntry {
  id: string;
  title: string;
  content: string;
  created_date: string;
  updated_date: string;
  author: rpcAuthor;
  tags: rpcTag[];
}

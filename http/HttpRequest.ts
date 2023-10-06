import { Request, Response} from 'express';

export interface HttpRequest{
    handle(res: Request, Response): void;
}
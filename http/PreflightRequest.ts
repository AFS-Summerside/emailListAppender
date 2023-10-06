import { Request, Response, Application } from 'express/lib/express';
import { HttpHeaders } from './HttpHeaders';
import { HttpRequest } from './HttpRequest';

class PreflightRequest implements HttpRequest{
    handle(req: Request, res: Response){
        res.setHeader.apply(HttpHeaders.accessControlAllowOrigin());
        res.setHeader.apply(HttpHeaders.accessControlAllowMethods());
        res.setHeader.apply(HttpHeaders.accessControlAllowHeaders());
        res.status(200).send();
    }
}
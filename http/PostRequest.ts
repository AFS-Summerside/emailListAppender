import { Request, Response} from 'express';
import { HttpHeaders } from './HttpHeaders';
import { HttpRequest } from './HttpRequest';

class PostRequest implements HttpRequest{

    handle(req: Request, res: Response) {
        res.setHeader.apply(HttpHeaders.accessControlAllowOrigin());
        res.setHeader.apply(HttpHeaders.accessControlAllowMethods());
        res.setHeader.apply(HttpHeaders.accessControlAllowHeaders());
        res.setHeader.apply(HttpHeaders.contentType());
        // message = aMessage
        res.status(200).json(this);
    }
}
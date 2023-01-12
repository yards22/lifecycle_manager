import { Request } from "../Utils/Fetch";

export class PollsRepo{
    baseUrl: string;
    rq: Request;

    constructor(baseUrl:string,rq:Request){
        this.baseUrl = baseUrl
        this.rq = rq;
    }

    
}
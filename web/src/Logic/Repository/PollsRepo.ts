import { AuthHeadersWithToken } from "../Utils/AuthHeaders";
import { Request } from "../Utils/Fetch";

export class PollsRepo{
    baseUrl: string;
    rq: Request;

    constructor(baseUrl:string,rq:Request){
        this.baseUrl = baseUrl
        this.rq = rq;
    }

    async getPolls(token:string){
        try {
            const response = await this.rq.Get(this.baseUrl,AuthHeadersWithToken(token));
            const res = response.json();
            return res
        } catch (error) {
            throw error
        }
    }

    async addPoll(poll_question:string, options:string[], token:string){
        try{
            const response = await this.rq.Post(this.baseUrl,{poll_question,options_count:options.length,options },AuthHeadersWithToken(token));
            const res = response.json();
            return res
        }catch(error){
            throw error
        }
    }
    
}
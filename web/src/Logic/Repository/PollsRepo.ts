import { AuthHeaders, AuthHeadersWithToken } from "../Utils/AuthHeaders";
import { Request } from "../Utils/Fetch";
import { CheckResponse } from "../Utils/ResponseHandler";

export class PollsRepo{
    baseUrl: string;
    rq: Request;

    constructor(baseUrl:string,rq:Request){
        this.baseUrl = baseUrl
        this.rq = rq;
    }

    async getPolls(token:string){
        try {
            console.log(token);
            const response = await this.rq.Get(this.baseUrl,AuthHeaders(token));
            const {body} = await CheckResponse(response,200)
            return {
                polls:body.data
            }
        } catch (error) {
            throw error
        }
    }

    async addPoll(poll_question:string, options:string[], token:string){
        try{
            const response = await this.rq.Post(this.baseUrl,{poll_question,options_count:options.length,options },AuthHeaders(token));
            const res = response.json();
            return res
        }catch(error){
            throw error
        }
    }
    
}